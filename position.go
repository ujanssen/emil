package emil

import (
	"strings"
)

type position struct {
	Board  *Board
	Player int
}

func NewPosition(board *Board, player int) *position {
	return &position{
		Board:  board,
		Player: player}
}

func PositionFromKey(key string) *position {
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

func (p *position) key() positionKey {
	return positionKey(p.String())
}
func (p *position) String() string {
	s := p.Board.String()
	if p.Player == WHITE {
		s += " w"
	} else {
		s += " b"
	}
	return s
}
