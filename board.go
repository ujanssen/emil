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

func symbol(piece int) string {
	switch piece {
	case WhiteKing:
		return "K"
	case BlackKing:
		return "k"
	case WhiteRock:
		return "R"
	case BlackRock:
		return "r"
	default:
		return " "
	}
}

//Setup a piece on a square
func (b *Board) Setup(piece, square int) {
	b.squares[square] = piece
}

//Moves prints all moves for a piece on square
func (b *Board) Moves(player int) string {
	var list []*Move
	for src, piece := range b.squares {
		if (player == WHITE && piece > 0) ||
			(player == BLACK && piece < 0) {
			switch abs(piece) {
			case kingValue:
				for _, dst := range kingDestinationsFrom(src) {
					dstPiece := b.squares[dst]
					if dstPiece == Empty {
						list = append(list, newSilentMove(piece, src, dst))
					}
				}
			case rockValue:
				for _, dsts := range rockDestinationsFrom(src) {
					for _, dst := range dsts {
						dstPiece := b.squares[dst]
						if dstPiece == Empty {
							list = append(list, newSilentMove(piece, src, dst))
						}
					}
				}
			}
		}
	}

	return moveList(list)
}
