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
	PAUSE_CHAR       byte = 112
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

	for {
		g.board.Render()
		g.displayInitInstructions()

		keyPressed := readInput()
		switch keyPressed {
		case ESCAPE_CHAR:
			_ = readInput() // skip the ] char
			arrow := readInput()
			switch arrow {
			case ARROW_UP_CHAR:
				newSelectedPoint := selectedCellPoint.North()
				if !g.board.IsInside(newSelectedPoint) {
					break
				}
				c := g.board.GetCell(selectedCellPoint)
				c.SetSelected(false)
				g.board.SetCell(c.Position(), c)
				selectedCellPoint = newSelectedPoint
				c = g.board.GetCell(selectedCellPoint)
				c.SetSelected(true)
				g.board.SetCell(c.Position(), c)
			case ARROW_DOWN_CHAR:
				newSelectedPoint := selectedCellPoint.South()
				if !g.board.IsInside(newSelectedPoint) {
					break
				}
				c := g.board.GetCell(selectedCellPoint)
				c.SetSelected(false)
				g.board.SetCell(c.Position(), c)
				selectedCellPoint = newSelectedPoint
				c = g.board.GetCell(selectedCellPoint)
				c.SetSelected(true)
				g.board.SetCell(c.Position(), c)
			case ARROW_RIGHT_CHAR:
				newSelectedPoint := selectedCellPoint.East()
				if !g.board.IsInside(newSelectedPoint) {
					break
				}
				c := g.board.GetCell(selectedCellPoint)
				c.SetSelected(false)
				g.board.SetCell(c.Position(), c)
				selectedCellPoint = newSelectedPoint
				c = g.board.GetCell(selectedCellPoint)
				c.SetSelected(true)
				g.board.SetCell(c.Position(), c)
			case ARROW_LEFT_CHAR:
				newSelectedPoint := selectedCellPoint.West()
				if !g.board.IsInside(newSelectedPoint) {
					break
				}
				c := g.board.GetCell(selectedCellPoint)
				c.SetSelected(false)
				g.board.SetCell(c.Position(), c)
				selectedCellPoint = newSelectedPoint
				c = g.board.GetCell(selectedCellPoint)
				c.SetSelected(true)
				g.board.SetCell(c.Position(), c)
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
			g.board.SetCell(c.Position(), c)
			clearScreen()
		case ENTER_CHAR:
			c := g.board.GetCell(selectedCellPoint)
			c.SetSelected(false)
			g.board.SetCell(c.Position(), c)
			fmt.Println("STARTING THE GAME!")
			return
		default:
			clearScreen()
			fmt.Printf("Unrecognized key, skipping it: %s (%d)\n", string(keyPressed), keyPressed)
		}
	}
}

// Play starts the loop of the game
func (g game) Play(updateInterval time.Duration) {
	runGame := true
	i := 0

	// get input while game is running
	userInput := make(chan byte, 1)
	go handlePause(userInput)

	clearScreen()

	fmt.Println("iteration:", i)
	fmt.Println("---------------")
	g.board.Render()
	g.displayInstructions()

	for {
		select {
		case <-userInput:
			runGame = !runGame // flip rungame state
		case <-time.After(updateInterval):
			if runGame {
				i++
				clearScreen()
				fmt.Println("iteration:", i)
				fmt.Println("---------------")

				g.board = g.board.Update()
				g.board.Render()
				g.displayInstructions()
			}
		}
	}
}

// handlePause is intented to run as goroutine to catch the user pause input
func handlePause(input chan<- byte) {
	for {
		keyPressed := make([]byte, 1)
		os.Stdin.Read(keyPressed)
		if keyPressed[0] == PAUSE_CHAR {
			input <- keyPressed[0]
		}
	}
}

func readInput() byte {
	keyPressed := make([]byte, 1)
	os.Stdin.Read(keyPressed)
	return keyPressed[0]
}

func (g game) displayInitInstructions() {
	for i := 0; i < g.board.GetWidth()*2; i++ {
		fmt.Printf("-")
	}
	fmt.Println()
	fmt.Println("Use <ARROW> keys to move through the board")
	fmt.Println("Use <SPACEBAR> key to set cells to ALIVE STATUS")
	fmt.Println("Use <ENTER> key to start the game")
}

func (g game) displayInstructions() {
	for i := 0; i < g.board.GetWidth()*2; i++ {
		fmt.Printf("-")
	}
	fmt.Println()
	fmt.Println("Press <p> to pause the game")
}

func clearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}
