package emil

import (
	"fmt"
	"strconv"
	"strings"
)

type DTM struct {
	dtm  int // Depth to mate
	move *Move
}

func (d *DTM) String() string {
	return fmt.Sprintf("%d/%s", d.dtm, d.move)
}

func DTMsFromString(s string) (list []*DTM) {
	if s == "[]" {
		return list
	}
	s = strings.Replace(s, "[", "", -1)
	s = strings.Replace(s, "]", "", -1)

	for _, item := range strings.Split(s, " ") {
		parts := strings.Split(item, "/")
		dtm, err := strconv.Atoi(parts[0])
		if err != nil {
			panic("can not parse " + parts[0] + " to integer")
		}
		move := MoveFromString(parts[1])
		d := &DTM{dtm: dtm, move: move}
		list = append(list, d)
	}
	return list
}

type Analysis struct {
	board    *Board `json:"-"`
	dtmWhite []*DTM `json:"dtmWhite"`
	dtmBlack []*DTM `json:"dtmBlack"`

	moves map[string]bool
}

func (a *Analysis) String() string {
	return fmt.Sprintf("WHITE: %s\nBLACK: %s\nFEN: %s\n%s\n",
		a.dtmWhite,
		a.dtmBlack,
		NewPosition(a.board, WHITE),
		a.board)
}

func (a *Analysis) DTMs(player int) []*DTM {
	if player == WHITE {
		return a.dtmWhite
	}
	return a.dtmBlack
}

func (a *Analysis) addDTM(move *Move, dtm int) {
	if _, ok := a.moves[move.String()]; ok {
		return // we have this move allready
	}
	a.moves[move.String()] = true

	if move.player == WHITE {
		a.dtmWhite = append(a.dtmWhite, &DTM{move: move, dtm: dtm})
	} else {
		a.dtmBlack = append(a.dtmBlack, &DTM{move: move, dtm: dtm})
	}
}

func (a *Analysis) playerHaveDTMs() bool {
	return (len(a.dtmWhite) + len(a.dtmBlack)) > 0
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
