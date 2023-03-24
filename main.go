package main

import (
	"github.com/ebarped/game-of-life/pkg/board"
)

const (
	BOARD_WIDTH  = 6
	BOARD_HEIGHT = 6
)

func main() {
	b := board.New(BOARD_WIDTH, BOARD_HEIGHT)
	b.Render()
}
