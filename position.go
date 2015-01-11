package emil

import (
	"strings"
)

type Position struct {
	Board  *Board
	Player int
}

func NewPosition(board *Board, player int) *Position {
	return &Position{
		Board:  board,
		Player: player}
}

func PositionFromKey(key string) *Position {
	parts := strings.Split(key, " ")
	var player int

	if parts[1] == "w" {
		player = WHITE
	}
	if parts[1] == "b" {
		player = BLACK
	}
	board := Fen2Board(parts[0])

	return NewPosition(board, player)
}

func (p *Position) key() positionKey {
	return positionKey(p.String())
}
func (p *Position) String() string {
	s := p.Board.String()
	if p.Player == WHITE {
		s += " w"
	} else {
		s += " b"
	}
	return s
}
