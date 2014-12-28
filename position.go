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
