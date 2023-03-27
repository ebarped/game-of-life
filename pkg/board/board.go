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

func (b Board) Render() {
	for y := 0; y < b.height; y++ {
		for x := 0; x < b.width; x++ {
			p := point.New(x, y)
			fmt.Printf("%s ", b.cells[p])
		}
		fmt.Println()
	}
}
