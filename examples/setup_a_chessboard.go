package main

import (
	"fmt"
	"github.com/ujanssen/emil"
)

func main() {
	emil.DEBUG = false

	fmt.Printf("A chess board\n")
	board := emil.NewBoard()
	board.Setup(emil.WhiteKing, emil.B5)
	board.Setup(emil.WhiteRock, emil.H7)

	board.Setup(emil.BlackKing, emil.B8)

	moves := emil.GenerateMoves(emil.NewPosition(board, emil.BLACK))

	fmt.Printf("black:%s\n%s\n", board, moves)

	for _, m := range moves {
		b := board.DoMove(m)
		fmt.Printf("black:%s=%v\n%s\n", m, emil.IsTheKingInCheck(emil.NewPosition(b, emil.WHITE)), b)
	}
}
