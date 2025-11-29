// Package spatial includes all spatial code
package spatial

import (
	"errors"
	"fmt"
	"strings"
)

type TileMap[T comparable] struct {
	Rect  Rect
	Cells []T
}

var (
	ErrOutOfBounds       = errors.New("out of bounds")
	ErrUndersized        = errors.New("not enough cells")
	ErrNoData            = errors.New("no data")
	ErrUndefinedAlias    = errors.New("alias not defined in legend")
	ErrUnequalRowLengths = errors.New("unequal row lengths")
)

func NewTileMap[T comparable](w, h int, cells []T) (*TileMap[T], error) {
	if len(cells) != w*h {
		return nil, ErrUndersized
	}
	return &TileMap[T]{
		Rect:  NewRectFromZero(Coord{w, h}),
		Cells: cells,
	}, nil
}

func normaliseStrMap(strmap string) ([]string, int, error) {
	lines := strings.Split(strings.TrimSpace(strmap), "\n")
	if len(lines) == 0 {
		return nil, 0, fmt.Errorf("strmap has no rows (%w)", ErrNoData)
	}
	width := -1
	for i := range lines {
		trimmed := strings.TrimSpace(lines[i])
		length := len([]rune(trimmed))
		if width > 0 && length != width {
			return nil, 0, ErrUnequalRowLengths
		}
		width = length
		lines[i] = trimmed
	}
	if width <= 0 {
		return nil, 0, fmt.Errorf("strmap's rows are empty (%w)", ErrNoData)
	}
	return lines, width, nil
}

func mapLines2Cells[T any](lines []string, width int, transform func(r rune) (T, bool)) ([]T, error) {
	size := width * len(lines)
	cells := make([]T, size)
	for y, line := range lines {
		for x, glyph := range []rune(line) {
			val, ok := transform(glyph)
			if !ok {
				return nil, fmt.Errorf("glyph '%c' is undefined (%w)", glyph, ErrUndefinedAlias)
			}
			cells[y*width+x] = val
		}
	}
	return cells, nil
}

func indexOf(c Coord, w int) int {
	return c.Y*w + c.X
}

func ParseStrMap[T comparable](strmap string, transform func(r rune) (T, bool)) (*TileMap[T], error) {
	lines, width, err := normaliseStrMap(strmap)
	if err != nil {
		return nil, err
	}
	cells, err := mapLines2Cells(lines, width, transform)
	if err != nil {
		return nil, err
	}
	return NewTileMap(width, len(lines), cells)
}

func (tm *TileMap[T]) IndexOf(c Coord) int {
	return indexOf(c, tm.Rect.Size.X)
}

func (tm *TileMap[T]) At(c Coord) (*T, error) {
	if !tm.Rect.InBounds(c) {
		return nil, ErrOutOfBounds
	}
	idx := tm.IndexOf(c)
	return &tm.Cells[idx], nil
}

func (tm *TileMap[T]) Put(c Coord, cell T) error {
	if !tm.Rect.InBounds(c) {
		return ErrOutOfBounds
	}
	idx := tm.IndexOf(c)
	tm.Cells[idx] = cell
	return nil
}

func (tm *TileMap[T]) Stringify(transform func(cell T, c Coord) string) (string, error) {
	var sb strings.Builder
	for y := 0; y < tm.Rect.Size.Y; y++ {
		for x := 0; x < tm.Rect.Size.X; x++ {
			c := Coord{x, y}
			cell, err := tm.At(c)
			if err != nil {
				return "", err
			}
			repr := transform(*cell, c)
			if repr == "" {
				return "", ErrUndefinedAlias
			}
			sb.WriteString(repr)
		}
		sb.WriteRune('\n')
	}
	return sb.String(), nil
}
