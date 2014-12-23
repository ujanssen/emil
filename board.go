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
			s += fmt.Sprintf("%s ", symbol(b.squares[r+f]))
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

func (b *Board) isKingCapturedAfter(m *Move) bool {
	_, kingCaptured := Search(b, otherPlayer(m.player), true)
	return kingCaptured != nil
}

func (b *Board) doMove(m *Move) {
	if DEBUG {
		println("do move", m.String())
	}
	b.squares[m.source] = Empty
	b.squares[m.destination] = m.piece
	if DEBUG {

		fmt.Printf("%s\n", b)
	}

}
func (b *Board) undoMove(m *Move) {
	if DEBUG {
		println("undo move", m.String())
	}
	b.squares[m.source] = m.piece
	b.squares[m.destination] = m.capture
	if DEBUG {
		fmt.Printf("%s\n", b)
	}
}
