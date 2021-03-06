package emil

import (
	"errors"
	"fmt"
	"strings"
)

var errPieceNotFound = errors.New("piece not found")
var errNotEmpty = errors.New("not empty")
var errKingsToClose = errors.New("Kings to close")

// Board with an array of field values, representing pieces
type Board struct {
	Squares []int

	str string
}

// NewBoard creates a new board
func NewBoard() *Board {
	return &Board{Squares: make([]int, SQUARES)}
}

// Fen2Board creates a new board from a fen string
func Fen2Board(fen string) *Board {
	b := NewBoard()
	b.str = fen

	for fs, part := range strings.Split(fen, "/") {
		sq := FirstSquares[fs]
		i := 0
		for _, r := range part {
			s := string(r)
			switch s {
			case "1", "8":
				//nothing to do
			case "7":
				i += 6
			case "6":
				i += 5
			case "5":
				i += 4
			case "4":
				i += 3
			case "3":
				i += 2
			case "2":
				i += 1
			default:
				piece, ok := Symbols[s]
				if !ok {
					panic("can't parse " + s + " in " + part)
				} else {
					b.Setup(piece, sq+i)
				}
			}
			i++
		}
	}
	return b
}
func (b *Board) find(piece int) (int, error) {
	for _, s := range b.Squares {
		if s == piece {
			return s, nil
		}
	}
	return -1, errPieceNotFound
}

func (b *Board) WhiteKing() int {
	s, err := b.find(WhiteKing)
	if err != nil {
		panic("White king not found")
	}
	return s
}
func (b *Board) BlackKing() int {
	s, err := b.find(BlackKing)
	if err != nil {
		panic("Black king not found")
	}
	return s
}
func (b *Board) Square(i int) int {
	return b.Squares[i]
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
			s += fmt.Sprintf("%s", Pieces[b.Squares[r+f]])
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
			s += fmt.Sprintf("%s ", Pieces[b.Squares[r+f]])
		}
		s += fmt.Sprintf("%d\n", BoardSquares[r].rank)
	}
	s += "  " + files + " \n"
	s += "FEN: " + b.String()
	return s
}

//Empty removes a piece from a square
func (b *Board) Empty(square int) {
	b.Squares[square] = Empty
	b.str = ""
}

//Setup a piece on a square
func (b *Board) Setup(piece, square int) (noError error) {
	if b.Squares[square] != Empty {
		return errNotEmpty
	}

	b.Squares[square] = piece

	return noError
}

func (b *Board) DoMove(m *Move) (newBoard *Board) {
	newBoard = NewBoard()
	copy(newBoard.Squares, b.Squares)
	newBoard.Squares[m.Source] = Empty
	newBoard.Squares[m.Destination] = m.Piece
	return newBoard
}
