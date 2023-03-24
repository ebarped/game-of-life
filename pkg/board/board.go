package board

import (
	"fmt"

	"github.com/ebarped/game-of-life/pkg/cell"
	"github.com/ebarped/game-of-life/pkg/point"
)

const (
	EmptyCellRepr = "□"
	FullCellRepr  = "▣"
)

type Board struct {
	width, height int
	cells         [][]cell.Cell
}

func New(width, height int) Board {
	cells := make([][]cell.Cell, height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			p := point.New(x, y)
			c := cell.New(p, true)
			cells[x] = append(cells[x], c)
		}
	}

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
			if cell.IsAlive() {
				fmt.Printf("%s ", FullCellRepr)
			} else {
				fmt.Printf("%s ", EmptyCellRepr)
			}
		}
		fmt.Println()
	}
}
