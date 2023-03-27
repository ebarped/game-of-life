package main

import (
	"os/exec"

	"github.com/ebarped/game-of-life/pkg/game"
)

const (
	BOARD_WIDTH  = 5
	BOARD_HEIGHT = 5
)

func main() {

	// input keystroke without pressing enter
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	g := game.New(BOARD_WIDTH, BOARD_HEIGHT)
	_ = g
	g.Init()
	//g.Play()
}
