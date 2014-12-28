package emil

type position struct {
	board  *Board
	player int
}

func newPosition(board *Board, player int) *position {
	return &position{
		board:  board,
		player: player}
}
