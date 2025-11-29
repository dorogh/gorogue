package spatial

import "errors"

type Rect struct {
	Size   Coord
	Origin Coord
}

var ErrEndReached = errors.New("end reached")

func NewRectFromZero(size Coord) Rect {
	return Rect{
		Size: size, Origin: Coord{0, 0},
	}
}

func (r Rect) TopLeft() Coord {
	return r.Origin
}

// BottomRight - The result is considered outside the bounds
func (r Rect) BottomRight() Coord {
	return r.Origin.Add(r.Size)
}

func (r Rect) InBounds(c Coord) bool {
	tl := r.Origin
	if c.X < tl.X || c.Y < tl.Y {
		return false
	}
	br := r.BottomRight()
	if c.X >= br.X || c.Y >= br.Y {
		return false
	}
	return true
}

func (r Rect) Area() int {
	return r.Size.X * r.Size.Y
}

// Next - returns the next coord when going row by row.
// returns ErrEndReached if we have reached the bottom right
func (r Rect) Next(c Coord) (Coord, error) {
	if !r.InBounds(c) {
		return c, ErrOutOfBounds
	}
	br := r.BottomRight()
	if c.X+1 < br.X {
		return Coord{c.X + 1, c.Y}, nil
	}
	if c.Y+1 < br.Y {
		return Coord{r.Origin.X, c.Y + 1}, nil
	}
	return c, ErrEndReached
}
