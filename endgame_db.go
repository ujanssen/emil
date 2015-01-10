package emil

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

var errNowNewAnalysis = errors.New("errNowNewAnalysis")

func (db *EndGameDb) addAnalysis(board *Board) {
	boardStr := board.String()
	a := NewAnalysis(board)
	db.AnalysisMap[boardStr] = a
	db.retrogradeAnalysisStep0(a)
	db.AnalysisStr[boardStr] = fmt.Sprintf("%v.%v", a.dtmWhite, a.dtmBlack)
}

func (db *EndGameDb) addAnalysisFromStr(board *Board, str string) {
	boardStr := board.String()
	a := NewAnalysis(board)
	db.AnalysisMap[boardStr] = a

	parts := strings.Split(str, ".")

	a.dtmWhite = DTMsFromString(parts[0])
	a.dtmBlack = DTMsFromString(parts[1])
}

func (db *EndGameDb) addMate(board *Board) {
	a := db.AnalysisMap[board.String()]
	a.dtm = 0
}

/*
func (db *EndGameDb) addDTMToAnalysis(board *Board, dtm int, move *Move) bool {
	if dtm == 0 {
		panic("dtm == 0")
	}
	if move == nil {
		panic("move == nil")
	}
	a := db.AnalysisMap[board.String()]
	added := a.addDTM(move.reverse(), dtm)
	if !added {
		return false
	}
	if move != nil {
		playerForStep := playerForStepN(dtm)
		if playerForStep != move.player {
			panic("playerForStep != move.player")
		}
	}
	return true
}
*/
func (db *EndGameDb) addDTMToAnalysis(board *Board, dtm int, move *Move) {
	if dtm == 0 {
		panic("dtm == 0")
	}
	if move == nil {
		panic("move == nil")
	}

	playerForStep := playerForStepN(dtm)
	if playerForStep != move.player {
		panic("playerForStep != move.player")
	}

	a := db.AnalysisMap[board.String()]
	a.addDTM(move, dtm)
}

// generate all moves
func (db *EndGameDb) retrogradeAnalysisStep0(a *Analysis) {
	player := WHITE
	p := NewPosition(a.board, player)
	moves := GenerateMoves(p)
	for _, m := range moves {
		newBoard := a.board.DoMove(m)
		a.addMoveToAnalysis(m, newBoard)
	}

	player = BLACK
	p = NewPosition(a.board, player)
	moves = GenerateMoves(p)
	for _, m := range moves {
		newBoard := a.board.DoMove(m)
		a.addMoveToAnalysis(m, newBoard)
	}
}
func playerForStepN(dtm int) (player int) {
	if dtm%2 == 0 {
		return BLACK
	}
	return WHITE
}

// find positions where black is checkmate
func (db *EndGameDb) retrogradeAnalysisStep1() {
	start := time.Now()

	player := BLACK
	for boardStr, a := range db.AnalysisMap {
		// mate only on border square
		blackKingSquare := BoardSquares[a.board.BlackKing()]
		if !blackKingSquare.isBorder {
			continue
		}
		// mate only with help from king
		if squaresDistances[a.board.BlackKing()][a.board.WhiteKing()] > 2 {
			continue
		}

		p := NewPosition(a.board, player)
		if len(a.dtmBlack) == 0 && isKingInCheck(p) {
			db.addMate(a.board)
			if DEBUG {
				fmt.Printf("mate: %s\n", boardStr)
			}
		}
	}
	end := time.Now()
	if DEBUG {
		fmt.Printf("db.dtmDb[0] %d\n", len(db.FindMates()))
		fmt.Printf("duration %v\n\n\n", end.Sub(start))
	}
}

func (db *EndGameDb) retrogradeAnalysisStepNForBlack(dtm int) (noError error) {
	start := time.Now()
	newMovesFound := 0

	if DEBUG {
		positions := 0
		for _, a := range db.AnalysisMap {
			if db.isMateIn0246(a.board, dtm) >= 0 {
				positions++
			}
		}
		fmt.Printf("BLACK Start %d positions %d / ignored positions %d\n",
			dtm, len(db.AnalysisMap)-positions, positions)
	}
	for _, a := range db.AnalysisMap {
		if db.isMateIn0246(a.board, dtm) >= 0 {
			continue
		}
		found := 0
		minDTM := 500
		for _, d := range a.dtmBlack {
			newDtm := db.isMateIn1357(d.board, dtm)
			if newDtm < minDTM {
				minDTM = newDtm
			}
			if newDtm >= 0 {
				found++
			}
		}

		if found == len(a.dtmBlack) {
			for _, d := range a.dtmBlack {
				db.addDTMToAnalysis(a.board, minDTM+1, d.move)
				newMovesFound++
			}
		}
	}
	end := time.Now()

	if DEBUG {
		for i := 0; i <= dtm; i++ {
			moves := 0
			dtms := 0
			for _, a := range db.AnalysisMap {
				if a.dtm == i {
					dtms++
					moves += a.dtmMoves()
				}
			}
			if dtms > 0 {
				fmt.Printf("db.dtmDb[%2d] %6d/%6d\n", i, dtms, moves)
			}
		}
		fmt.Printf("\nnewMovesFound %d\n", newMovesFound)
		fmt.Printf("duration %v\n\n", end.Sub(start))
	}

	if newMovesFound == 0 {
		return errNowNewAnalysis
	}
	return noError
}

func (db *EndGameDb) retrogradeAnalysisStepNforWhite(dtm int) (noError error) {
	start := time.Now()
	newMovesFound := 0
	positions := 0
	var list []*Analysis
	for _, a := range db.AnalysisMap {
		if db.isMateIn0246(a.board, 0) >= 0 {
			positions++
			list = append(list, a)
		}
	}
	fmt.Printf("WHITE Start %d positions %d / ignored positions %d\n",
		dtm, positions, len(db.AnalysisMap)-positions)

	// get mate positions
	for _, a := range list {
		for _, newA := range db.AnalysisMap {
			for _, d := range newA.dtmWhite {
				if a.board.String() == d.board.String() {
					if newA.dtm == initial {
						newA.dtm = dtm
					}
					if d.dtm == initial {
						d.dtm = dtm
						newMovesFound++
					}
				}
			}
		}
	}
	end := time.Now()

	if DEBUG {
		for i := 0; i <= dtm; i++ {
			dtms := 0
			for _, a := range db.AnalysisMap {
				if a.dtm == i {
					dtms++
				}
			}
			fmt.Printf("db.dtmDb[%2d] %6d\n", i, dtms)
		}
		fmt.Printf("\nnewMovesFound %d\n", newMovesFound)
		fmt.Printf("duration %v\n\n", end.Sub(start))
	}

	// search mate positions in DTM.board
	//mark als dtm 1
	fmt.Printf("duration %v\n\n", end.Sub(start))

	return noError

}
func (db *EndGameDb) isMateIn0246(board *Board, maxDtm int) int {
	if a, ok := db.AnalysisMap[board.String()]; ok {
		for dtm := 0; dtm <= maxDtm; dtm += 2 {
			if a.dtm == dtm {
				return dtm
			}
		}
	}
	return -1
}
func (db *EndGameDb) isMateIn1357(board *Board, maxDtm int) int {
	if a, ok := db.AnalysisMap[board.String()]; ok {
		for dtm := 1; dtm <= maxDtm; dtm += 2 {
			if a.dtm == dtm {
				return dtm
			}
		}
	}
	return -1
}

func (db *EndGameDb) retrogradeAnalysis() {
	// find positions where black is checkmate
	db.retrogradeAnalysisStep1()
	db.retrogradeAnalysisStepNforWhite(1)

	/*
		for {
			err := db.retrogradeAnalysisStepNforWhite(dtm)
			dtm++
			if err != nil {
				break
			}
			if IN_TEST {
				return
			}
			err = db.retrogradeAnalysisStepNForBlack(dtm)
			dtm++
			if err != nil {
				break
			}

			if dtm >= 114 {
				break
			}
		}
	*/
}
func generateMoves(p *position) (list []*Move) {
	for src, piece := range p.board.squares {
		if isOwnPiece(p.player, piece) {
			switch abs(piece) {
			case kingValue:
				for _, dst := range kingDestinationsFrom(src) {
					capture := p.board.squares[dst]
					if isOtherKing(p.player, capture) {
						continue
					}
					if capture == Empty {
						list = append(list, newSilentMove(p.player, piece, src, dst))
					} else if !isOwnPiece(p.player, capture) {
						list = append(list, newCaptureMove(p.player, piece, capture, src, dst))
					}
				}
			case rockValue:
				for _, dsts := range rockDestinationsFrom(src) {
					for _, dst := range dsts {
						capture := p.board.squares[dst]
						if isOtherKing(p.player, capture) {
							break
						}
						if capture == Empty {
							list = append(list, newSilentMove(p.player, piece, src, dst))
						} else if !isOwnPiece(p.player, capture) {
							list = append(list, newCaptureMove(p.player, piece, capture, src, dst))
							break
						} else {
							break // onOwnPiece
						}
					}
				}
			}
		}
	}
	return list
}
