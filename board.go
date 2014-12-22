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

func isOwnPiece(player, capture int) bool {
	return (player == WHITE && capture > 0) ||
		(player == BLACK && capture < 0)
}

func otherPlayer(player int) int {
	if player == WHITE {
		return BLACK
	}
	return WHITE
}

//Moves prints all moves for a piece on square
func (b *Board) Moves(player int, testKingCapture bool) string {
	var result, list []*Move
	for src, piece := range b.squares {
		if isOwnPiece(player, piece) {
			switch abs(piece) {
			case kingValue:
				for _, dst := range kingDestinationsFrom(src) {
					capture := b.squares[dst]
					if isKing(capture) {
						return "KingCaptured"
					}
					if capture == Empty {
						list = append(list, newSilentMove(player, piece, src, dst))
					} else if !isOwnPiece(player, capture) {
						list = append(list, newCaptureMove(player, piece, capture, src, dst))
					}
				}
			case rockValue:
				for _, dsts := range rockDestinationsFrom(src) {
					for _, dst := range dsts {
						capture := b.squares[dst]
						if isKing(capture) {
							return "KingCaptured"
						}
						if capture == Empty {
							list = append(list, newSilentMove(player, piece, src, dst))
						} else if !isOwnPiece(player, capture) {
							list = append(list, newCaptureMove(player, piece, capture, src, dst))
							break
						} else {
							break // onOwnPiece
						}
					}
				}
			}
		}
	}

	if testKingCapture {
		return "no KingCaptured"
	}

	for _, m := range list {
		println("TEST move", m.String())
		b.doMove(m)
		if !b.isKingCapturedAfter(m) {
			result = append(result, m)
		} else {
			println("KingCaptured")
		}
		b.undoMove(m)
		fmt.Printf("\n\n\n")
	}

	return moveList(result)
}
func isKing(piece int) bool {
	return abs(piece) == kingValue
}

func (b *Board) isKingCapturedAfter(m *Move) bool {
	return "KingCaptured" == b.Moves(otherPlayer(m.player), true)
}

func (b *Board) doMove(m *Move) {
	println("do move", m.String())
	b.squares[m.source] = Empty
	b.squares[m.destination] = m.piece
	fmt.Printf("%s\n", b)

}
func (b *Board) undoMove(m *Move) {
	println("undo move", m.String())
	b.squares[m.source] = m.piece
	b.squares[m.destination] = m.capture
	fmt.Printf("%s\n", b)
}
