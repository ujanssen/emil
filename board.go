package emil

import (
	"fmt"
)

// Board with an array of field values, representing pieces
type Board struct {
	squares []int
}

// NewBoard creates a new Board
func NewBoard() *Board {
	return &Board{squares: make([]int, SQUARES)}
}
func (b *Board) String() string {
	files := "a b c d e f g h "
	s := "  " + files + " \n"
	for _, r := range FirstSquares {
		s += fmt.Sprintf("%d ", BoardSquares[r].rank)
		for f := 0; f < 8; f++ {
			switch b.squares[r+f] {
			case WhiteKing:
				s += "K "
			case BlackKing:
				s += "k "
			case WhiteRock:
				s += "R "
			case BlackRock:
				s += "r "
			default:
				s += "  "
			}
		}
		s += fmt.Sprintf("%d\n", BoardSquares[r].rank)
	}
	s += "  " + files + " \n"
	return s
}

//Setup a piece on a square
func (b *Board) Setup(piece, square int) {
	b.squares[square] = piece
}

//Square of a chess board
type Square struct {
	name  string
	file  string
	rank  int
	index int
}

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
	// define BoardSquares
	for j := 0; j < 8; j++ {
		for i, file := range FILES {
			index := i + j*8
			BoardSquares[index] = newSquare(string(file), j+1, index)
		}
	}
}
