package emil

import (
	"fmt"
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

func (p *position) String() string {
	s := ""
	for _, r := range FirstSquares {
		if len(s) > 0 {
			s += "/"
		}
		for f := 0; f < 8; f++ {
			s += fmt.Sprintf("%s", symbol(p.board.squares[r+f]))
		}
	}
	s = strings.Replace(s, "        ", "8", -1)
	s = strings.Replace(s, "       ", "7", -1)
	s = strings.Replace(s, "      ", "6", -1)
	s = strings.Replace(s, "     ", "5", -1)
	s = strings.Replace(s, "    ", "4", -1)
	s = strings.Replace(s, "   ", "3", -1)
	s = strings.Replace(s, "  ", "2", -1)
	s = strings.Replace(s, " ", "1", -1)
	return s
}
