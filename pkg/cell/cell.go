package cell

import (
	"github.com/ebarped/game-of-life/pkg/point"
)

const (
	deadCellRepr             = "□"
	aliveCellRepr            = "▣"
	selectedCellRepr         = "▢"
	selectedAndAliveCellRepr = "▣"

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

// IsAlive returns true if the cell is alive
func (c Cell) IsAlive() bool {
	return c.isAlive
}

// SetSelected sets the cell as selected by the UI
func (c *Cell) SetSelected(selected bool) {
	c.isSelected = selected
}

// SetAlive sets the cell as alive by the UI
func (c *Cell) SetAlive(alive bool) {
	c.isAlive = alive
}

// MoveTo changes c position to p
func (c *Cell) MoveTo(p point.Point) {
	c.position = p
}

// String allows pretty printing of the Cell
func (c Cell) String() string {
	if c.isAlive && c.isSelected {
		return string(colorGreen) + selectedAndAliveCellRepr + string(colorReset)
	} else if c.isAlive {
		return string(colorRed) + aliveCellRepr + string(colorReset)
	} else if c.isSelected {
		return string(colorGreen) + selectedCellRepr + string(colorReset)
	}
	return deadCellRepr
}
