package game

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/ebarped/game-of-life/pkg/board"
	"github.com/ebarped/game-of-life/pkg/point"
)

const (
	ENTER_CHAR       byte = 10
	ESCAPE_CHAR      byte = 27
	SPACEBAR_CHAR    byte = 32
	ARROW_UP_CHAR    byte = 65
	ARROW_DOWN_CHAR  byte = 66
	ARROW_RIGHT_CHAR byte = 67
	ARROW_LEFT_CHAR  byte = 68
)

type game struct {
	board board.Board
}

func New(width, height int) game {
	return game{board: board.New(width, height)}
}

// Init lets the user to select the alive cells and the initial conditions before going to play
func (g game) Init() {
	clearScreen()

	// get the point of the currently selected cell, initially its (0,0)
	selectedCellPoint := point.New(0, 0)

	// on each arrow press, change the selected cell
	// on spacebar, set that cell indexed by its point to alive/dead
	// on enter, start the game

	for {
		g.board.Render()
		g.displayInstructions()

		keyPressed := readInput()
		switch keyPressed {
		case ESCAPE_CHAR:
			_ = readInput() // skip the ] char
			arrow := readInput()
			switch arrow {
			case ARROW_UP_CHAR:
				newSelectedPoint := selectedCellPoint.GetNorth()
				if !g.board.IsInside(newSelectedPoint) {
					break
				}
				g.board.GetCell(selectedCellPoint).SetSelected(false)
				selectedCellPoint = newSelectedPoint
				g.board.GetCell(selectedCellPoint).SetSelected(true)

			case ARROW_DOWN_CHAR:
				newSelectedPoint := selectedCellPoint.GetSouth()
				if !g.board.IsInside(newSelectedPoint) {
					break
				}
				g.board.GetCell(selectedCellPoint).SetSelected(false)
				selectedCellPoint = newSelectedPoint
				g.board.GetCell(selectedCellPoint).SetSelected(true)

			case ARROW_RIGHT_CHAR:
				newSelectedPoint := selectedCellPoint.GetEast()
				if !g.board.IsInside(newSelectedPoint) {
					break
				}
				g.board.GetCell(selectedCellPoint).SetSelected(false)
				selectedCellPoint = newSelectedPoint
				g.board.GetCell(selectedCellPoint).SetSelected(true)

			case ARROW_LEFT_CHAR:
				newSelectedPoint := selectedCellPoint.GetWest()
				if !g.board.IsInside(newSelectedPoint) {
					break
				}
				g.board.GetCell(selectedCellPoint).SetSelected(false)
				selectedCellPoint = newSelectedPoint
				g.board.GetCell(selectedCellPoint).SetSelected(true)
			default:
				fmt.Println("error: key not recognized:", arrow)
			}
			clearScreen()
		case SPACEBAR_CHAR:
			c := g.board.GetCell(selectedCellPoint)
			if c.IsAlive() {
				c.SetAlive(false)
			} else {
				c.SetAlive(true)
			}
			clearScreen()
		case ENTER_CHAR:
			fmt.Println("STARTING THE GAME!")
			return
		default:
			clearScreen()
			fmt.Printf("Unrecognized key, skipping it: %s (%d)\n", string(keyPressed), keyPressed)
		}
	}
}

func (g game) Play() {
	fmt.Println("START")
	for {
		i := 0

		g.board.Render()
		time.Sleep(200 * time.Millisecond)
		fmt.Println("iteration:", i)
		i++
		fmt.Println("---------------")
	}
}

func readInput() byte {
	keyPressed := make([]byte, 1)
	os.Stdin.Read(keyPressed)
	return keyPressed[0]
}

func (g game) Update() {

}

func (g game) displayInstructions() {
	for i := 0; i < g.board.GetWidth()*2; i++ {
		fmt.Printf("-")
	}
	fmt.Println()
	fmt.Println("Use <ARROW> keys to move through the board")
	fmt.Println("Use <SPACEBAR> key to set cells to ALIVE STATUS")
	fmt.Println("Use <ENTER> key to start the game")
}

func clearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}
