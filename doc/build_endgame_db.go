package main

import (
	"fmt"
	"github.com/ujanssen/emil"
)

var err error

func main() {
	fmt.Printf("Generating all possible positions for KRK\n")

	emil.DEBUG = false
	positions := 0
	for wk := emil.A1; wk <= emil.H8; wk++ {
		for wr := emil.A1; wr <= emil.H8; wr++ {
			for bk := emil.A1; bk <= emil.H8; bk++ {

				board := emil.NewBoard()

				err = board.Setup(emil.WhiteKing, wk)
				if err != nil {
					continue
				}

				err = board.Setup(emil.WhiteRock, wr)
				if err != nil {
					continue
				}

				err = board.Setup(emil.BlackKing, bk)
				if err != nil {
					continue
				}

				positions++
			}
		}
	}
	fmt.Printf("positions %d\n", positions)

}
