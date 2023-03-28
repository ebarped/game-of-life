package board

import (
	"fmt"

	"github.com/ebarped/game-of-life/pkg/cell"
	"github.com/ebarped/game-of-life/pkg/point"
)

type Board struct {
	width, height int
	cells         map[point.Point]*cell.Cell
}

func New(width, height int) Board {
	cells := make(map[point.Point]*cell.Cell)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			p := point.New(x, y)
			c := cell.New(p)
			cells[p] = &c
		}
	}
	// select initial cell as starting point of the terminal cursor
	cells[point.New(0, 0)].SetSelected(true)

	fmt.Printf("%#v\n", cells)
	return Board{width: width, height: height, cells: cells}
}

// Render prints the game board
func (b Board) Render() {
	for y := 0; y < b.height; y++ {
		for x := 0; x < b.width; x++ {
			p := point.New(x, y)
			fmt.Printf("%s ", b.cells[p])
		}
		fmt.Println()
	}
}

// GetCell returns a pointer to a cell
// implemented in order the cell field of Board can remain abstract to the clients
func (b Board) GetCell(p point.Point) *cell.Cell {
	return b.cells[p]
}

// IsInside returns true if the point is inside the the board
func (b Board) IsInside(p point.Point) bool {
	// check horizontal
	if 0 > p.X || p.X > b.width-1 {
		return false
	}
	// check vertical
	if 0 > p.Y || p.Y > b.height-1 {
		return false
	}
	return true
}

// GetWidth returns the width of the board
func (b Board) GetWidth() int {
	return b.width
}
