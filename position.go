package emil

type position struct {
	board  *Board
	player int
}

func NewPosition(board *Board, player int) *position {
	return &position{
		board:  board,
		player: player}
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
