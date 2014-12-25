package main

import (
	"fmt"
	"github.com/ujanssen/emil"
)

func main() {

	board := emil.NewBoard()
	board.Setup(emil.WhiteKing, emil.E6)
	board.Setup(emil.BlackKing, emil.E8)
	board.Setup(emil.WhiteRock, emil.H1)

	db := emil.NewEndGameDb()

	move := db.Find(board)
	fmt.Printf("move should be Rh1h8, got %s", move)

}
