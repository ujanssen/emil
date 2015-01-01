package emil

import (
	"fmt"
)

type DTM struct {
	dtm  int   `json:"dtm"` // Depth to mate
	move *Move `json:"move"`
}

func (d *DTM) String() string {
	return fmt.Sprintf("%d/%s", d.dtm, d.move)
}

type Analysis struct {
	board     *Board `json:"board"`
	dtmWhite  []*DTM `json:"dtmWhite"`
	dtmWBlack []*DTM `json:"dtmWBlack"`

	moves map[string]bool `json:"moves"`
}

func (a *Analysis) String() string {
	return fmt.Sprintf("WHITE: %s\nBLACK: %s\nFEN: %s\n%s\n",
		a.dtmWhite,
		a.dtmWBlack,
		NewPosition(a.board, WHITE),
		a.board)
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
