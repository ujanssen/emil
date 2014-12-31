package emil

import (
	"fmt"
)

type DTM struct {
	dtm  int // Depth to mate
	move *Move
}

func (d *DTM) String() string {
	return fmt.Sprintf("%d/%s", d.dtm, d.move)
}

type Analysis struct {
	board     *Board
	dtmWhite  []*DTM
	dtmWBlack []*DTM

	moves map[string]bool
}

func (a *Analysis) String() string {
	return fmt.Sprintf("WHITE:%s\nBLACK:%s\n%s\nFEN: %s\n\n",
		a.dtmWhite,
		a.dtmWBlack,
		a.board, NewPosition(a.board, WHITE))
}

func (a *Analysis) DTMs(player int) []*DTM {
	if player == WHITE {
		return a.dtmWhite
	}
	return a.dtmWBlack
}
func (a *Analysis) addDTM(move *Move, dtm int) {
	if _, ok := a.moves[move.String()]; ok {
		return // we have this move allready
	}
	a.moves[move.String()] = true

	if move.player == WHITE {
		a.dtmWhite = append(a.dtmWhite, &DTM{move: move, dtm: dtm})
	} else {
		a.dtmWBlack = append(a.dtmWBlack, &DTM{move: move, dtm: dtm})
	}
}

func (a *Analysis) playerHaveDTMs() bool {
	return (len(a.dtmWhite) + len(a.dtmWBlack)) > 0
}

func (a *Analysis) BestMove(player int) (bestMove *Move) {
	dtms := a.DTMs(player)
	minDTM := 9999

	// white minimizes DTM
	for _, d := range dtms {
		if d.dtm < minDTM {
			minDTM = d.dtm
			bestMove = d.move
		}
	}
	// TODO best move for black
	return bestMove
}

func (a *Analysis) Board() *Board {
	return a.board
}
