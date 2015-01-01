package emil

import (
	"errors"
	"fmt"
	"strings"
)

var errNotEmpty = errors.New("not empty")
var errKingsToClose = errors.New("Kings to close")

// Board with an array of field values, representing pieces
type Board struct {
	Squares   []int `json:"squares"`
	whiteKing int
	blackKing int

	str string
}

// NewBoard creates a new Board
func NewBoard() *Board {
	return &Board{Squares: make([]int, SQUARES)}
}
func (b *Board) String() string {
	if len(b.str) > 0 {
		return b.str
	}
	s := ""
	for _, r := range FirstSquares {
		if len(s) > 0 {
			s += "/"
		}
		for f := 0; f < 8; f++ {
			s += fmt.Sprintf("%s", symbol(b.Squares[r+f]))
		}
	}
	s = strings.Replace(s, "        ", "8", -1)
	s = strings.Replace(s, "       ", "7", -1)
	s = strings.Replace(s, "      ", "6", -1)
	s = strings.Replace(s, "     ", "5", -1)
	s = strings.Replace(s, "    ", "4", -1)
	s = strings.Replace(s, "   ", "3", -1)
	s = strings.Replace(s, "  ", "2", -1)
	s = strings.Replace(s, " ", "1", -1)
	b.str = s
	return s
}
func (b *Board) Picture() string {
	files := "a b c d e f g h "
	s := "  " + files + " \n"
	for _, r := range FirstSquares {
		s += fmt.Sprintf("%d ", BoardSquares[r].rank)
		for f := 0; f < 8; f++ {
			s += fmt.Sprintf("%s ", symbol(b.Squares[r+f]))
		}
		s += fmt.Sprintf("%d\n", BoardSquares[r].rank)
	}
	s += "  " + files + " \n"
	return s
}

//Setup a piece on a square
func (b *Board) Setup(piece, square int) (noError error) {
	if b.Squares[square] != Empty {
		return errNotEmpty
	}
	b.Squares[square] = piece
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
	copy(newBoard.Squares, b.Squares)
	newBoard.Squares[m.source] = Empty
	newBoard.Squares[m.destination] = m.piece
	// if DEBUG {
	// 	fmt.Printf("%s\n", b)
	// }
	return newBoard
}
