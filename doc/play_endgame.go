package main

import (
	"fmt"
	"github.com/ujanssen/emil"
)

func main() {
	emil.DEBUG = true

	board := emil.NewBoard()
	board.Setup(emil.WhiteKing, emil.H1)
	board.Setup(emil.BlackKing, emil.A8)
	board.Setup(emil.WhiteRock, emil.G2)

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
