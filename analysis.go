package emil

import (
	"fmt"
	"strconv"
	"strings"
)

const initial = -1

type DTM struct {
	dtm   int // Depth to mate
	move  *Move
	board *Board
}

func (d *DTM) String() string {
	return fmt.Sprintf("%d,%s,%s", d.dtm, d.move, d.board)
}

func DTMsFromString(s string) (list []*DTM) {
	if s == "[]" {
		return list
	}
	s = strings.Replace(s, "[", "", -1)
	s = strings.Replace(s, "]", "", -1)

	for _, item := range strings.Split(s, " ") {
		parts := strings.Split(item, ",")
		dtm, err := strconv.Atoi(parts[0])
		if err != nil {
			panic("can not parse " + parts[0] + " to integer")
		}
		move := MoveFromString(parts[1])
		board := Fen2Board(parts[2])
		d := &DTM{dtm: dtm, move: move, board: board}
		list = append(list, d)
	}
	return list
}

type Analysis struct {
	board    *Board
	dtm      int
	dtmWhite []*DTM
	dtmBlack []*DTM

	moves map[string]bool
}

func (a *Analysis) String() string {
	return fmt.Sprintf("WHITE: %s\nBLACK: %s\nFEN: %s\n%s\n",
		a.dtmWhite,
		a.dtmBlack,
		NewPosition(a.board, WHITE),
		a.board)
}
func NewAnalysis(board *Board) *Analysis {
	return &Analysis{
		dtmWhite: make([]*DTM, 0),
		dtmBlack: make([]*DTM, 0),
		board:    board,
		dtm:      initial,
		moves:    make(map[string]bool)}
}
func (a *Analysis) DTMs(player int) []*DTM {
	if player == WHITE {
		return a.dtmWhite
	}
	return a.dtmBlack
}
func (a *Analysis) addMoveToAnalysis(move *Move, board *Board) {
	if move.player == WHITE {
		a.dtmWhite = append(a.dtmWhite, &DTM{move: move, board: board})
	} else {
		a.dtmBlack = append(a.dtmBlack, &DTM{move: move, board: board})
	}
}

func (a *Analysis) addDTM(move *Move, dtm int) bool {
	if _, ok := a.moves[move.String()]; ok {
		return false // we have this move allready
	}
	if dtm < a.dtm {
		a.dtm = dtm
	}
	a.moves[move.String()] = true

	if move.player == WHITE {
		a.dtmWhite = append(a.dtmWhite, &DTM{move: move, dtm: dtm})
	} else {
		a.dtmBlack = append(a.dtmBlack, &DTM{move: move, dtm: dtm})
	}
	return true
}

func (a *Analysis) playerHaveDTMs() bool {
	return (len(a.dtmWhite) + len(a.dtmBlack)) > 0
}
func (a *Analysis) dtmMoves() int {
	return len(a.dtmWhite) + len(a.dtmBlack)
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
