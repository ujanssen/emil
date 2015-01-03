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

func (db *EndGameDb) addMate(board *Board) {
	a := db.AnalysisMap[board.String()]
	a.dtm = 0
}

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

// find positions where black is checkmate
func (db *EndGameDb) retrogradeAnalysisStep1() {
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
				db.addMate(a.board)
				if DEBUG {
					fmt.Printf("mate:\n%s\n", boardStr)
				}
			}
		}
	}
	end := time.Now()
	if DEBUG {
		fmt.Printf("db.dtmDb[0] %d\n", len(db.FindMates()))
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
	newMovesFound := 0

	player := playerForStepN(dtm)

	if DEBUG {
		positions := 0
		for _, a := range db.AnalysisMap {
			if db.isMateIn0246(a.board, dtm) >= 0 {
				positions++
			}
		}
		if player == WHITE {
			fmt.Printf("WHITE Start %d positions %d / ignored positions %d\n",
				dtm, positions, len(db.AnalysisMap)-positions)
		} else {
			fmt.Printf("BLACK Start %d positions %d / ignored positions %d\n",
				dtm, len(db.AnalysisMap)-positions, positions)
		}
	}
	if player == WHITE {
		for _, a := range db.AnalysisMap {
			positionDtm := db.isMateIn0246(a.board, dtm)
			if positionDtm == -1 {
				continue
			}

			p := NewPosition(a.board, player)
			list := generateMoves(p)
			moves := filterKingCaptures(p, list)
			moves = filterKingCaptures(NewPosition(a.board, otherPlayer(player)), list)

			for _, m := range moves {
				newBoard := a.board.DoMove(m)
				f := db.addDTMToAnalysis(newBoard, positionDtm+1, m)
				if f {
					newMovesFound++
				}
			}
		}
	} else {
		for _, a := range db.AnalysisMap {
			if db.isMateIn0246(a.board, dtm) >= 0 {
				continue
			}
			p := NewPosition(a.board, player)
			moves := GenerateMoves(p)

			found := 0
			minDTM := 500
			for _, m := range moves {
				newBoard := a.board.DoMove(m)
				newDtm := db.isMateIn1357(newBoard, dtm)
				if newDtm < minDTM {
					minDTM = newDtm
				}
				if db.isMateIn1357(newBoard, dtm) >= 0 {
					found++
				}
			}

			if found == len(moves) {
				for _, m := range moves {
					f := db.addDTMToAnalysis(a.board, minDTM+1, m)
					if f {
						newMovesFound++
					}
				}
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
			fmt.Printf("db.dtmDb[%2d] %6d/%6d\n", i, dtms, moves)
		}
		fmt.Printf("\nnewMovesFound %d\n", newMovesFound)
		fmt.Printf("duration %v\n\n", end.Sub(start))
	}

	if newMovesFound == 0 {
		return errNowNewAnalysis
	}
	return noError
}
func (db *EndGameDb) isMateIn0246(board *Board, maxDtm int) int {
	if a, ok := db.AnalysisMap[board.String()]; ok {
		for dtm := 0; dtm < maxDtm; dtm += 2 {
			if a.dtm == dtm {
				return dtm
			}
		}
	}
	return -1
}
func (db *EndGameDb) isMateIn1357(board *Board, maxDtm int) int {
	if a, ok := db.AnalysisMap[board.String()]; ok {
		for dtm := 1; dtm < maxDtm; dtm += 2 {
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
