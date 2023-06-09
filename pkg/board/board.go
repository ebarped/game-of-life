package board

import (
	"fmt"

	"github.com/ebarped/game-of-life/pkg/cell"
	"github.com/ebarped/game-of-life/pkg/point"
)

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

type Board struct {
	width  int
	height int
	cells  map[point.Point]cell.Cell
}

func New(width, height int) Board {
	cells := make(map[point.Point]cell.Cell)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			p := point.New(x, y)
			c := cell.New(p)
			cells[p] = c
		}
	}
	// select initial cell as starting point of the terminal cursor
	initialPos := point.New(0, 0)
	c := cells[initialPos]
	c.SetSelected(true)
	cells[point.New(0, 0)] = c

	return Board{width: width, height: height, cells: cells}
}

// Clone returns a clone of the board b
func (b Board) Clone() Board {
	clon := Board{
		width:  b.width,
		height: b.height,
		cells:  make(map[point.Point]cell.Cell),
	}

	for k, v := range b.cells {
		clon.cells[k] = v
	}
	return clon
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

// Update will calculate and return a new board based on the actual board
func (b Board) Update() Board {
	// create a copy of the board
	newBoard := b.Clone()

	// update the new board based on the state of the old board
	for _, c := range b.cells {
		newCell := c

		// handle alive cell
		if c.IsAlive() {
			switch b.NeighboursCount(c) {
			case 0, 1, 4, 5, 6, 7, 8:
				newCell.SetAlive(false)
				newBoard.SetCell(newCell.Position(), newCell)
			// this case can be ommited, if the cell is alve and do not have 0 1 4 ... 8 neighbours, it has to have 2 or 3
			case 2, 3:
				newCell.SetAlive(true)
				newBoard.SetCell(newCell.Position(), newCell)
			}
		} else {
			// handle dead cell
			if b.NeighboursCount(c) == 3 {
				newCell.SetAlive(true)
				newBoard.SetCell(newCell.Position(), newCell)
			}
		}
	}

	return newBoard
}

// SetCell updates the cell in point p with cell c
func (b Board) SetCell(p point.Point, c cell.Cell) {
	b.cells[p] = c
}

// GetCell returns a pointer to a cell
// THIS SHOULD CHECK EXISTENCE OF THE CELL?
func (b Board) GetCell(p point.Point) cell.Cell {
	return b.cells[p]
}

// IsInside returns true if the point is inside the the board
func (b Board) IsInside(p point.Point) bool {
	// check horizontal
	if 0 > p.GetX() || p.GetX() > b.width-1 {
		return false
	}
	// check vertical
	if 0 > p.GetY() || p.GetY() > b.height-1 {
		return false
	}
	return true
}

// NeighboursCount returns the number of alive neighbours of a cell
func (b Board) NeighboursCount(c cell.Cell) int {
	count := 0

	// array of positions to check
	neighbourPositions := []point.Point{
		b.CalculatePosition(c.Position(), North),
		b.CalculatePosition(c.Position(), North, East),
		b.CalculatePosition(c.Position(), East),
		b.CalculatePosition(c.Position(), South, East),
		b.CalculatePosition(c.Position(), South),
		b.CalculatePosition(c.Position(), South, West),
		b.CalculatePosition(c.Position(), West),
		b.CalculatePosition(c.Position(), North, West),
	}

	for _, pos := range neighbourPositions {
		if b.GetCell(pos).IsAlive() {
			count++
		}
	}

	return count
}

// Width returns the width of the board
func (b Board) Width() int {
	return b.width
}

// Height returns the height of the board
func (b Board) Height() int {
	return b.height
}

// CalculatePosition returns the final position of applying all the "dir"
// Directions to the "initialPos" position, accounting that the board is infinite
func (b Board) CalculatePosition(initialPos point.Point, dirs ...Direction) point.Point {
	actualPos := initialPos

	for _, dir := range dirs {
		switch dir {
		case North:
			actualPos = actualPos.North()
			if !b.IsInside(actualPos) {
				actualPos = point.New(actualPos.GetX(), b.height-1)
			}
		case East:
			actualPos = actualPos.East()
			if !b.IsInside(actualPos) {
				actualPos = point.New(0, actualPos.GetY())
			}
		case South:
			actualPos = actualPos.South()
			if !b.IsInside(actualPos) {
				actualPos = point.New(actualPos.GetX(), 0)
			}
		case West:
			actualPos = actualPos.West()
			if !b.IsInside(actualPos) {
				actualPos = point.New(b.width-1, actualPos.GetY())
			}
		default:
			panic("this direction is not contemplated!")
		}
	}

	return actualPos
}
