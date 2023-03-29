package main

import (
	"os/exec"
	"time"

	"github.com/ebarped/game-of-life/pkg/game"
)

const (
	BOARD_WIDTH     = 5
	BOARD_HEIGHT    = 5
	UPDATE_INTERVAL = 500 * time.Millisecond
)

func main() {
	// input keystroke without pressing enter
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	// do not display entered characters on the screen
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	g := game.New(BOARD_WIDTH, BOARD_HEIGHT)
	g.Init()
	g.Play(UPDATE_INTERVAL)
}
