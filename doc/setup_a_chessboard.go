package main

import (
	"fmt"
	"github.com/ujanssen/emil"
)

func main() {
	fmt.Printf("A chess board\n")
	board := emil.NewBoard()
	board.Setup(emil.WhiteKing, emil.F6)
	board.Setup(emil.BlackKing, emil.G8)
	board.Setup(emil.WhiteRock, emil.B1)
	board.Setup(emil.BlackRock, emil.A1)
	fmt.Printf("\n")
	fmt.Printf("%s\n", board)

	// emil.DEBUG = true

	move := emil.Search(board, emil.WHITE)

	fmt.Printf("black: %s\n", move)
}
