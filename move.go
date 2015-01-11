package emil

import (
	"fmt"
)

// Move represents a move on the board
type Move struct {
	Player      int
	Piece       int
	Capture     int
	Source      int
	Destination int
	IsCapture   bool
}

func (m *Move) String() string {
	if !m.IsCapture {
		return fmt.Sprintf("%s%s%s",
			Pieces[m.Piece],
			BoardSquares[m.Source].name,
			BoardSquares[m.Destination].name)
	}
	return fmt.Sprintf("%s%sx%s",
		Pieces[m.Piece],
		BoardSquares[m.Source].name,
		BoardSquares[m.Destination].name)
}

func (m *Move) reverse() *Move {
	return &Move{
		Player:      m.Player,
		Piece:       m.Piece,
		Capture:     m.Capture,
		Source:      m.Destination,
		Destination: m.Source,
		IsCapture:   m.IsCapture}

}

// MoveFromString parses a string like Ke1e2
// to a *Move
func MoveFromString(str string) *Move {
	player := WHITE
	piece := Symbols[string(str[0])]
	if piece < 0 {
		player = BLACK
	}

	return &Move{
		Player:      player,
		Piece:       piece,
		Capture:     Empty,
		Source:      squareMap[string(str[1:3])],
		Destination: squareMap[string(str[3:5])],
		IsCapture:   false}
}

func newSilentMove(player, piece, src, dst int) *Move {
	return &Move{
		Player:      player,
		Piece:       piece,
		Capture:     Empty,
		Source:      src,
		Destination: dst,
		IsCapture:   false}
}
func newCaptureMove(player, piece, capture, src, dst int) *Move {
	m := newSilentMove(player, piece, src, dst)
	m.IsCapture = true
	return m
}
