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
	fmt.Printf("\n")
	fmt.Printf("%s\n", board)

	fmt.Printf("K F6: %s\n", board.Moves(emil.WhiteKing, emil.F6))
	fmt.Printf("R B1: %s\n", board.Moves(emil.WhiteRock, emil.B1))
	fmt.Printf("k G8: %s\n", board.Moves(emil.BlackKing, emil.G8))
}
