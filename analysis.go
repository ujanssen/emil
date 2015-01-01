package emil

import (
	"fmt"
)

type DTM struct {
	Dtm  int   `json:"dtm"` // Depth to mate
	move *Move `json:"move"`
}

func (d *DTM) String() string {
	return fmt.Sprintf("%d/%s", d.Dtm, d.move)
}

type Analysis struct {
	Board    *Board `json:"board"`
	DtmWhite []*DTM `json:"dtmWhite"`
	DtmBlack []*DTM `json:"dtmBlack"`

	moves map[string]bool
}

func (a *Analysis) String() string {
	return fmt.Sprintf("WHITE: %s\nBLACK: %s\nFEN: %s\n%s\n",
		a.DtmWhite,
		a.DtmBlack,
		NewPosition(a.Board, WHITE),
		a.Board)
}

func (a *Analysis) DTMs(player int) []*DTM {
	if player == WHITE {
		return a.DtmWhite
	}
	return a.DtmBlack
}
func (a *Analysis) addDTM(move *Move, dtm int) {
	if _, ok := a.moves[move.String()]; ok {
		return // we have this move allready
	}
	a.moves[move.String()] = true

	if move.player == WHITE {
		a.DtmWhite = append(a.DtmWhite, &DTM{move: move, Dtm: dtm})
	} else {
		a.DtmBlack = append(a.DtmBlack, &DTM{move: move, Dtm: dtm})
	}
}

func (a *Analysis) playerHaveDTMs() bool {
	return (len(a.DtmWhite) + len(a.DtmBlack)) > 0
}

func (a *Analysis) BestMove(player int) (bestMove *Move) {
	dtms := a.DTMs(player)
	minDTM := 9999

	// white minimizes DTM
	for _, d := range dtms {
		if d.Dtm < minDTM {
			minDTM = d.Dtm
			bestMove = d.move
		}
	}
	// TODO best move for black
	return bestMove
}
