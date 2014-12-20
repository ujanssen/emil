package emil

import (
	"fmt"
)

//Square of a chess board
type Square struct {
	name  string
	file  string
	rank  int
	index int
}

//SQUARES: the number of quares
const SQUARES = 64

//FILES: the files of a board
const FILES = "abcdefgh"

//BoardSquares the squares of the board
var BoardSquares [SQUARES]*Square

func (s *Square) String() string {
	return s.name
}

func newSquare(file string, rank, index int) *Square {
	name := fmt.Sprintf("%v%d", file, rank)

	return &Square{
		name:  name,
		file:  file,
		rank:  rank,
		index: index}
}

func init() {
	// define boardSquares
	for j := 0; j < 8; j++ {
		for i, file := range FILES {
			index := i + j*8
			BoardSquares[index] = newSquare(string(file), j+1, index)
		}
	}
}
