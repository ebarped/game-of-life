package cell

import (
	"github.com/ebarped/game-of-life/pkg/point"
)

const (
	deadCellRepr     = "□"
	aliveCellRepr    = "▣"
	selectedCellRepr = "▢"

	colorReset = "\033[0m"  // reset color
	colorRed   = "\033[31m" // cell alive
	colorGreen = "\033[32m" // cell selected
)

type Cell struct {
	position   point.Point
	isAlive    bool
	neighbours int
	isSelected bool
}

func New(p point.Point) Cell {
	return Cell{position: p}
}

func (c Cell) IsAlive() bool {
	return c.isAlive
}

func (c Cell) String() string {
	if c.isAlive {
		return string(colorRed) + aliveCellRepr + string(colorReset)
	} else if c.isSelected {
		return string(colorGreen) + selectedCellRepr + string(colorReset)
	}
	return deadCellRepr
}

func (c *Cell) SetSelected(selected bool) {
	c.isSelected = selected
}
