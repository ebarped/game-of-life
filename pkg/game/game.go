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
	RESTART_CHAR     byte = 114
	QUIT_CHAR        byte = 113
)

const (
	colorReset = "\x1b[0m"  // reset color
	colorRed   = "\x1b[31m" // red color
	blink      = "\x1b[5m"  // blink effect
)

type game struct {
	board             board.Board
	userInput         chan byte
	selectedCellPoint point.Point
}

func New(width, height int) game {
	return game{
		board:             board.New(width, height),
		userInput:         make(chan byte),
		selectedCellPoint: point.New(0, 0),
	}
}

// Start starts the game, showing the main menu
func (g game) Start(updateInterval time.Duration) {
	clearScreen()

	// get input while game is running
	go g.handleInput()

	// menu loop
	for {
		g.board.Render()
		g.displayMenuInstructions()

		keyPressed := <-g.userInput

		switch keyPressed {
		case ESCAPE_CHAR:
			_ = <-g.userInput // skip the ] char
			input := <-g.userInput
			switch input {
			case ARROW_UP_CHAR:
				g.moveSelectedCell(board.North)
			case ARROW_DOWN_CHAR:
				g.moveSelectedCell(board.South)
			case ARROW_RIGHT_CHAR:
				g.moveSelectedCell(board.East)
			case ARROW_LEFT_CHAR:
				g.moveSelectedCell(board.West)
			default:
				fmt.Println("error: key not recognized:", input)
			}
			clearScreen()
		case SPACEBAR_CHAR:
			c := g.board.GetCell(g.selectedCellPoint)
			if c.IsAlive() {
				c.SetAlive(false)
			} else {
				c.SetAlive(true)
			}
			g.board.SetCell(c.Position(), c)
			clearScreen()
		case ENTER_CHAR:
			g.changeCellSelectStatus(g.selectedCellPoint, false)

			g.play(updateInterval, g.userInput)

			clearScreen()
			g.restart()
		case QUIT_CHAR:
			fmt.Println("Bye!")
			os.Exit(0)
		default:
			clearScreen()
			fmt.Printf("Unrecognized key, skipping it: %s (%d)\n", string(keyPressed), keyPressed)
		}
	}
}

// play starts the loop of the game
func (g game) play(updateInterval time.Duration, userInput chan byte) {
	runGame := true
	i := 0

	clearScreen()

	fmt.Println("iteration:", i)
	fmt.Println("---------------")
	g.board.Render()
	g.displayGameInstructions()

	for {
		select {
		case input := <-userInput:
			switch input {
			case PAUSE_CHAR:
				runGame = !runGame // flip game state

				// when resuming the game, render immediately, dont wait "updateInterval" until next cycle
				if runGame {
					i++
					clearScreen()
					fmt.Println("iteration:", i)
					fmt.Println("---------------")

					g.board = g.board.Update()
					g.board.Render()
					g.displayGameInstructions()
				} else {
					fmt.Println(string(blink) + string(colorRed) + "GAME PAUSED" + string(colorReset))
				}
			case RESTART_CHAR:
				return
			case QUIT_CHAR:
				fmt.Println("Bye!")
				os.Exit(0)
			}

		case <-time.After(updateInterval):
			if runGame {
				i++
				clearScreen()
				fmt.Println("iteration:", i)
				fmt.Println("---------------")

				g.board = g.board.Update()
				g.board.Render()
				g.displayGameInstructions()
			}
		}
	}
}

// handleInput is intented to run as goroutine to catch the user input
func (g game) handleInput() {
	for {
		keyPressed := make([]byte, 1)
		os.Stdin.Read(keyPressed)
		g.userInput <- keyPressed[0]
	}
}

// displayMenuInstructions displays the instructions of the menu
func (g game) displayMenuInstructions() {
	for i := 0; i < g.board.Width()*2; i++ {
		fmt.Printf("-")
	}
	fmt.Println()
	fmt.Println("Press <ARROW> keys to move through the board")
	fmt.Println("Press <SPACEBAR> key to set cells to ALIVE STATUS")
	fmt.Println("Press <ENTER> key to start the game")
	fmt.Println("Press <q> to quit the game")
}

// displayGameInstructions displays the instructions of the running game
func (g game) displayGameInstructions() {
	for i := 0; i < g.board.Width()*2; i++ {
		fmt.Printf("-")
	}
	fmt.Println()
	fmt.Println("Press <p> to pause the game")
	fmt.Println("Press <r> to restart the game")
	fmt.Println("Press <q> to quit the game")
}

// clearScreen clears the screen (unix terminals only)
func clearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

// moveSelectedCell moves the selected cell in the "dir" direction
func (g *game) moveSelectedCell(dir board.Direction) {
	newSelectedPoint := g.board.CalculatePosition(g.selectedCellPoint, dir)
	g.changeCellSelectStatus(g.selectedCellPoint, false)
	g.selectedCellPoint = newSelectedPoint
	g.changeCellSelectStatus(g.selectedCellPoint, true)
}

// restart restarts the initial conditions of the game
func (g *game) restart() {
	g.board = board.New(g.board.Width(), g.board.Height())
	g.selectedCellPoint = point.New(0, 0)
}

// changeCellSelectStatus changes the selected status of the cell in position "p" to the "status"
func (g *game) changeCellSelectStatus(p point.Point, status bool) {
	c := g.board.GetCell(g.selectedCellPoint)
	c.SetSelected(status)
	g.board.SetCell(c.Position(), c)
}
