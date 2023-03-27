package board

import (
	"fmt"

	"github.com/ebarped/game-of-life/pkg/cell"
	"github.com/ebarped/game-of-life/pkg/point"
)

type Board struct {
	width, height int
	cells         [][]cell.Cell
	// we should use a structure to store the cells that allows getting a cell
	// by its position fast! with this we have to loop over all the cells to find the cells with certain conditions..
	// maybe a map[Point]Cell?
	// 	- we can locate a cell by its point
	//  - we are faster (map created in compile time and only accessed and modified in runtime)
	//  -
}

func New(width, height int) Board {
	cells := make([][]cell.Cell, height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			p := point.New(x, y)
			c := cell.New(p)
			cells[x] = append(cells[x], c)
		}
	}
	// select initial cell as starting point of the terminal cursor
	cells[0][0].SetSelected(true)

	return Board{width: width, height: height, cells: cells}
}

//func (b Board) Render() {
//	for i := 0; i < b.width; i++ {
//		fmt.Printf(" _")
//	}
//	fmt.Println()
//	for i := 0; i < b.height; i++ {
//		fmt.Printf("|")
//		for j := 0; j < b.width; j++ {
//			fmt.Printf("_|")
//		}
//		fmt.Println()
//	}
//}

func (b Board) Render() {
	for _, cells := range b.cells {
		for _, cell := range cells {
			fmt.Printf("%s ", cell)
		}
		fmt.Println()
	}
}
