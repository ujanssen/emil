package main

import (
	"fmt"
	"github.com/ujanssen/emil"
)

func main() {
	emil.DEBUG = false

	fmt.Printf("A chess board\n")
	board := emil.NewBoard()
	board.Setup(emil.WhiteKing, emil.F6)
	board.Setup(emil.BlackKing, emil.G8)
	board.Setup(emil.WhiteRock, emil.B1)
	board.Setup(emil.BlackRock, emil.A1)
	fmt.Printf("\n")
	fmt.Printf("%s\n", board)

	move := emil.Search(board, emil.WHITE)

	fmt.Printf("white: %s\n", move)
}
