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
}

func (a *Analysis) DTMs(player int) []*DTM {
	if player == WHITE {
		return a.dtmWhite
	}
	return a.dtmWBlack
}
func (a *Analysis) addDTM(move *Move, dtm int) {
	dtms := a.DTMs(move.player)
	dtms = append(dtms, &DTM{move: move, dtm: dtm})
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
	return bestMove
}

func (a *Analysis) Board() *Board {
	return a.board
}
