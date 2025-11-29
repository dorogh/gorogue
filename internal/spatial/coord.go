package spatial

import (
	"math"
	"rogue/internal/util"
)

type Coord struct {
	X, Y int
}

func XY(x, y int) Coord {
	return Coord{X: x, Y: y}
}

func (c Coord) Add(o Coord) Coord {
	return Coord{
		c.X + o.X,
		c.Y + o.Y,
	}
}

func (c Coord) Sub(o Coord) Coord {
	return Coord{
		c.X - o.X,
		c.Y - o.Y,
	}
}

func (c Coord) Right() Coord {
	return Coord{c.X + 1, c.Y}
}

func (c Coord) Left() Coord {
	return Coord{c.X - 1, c.Y}
}

func (c Coord) Down() Coord {
	return Coord{c.X, c.Y + 1}
}

func (c Coord) Up() Coord {
	return Coord{c.X, c.Y - 1}
}

func (c Coord) Mag() float64 {
	return math.Hypot(float64(c.X), float64(c.Y))
}

func (c Coord) ManhattanDistance(o Coord) int {
	return util.AbsInt(c.X-o.X) + util.AbsInt(c.Y-o.Y)

}

func (c Coord) ChebyshevDistance(o Coord) int {
	x := util.AbsInt(c.X - o.X)
	y := util.AbsInt(c.Y - o.Y)
	if x > y {
		return x
	}
	return y
}
