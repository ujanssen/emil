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

	db, _ := emil.LoadEndGameDb()

	for {
		move := db.Find(emil.NewPosition(board, emil.WHITE))
		if move == nil {
			break
		}
		fmt.Printf("move %s\n", move)
		board = board.DoMove(move)
		fmt.Printf("%s\n", board.Picture())

		move = emil.Search(emil.NewPosition(board, emil.BLACK))
		if move == nil {
			break
		}
		fmt.Printf("move %s\n", move)
		board = board.DoMove(move)
		fmt.Printf("%s\n", board.Picture())
	}
}
