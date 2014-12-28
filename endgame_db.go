package emil

import (
	"errors"
	"fmt"
	"time"
)

var errNowNewAnalysis = errors.New("errNowNewAnalysis")

type analysis struct {
	analysisDone bool
	dtm          int // Depth to mate

	board *Board
	move  *Move
}

// EndGameDb to query for mate in 1,2, etc.
type EndGameDb struct {
	positionDb map[string]*analysis

	dtmDb []map[string]bool

	searchedPositions int
}

func (db *EndGameDb) Find(board *Board) (bestMove *Move) {
	if DEBUG {
		fmt.Printf("Find:\n%s\n", board.String())
	}
	a := db.positionDb[board.String()]
	if DEBUG {
		fmt.Printf("Found: positionDb with dtm %d\n", a.dtm)
	}
	return a.move
}
func (db *EndGameDb) FindMates() (boards []*Board) {
	for str := range db.dtmDb[0] {
		a := db.positionDb[str]
		boards = append(boards, a.board)
	}
	return boards
}

func (db *EndGameDb) FindMate(piece, square int) (boards []*Board) {
	for str := range db.dtmDb[0] {
		a := db.positionDb[str]
		if a.board.squares[square] == piece {
			boards = append(boards, a.board)
		}
	}
	return boards
}

func (db *EndGameDb) addPosition(board *Board) {
	db.addAnalysis(board, -1, nil)
}

func (db *EndGameDb) addAnalysis(board *Board, dtm int, move *Move) {
	a := &analysis{
		dtm:   dtm,
		board: board}
	if move != nil {
		a.move = move.reverse()
	}
	done := dtm >= 0
	if done {
		db.dtmDb[dtm][a.board.String()] = true
	}
	a.analysisDone = done

	db.positionDb[a.board.String()] = a
}

func (db *EndGameDb) positions() int {
	return len(db.positionDb)
}

// find positions where black is checkmate
func (db *EndGameDb) retrogradeAnalysisStep1() {
	db.dtmDb = append(db.dtmDb, make(map[string]bool))

	start := time.Now()

	player := BLACK
	for boardStr, a := range db.positionDb {
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
		db.searchedPositions++
		if move == nil {
			if isKingInCheck(p) {
				a.dtm = 0
				db.addAnalysis(a.board, 0, nil)
				if DEBUG {
					fmt.Printf("mate:\n%s\n", boardStr)
				}
			}
		}
	}
	end := time.Now()
	if DEBUG {
		fmt.Printf("searchedPositions %d\n", db.searchedPositions)
		fmt.Printf("db.dtmDb[0] %d\n", len(db.dtmDb[0]))
		fmt.Printf("duration %v\n\n\n", end.Sub(start))
	}
}
func playerForStepN(dtm int) (player int) {
	if dtm%2 == 0 {
		player = BLACK
		if DEBUG {
			fmt.Printf("retrogradeAnalysis BLACK %d\n", dtm)
		}
	} else {
		player = WHITE
		if DEBUG {
			fmt.Printf("retrogradeAnalysis WHITE %d\n", dtm)
		}
	}
	return player
}

func (db *EndGameDb) retrogradeAnalysisStepN(dtm int) (noError error) {
	start := time.Now()
	db.dtmDb = append(db.dtmDb, make(map[string]bool))

	player := playerForStepN(dtm)

	if DEBUG {
		fmt.Printf("Start positions %d\n", len(db.dtmDb[dtm-1]))
	}
	for str := range db.dtmDb[dtm-1] {
		a := db.positionDb[str]
		p := NewPosition(a.board, player)
		list := generateMoves(p)
		moves := filterKingCaptures(p, list)
		moves = filterKingCaptures(NewPosition(a.board, otherPlayer(player)), list)
		for _, m := range moves {
			newBoard := a.board.doMove(m)
			newAnalysis, ok := db.positionDb[newBoard.String()]
			if ok && !newAnalysis.analysisDone {
				db.addAnalysis(newBoard, dtm, m)
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

func (db *EndGameDb) retrogradeAnalysis() {
	// find positions where black is checkmate
	db.retrogradeAnalysisStep1()
	dtm := 1
	for {
		err := db.retrogradeAnalysisStepN(dtm)
		if err != nil {
			break
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

// NewEndGameDb generates an end game DB for KRK
func NewEndGameDb() *EndGameDb {
	var err error
	start := time.Now()

	endGames := &EndGameDb{
		positionDb: make(map[string]*analysis),
		dtmDb:      make([]map[string]bool, 0)}

	for wk := A1; wk <= H8; wk++ {
		//for wk := E3; wk <= E3; wk++ {
		if DEBUG {
			fmt.Printf("White king on %s\n", BoardSquares[wk])
		}
		for wr := A1; wr <= H8; wr++ {
			for bk := A1; bk <= H8; bk++ {

				board := NewBoard()

				err = board.Setup(WhiteKing, wk)
				if err != nil {
					continue
				}

				err = board.Setup(WhiteRock, wr)
				if err != nil {
					continue
				}

				err = board.Setup(BlackKing, bk)
				if err != nil {
					continue
				}

				err = board.kingsToClose()
				if err != nil {
					continue
				}

				endGames.addPosition(board)
			}
		}
	}
	end := time.Now()
	if DEBUG {
		fmt.Printf("all positions %d\n", 64*63*62)
		fmt.Printf("endGames.positions() %d\n", endGames.positions())
		fmt.Printf("difference %d\n", 64*63*62-endGames.positions())
		fmt.Printf("duration %v\n", end.Sub(start))
	}
	endGames.retrogradeAnalysis()

	return endGames
}
