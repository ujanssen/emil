package emil

import (
	"fmt"
)

// Move represents a move on the board
type Move struct {
	player, piece           int
	capture, promotion      int
	source, destination     int
	isCapture               bool
	isQueenside, isKingside bool
}

func (m *Move) String() string {
	if !m.isCapture {
		return fmt.Sprintf("%s%s%s",
			Pieces[m.piece],
			BoardSquares[m.source].name,
			BoardSquares[m.destination].name)
	}
	return fmt.Sprintf("%s%sx%s",
		Pieces[m.piece],
		BoardSquares[m.source].name,
		BoardSquares[m.destination].name)
}

func (m *Move) reverse() *Move {
	return &Move{
		player:      m.player,
		piece:       m.piece,
		capture:     m.capture,
		promotion:   m.promotion,
		source:      m.destination,
		destination: m.source,
		isCapture:   m.isCapture,
		isQueenside: m.isQueenside,
		isKingside:  m.isKingside}
}

func newSilentMove(player, piece, src, dst int) *Move {
	return &Move{
		player:      player,
		piece:       piece,
		capture:     Empty,
		promotion:   Empty,
		source:      src,
		destination: dst,
		isCapture:   false,
		isQueenside: false,
		isKingside:  false}
}
func newCaptureMove(player, piece, capture, src, dst int) *Move {
	return &Move{
		player:      player,
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
