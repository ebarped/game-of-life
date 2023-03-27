package game

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/ebarped/game-of-life/pkg/board"
)

const (
	ENTER_CHAR    byte = 10
	ESCAPE_CHAR   byte = 27
	SPACEBAR_CHAR byte = 32
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

	// get the point of the initially selected cell
	// on each arrow press, change the selected cell
	// on spacebar, set that cell indexed by its point to alive/dead
	// on enter, start the game

	for {
		g.board.Render()
		displayInstructions()

		keyPressed := readInput()
		switch keyPressed {
		case ESCAPE_CHAR:
			_ = readInput() // skip the ] char
			arrow := readInput()
			fmt.Println("Pressed the arrow:", string(arrow))
			clearScreen()
		case SPACEBAR_CHAR:
			clearScreen()
			fmt.Println("Setting cell status to alive!")
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

func displayInstructions() {
	fmt.Println("---------------")
	fmt.Println("Use the arrows to move through the board")
	fmt.Println("Use <SPACEBAR> to set cells to ALIVE STATUS")
	fmt.Println("Use <ENTER> to start the game")
}

func clearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}
