package emil

import (
	"fmt"
)

// Move represents a move on the board
type Move struct {
	player      int
	piece       int
	capture     int
	promotion   int
	source      int
	destination int
	isCapture   bool
	isQueenside bool
	isKingside  bool
	Str         string `json:"move"`
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
	m2 := &Move{
		player:      m.player,
		piece:       m.piece,
		capture:     m.capture,
		promotion:   m.promotion,
		source:      m.destination,
		destination: m.source,
		isCapture:   m.isCapture,
		isQueenside: m.isQueenside,
		isKingside:  m.isKingside}

	m2.Str = m2.String()
	return m2

}
func MoveFromString(str string) *Move {
	m := &Move{
		player:      player,
		piece:       piece,
		capture:     Empty,
		promotion:   Empty,
		source:      src,
		destination: dst,
		isCapture:   false,
		isQueenside: false,
		isKingside:  false}
	m.Str = m.str
	return m
}

func newSilentMove(player, piece, src, dst int) *Move {
	m := &Move{
		player:      player,
		piece:       piece,
		capture:     Empty,
		promotion:   Empty,
		source:      src,
		destination: dst,
		isCapture:   false,
		isQueenside: false,
		isKingside:  false}
	m.Str = m.String()
	return m
}
func newCaptureMove(player, piece, capture, src, dst int) *Move {
	m := &Move{
		player:      player,
		piece:       piece,
		capture:     capture,
		promotion:   Empty,
		source:      src,
		destination: dst,
		isCapture:   true,
		isQueenside: false,
		isKingside:  false}
	m.Str = m.String()
	return m
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
