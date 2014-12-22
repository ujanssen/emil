package emil

import (
	"fmt"
)

type Move struct {
	piece, capture, promotion int
	source, destination       int
	isCapture                 bool
	isQueenside, isKingside   bool
}

func (m *Move) String() string {
	return fmt.Sprintf("%s%s%s",
		Pieces[m.piece],
		BoardSquares[m.source].name,
		BoardSquares[m.destination].name)
}

func newSilentMove(piece, src, dst int) *Move {
	return &Move{
		piece:       piece,
		capture:     Empty,
		promotion:   Empty,
		source:      src,
		destination: dst,
		isCapture:   false,
		isQueenside: false,
		isKingside:  false}
}
func newCaptureMove(piece, capture, src, dst int) *Move {
	return &Move{
		piece:       piece,
		capture:     capture,
		promotion:   Empty,
		source:      src,
		destination: dst,
		isCapture:   true,
		isQueenside: false,
		isKingside:  false}
}

func moveList(list []*Move) string {
	r := "["
	for i, m := range list {
		if i > 0 {
			r += ", "
		}
		r += m.String()
	}
	r += "]"
	return r
}
