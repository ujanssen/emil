package main

import (
	"fmt"
	"github.com/ujanssen/emil"
	"time"
)

func main() {
	emil.IN_TEST = !true
	emil.DEBUG = true

	start := time.Now()
	emil.LoadEndGameDb()
	end := time.Now()
	fmt.Printf("\n\n\nload duration %v\n", end.Sub(start))
	/*
		if err == nil {

			fmt.Printf("len(db.AnalysisMap) %v\n\n", len(db.AnalysisMap))

			board := emil.NewBoard()
			board.Setup(emil.WhiteKing, emil.E6)
			board.Setup(emil.BlackKing, emil.E8)
			board.Setup(emil.WhiteRock, emil.H1)

			want := "Rh1h8"
			move := db.Find(emil.NewPosition(board, emil.WHITE))
			if move == nil {
				fmt.Printf("the move should be %s, got nil", want)
			} else {
				got := move.String()
				if got != want {
					fmt.Printf("the move should be %s, got %s", want, got)
				}
			}
		}
	*/
}
