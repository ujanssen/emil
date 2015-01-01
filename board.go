package emil

import (
	"errors"
	"fmt"
)

var errNotEmpty = errors.New("not empty")
var errKingsToClose = errors.New("Kings to close")

// Board with an array of field values, representing pieces
type Board struct {
	squares   []int `json:"squares"`
	whiteKing int   `json:"whiteKing"`
	blackKing int   `json:"blackKing"`

	str string `json:"str"`
}

// NewBoard creates a new Board
func NewBoard() *Board {
	return &Board{squares: make([]int, SQUARES)}
}
func (b *Board) String() string {
	if len(b.str) > 0 {
		return b.str
	}
	files := "a b c d e f g h "
	s := "  " + files + " \n"
	for _, r := range FirstSquares {
		s += fmt.Sprintf("%d ", BoardSquares[r].rank)
		for f := 0; f < 8; f++ {
			s += fmt.Sprintf("%s ", symbol(b.squares[r+f]))
		}
		s += fmt.Sprintf("%d\n", BoardSquares[r].rank)
	}
	s += "  " + files + " \n"
	b.str = s
	return s
}

//Setup a piece on a square
func (b *Board) Setup(piece, square int) (noError error) {
	if b.squares[square] != Empty {
		return errNotEmpty
	}
	b.squares[square] = piece
	if piece == BlackKing {
		b.blackKing = square
	}
	if piece == WhiteKing {
		b.whiteKing = square
	}
	return noError
}

func (b *Board) kingsToClose() (noError error) {
	if squaresDistances[b.whiteKing][b.blackKing] <= 1 {
		return errKingsToClose
	}
	return noError
}
func (b *Board) DoMove(m *Move) (newBoard *Board) {
	return b.doMove(m)
}

func (b *Board) doMove(m *Move) (newBoard *Board) {
	// if DEBUG {
	// 	fmt.Printf("do move: %s\n", m)
	// }
	newBoard = NewBoard()
	newBoard.whiteKing = b.whiteKing
	newBoard.blackKing = b.blackKing
	copy(newBoard.squares, b.squares)
	newBoard.squares[m.source] = Empty
	newBoard.squares[m.destination] = m.piece
	// if DEBUG {
	// 	fmt.Printf("%s\n", b)
	// }
	return newBoard
}
