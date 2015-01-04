package emil

import (
	"strings"
)

type position struct {
	board  *Board
	player int
}

func NewPosition(board *Board, player int) *position {
	return &position{
		board:  board,
		player: player}
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
	s := p.board.String()
	if p.player == WHITE {
		s += " w"
	} else {
		s += " b"
	}
	return s
}
