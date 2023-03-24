package cell

import "github.com/ebarped/game-of-life/pkg/point"

type Cell struct {
	position   point.Point
	isAlive    bool
	neighbours int
}

func New(p point.Point, isAlive bool) Cell {
	return Cell{position: p, isAlive: isAlive}
}

func (c Cell) IsAlive() bool {
	return c.isAlive
}
