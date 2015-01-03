package emil

import (
	"errors"
	"fmt"
	"time"
)

var errNowNewAnalysis = errors.New("errNowNewAnalysis")

func (db *EndGameDb) addAnalysis(board *Board) {
	a := NewAnalysis(board)
	db.AnalysisMap[a.board.String()] = a
}

func (db *EndGameDb) addDTMToAnalysis(board *Board, dtm int, move *Move) {
	a := db.AnalysisMap[board.String()]
	if move != nil {
		a.addDTM(move.reverse(), dtm)
	}
	if dtm >= 0 {
		db.dtmDb[dtm][board.String()] = true
		if move != nil {
			playerForStep := playerForStepN(dtm)
			if playerForStep != move.player {
				panic("playerForStep != move.player")
			}
		}
	}
}

// find positions where black is checkmate
func (db *EndGameDb) retrogradeAnalysisStep1() {
	db.dtmDb = append(db.dtmDb, make(map[string]bool))

	start := time.Now()

	player := BLACK
	for boardStr, a := range db.AnalysisMap {
		// mate only on border square
		blackKingSquare := BoardSquares[a.board.blackKing]
		if !blackKingSquare.isBorder {
			continue
		}
		// mate only with help from king
		if squaresDistances[a.board.blackKing][a.board.whiteKing] > 2 {
			continue
		}

		p := NewPosition(a.board, player)

		move := Search(p)
		if move == nil {
			if isKingInCheck(p) {
				db.addDTMToAnalysis(a.board, 0, nil)
				if DEBUG {
					fmt.Printf("mate:\n%s\n", boardStr)
				}
			}
		}
	}
	end := time.Now()
	if DEBUG {
		fmt.Printf("db.dtmDb[0] %d\n", len(db.dtmDb[0]))
		fmt.Printf("duration %v\n\n\n", end.Sub(start))
	}
}
func playerForStepN(dtm int) (player int) {
	if dtm%2 == 0 {
		return BLACK
	}
	return WHITE
}

func (db *EndGameDb) retrogradeAnalysisStepN(dtm int) (noError error) {
	start := time.Now()
	db.dtmDb = append(db.dtmDb, make(map[string]bool))

	player := playerForStepN(dtm)

	positions := 0
	if player == WHITE {
		if DEBUG {
			fmt.Printf("WHITE Start positions %d\n", len(db.dtmDb[dtm-1]))
		}
		for str := range db.dtmDb[dtm-1] {
			a := db.AnalysisMap[str]
			p := NewPosition(a.board, player)
			list := generateMoves(p)
			moves := filterKingCaptures(p, list)
			moves = filterKingCaptures(NewPosition(a.board, otherPlayer(player)), list)

			for _, m := range moves {
				newBoard := a.board.doMove(m)
				db.addDTMToAnalysis(newBoard, dtm, m)
			}
		}
	} else {
		for _, a := range db.AnalysisMap {
			if db.isMateIn0246(a.board, dtm) >= 0 {
				positions++
			}
		}
		if DEBUG {
			fmt.Printf("BLACK Start positions %d\n", len(db.AnalysisMap)-positions)
		}
		for _, a := range db.AnalysisMap {
			if db.isMateIn0246(a.board, dtm) >= 0 {
				continue
			}
			p := NewPosition(a.board, player)
			moves := GenerateMoves(p)

			found := 0
			maxDTM := -1
			for _, m := range moves {
				newBoard := a.board.doMove(m)
				newDtm := db.isMateIn1357(newBoard, dtm)
				if newDtm > maxDTM {
					maxDTM = newDtm
				}
				if db.isMateIn1357(newBoard, dtm) >= 0 {
					found++
				}
			}

			if found == len(moves) {
				for _, m := range moves {
					db.addDTMToAnalysis(a.board, maxDTM+1, m)
				}
			}
		}
	}
	end := time.Now()

	if DEBUG {
		fmt.Printf("db.dtmDb[%d] %d\n", dtm, len(db.dtmDb[dtm]))
		fmt.Printf("duration %v\n\n\n", end.Sub(start))
	}

	if len(db.dtmDb[dtm]) == 0 {
		return errNowNewAnalysis
	}
	return noError
}
func (db *EndGameDb) isMateIn0246(board *Board, maxDtm int) int {
	for dtm := 0; dtm < maxDtm; dtm += 2 {
		_, ok := db.dtmDb[dtm][board.String()]
		if ok {
			return dtm
		}
	}
	return -1
}
func (db *EndGameDb) isMateIn1357(board *Board, maxDtm int) int {
	for dtm := 1; dtm < maxDtm; dtm += 2 {
		_, ok := db.dtmDb[dtm][board.String()]
		if ok {
			return dtm
		}
	}
	return -1
}

func (db *EndGameDb) retrogradeAnalysis() {
	// find positions where black is checkmate
	db.retrogradeAnalysisStep1()
	dtm := 1
	for {
		err := db.retrogradeAnalysisStepN(dtm)
		if err != nil {
			break
		}
		if IN_TEST {
			return
		}
		dtm++
	}
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
