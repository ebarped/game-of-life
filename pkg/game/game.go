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
	board board.Board
}

func New(width, height int) game {
	return game{board: board.New(width, height)}
}

// Start starts the game, showing the main menu
func (g game) Start(updateInterval time.Duration) {
	clearScreen()

	// get input while game is running
	userInput := make(chan byte)
	go handleInput(userInput)

	// get the point of the currently selected cell, initially its (0,0)
	selectedCellPoint := point.New(0, 0)

	// menu loop
	for {
		g.board.Render()
		g.displayMenuInstructions()

		keyPressed := <-userInput

		switch keyPressed {
		case ESCAPE_CHAR:
			_ = <-userInput // skip the ] char
			input := <-userInput
			switch input {
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
				fmt.Println("error: key not recognized:", input)
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

			g.play(updateInterval, userInput)
			clearScreen()

			// restart game
			g.board = board.New(g.board.Width(), g.board.Hight())
			selectedCellPoint = point.New(0, 0)
			//
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

// handleInput is intented to run as goroutine to catch the user input (pause or restart)
func handleInput(input chan<- byte) {
	for {
		keyPressed := make([]byte, 1)
		os.Stdin.Read(keyPressed)
		input <- keyPressed[0]
	}
}

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

func (g game) displayGameInstructions() {
	for i := 0; i < g.board.Width()*2; i++ {
		fmt.Printf("-")
	}
	fmt.Println()
	fmt.Println("Press <p> to pause the game")
	fmt.Println("Press <r> to restart the game")
	fmt.Println("Press <q> to quit the game")
}

func clearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}
