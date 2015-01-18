package main

import (
	"fmt"
	"github.com/ujanssen/emil"
)

func main() {
	emil.DEBUG = true

	board := emil.NewBoard()
	board.Setup(emil.WhiteKing, emil.B5)
	board.Setup(emil.BlackKing, emil.A7)
	board.Setup(emil.WhiteRock, emil.H2)

	db := emil.NewEndGameDb()

	for {
		move := db.Find(board)
		if move == nil {
			break
		}
		fmt.Printf("move %s\n", move)
		board = board.DoMove(move)
	}
}
