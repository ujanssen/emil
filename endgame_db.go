package emil

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"time"
)

var IN_TEST = false
var errNowNewAnalysis = errors.New("errNowNewAnalysis")

type EndGameSave struct {
	AnalysisMap map[string]string `json:"analysis"`
}

// EndGameDb to query for mate in 1,2, etc.
type EndGameDb struct {
	Start       time.Time            `json:"startTime"`
	Duration    time.Duration        `json:"duration"`
	AnalysisMap map[string]*Analysis `json:"analysis"`
	dtmDb       []map[string]bool
}

func (db *EndGameDb) Find(p *position) (bestMove *Move) {
	if DEBUG {
		fmt.Printf("Find:\n%s\n", p.board)
	}
	a := db.AnalysisMap[p.board.String()]
	if DEBUG {
		fmt.Printf("Found: AnalysisMap with dtms %v\n", a.DTMs(p.player))
	}
	return a.BestMove(p.player)
}

func (db *EndGameDb) FindMatesIn(dtm int) (as []*Analysis) {
	if dtm == -1 {
		for _, a := range db.AnalysisMap {
			if a.playerHaveDTMs() {
				as = append(as, a)
			}
		}
	} else {
		for str := range db.dtmDb[dtm] {
			as = append(as, db.AnalysisMap[str])
		}
	}
	return as
}

func (db *EndGameDb) FindMates() (as []*Analysis) {
	return db.FindMatesIn(0)
}

func (db *EndGameDb) FindMate(piece, square int) (boards []*Board) {
	for str := range db.dtmDb[0] {
		a := db.AnalysisMap[str]
		if a.board.squares[square] == piece {
			boards = append(boards, a.board)
		}
	}
	return boards
}

func (db *EndGameDb) addPosition(board *Board) {
	a := NewAnalysis(board)
	db.AnalysisMap[a.board.String()] = a
}

func (db *EndGameDb) addAnalysis(board *Board, dtm int, move *Move) {
	a := db.AnalysisMap[board.String()]
	if move != nil {
		a.addDTM(move.reverse(), dtm)
	}
	if dtm >= 0 {
		if move != nil {
			playerForStep := playerForStepN(dtm)
			if playerForStep != move.player {
				panic("playerForStep != move.player")
			}
		}
		db.dtmDb[dtm][board.String()] = true
	}
}

func (db *EndGameDb) positions() int {
	return len(db.AnalysisMap)
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
				db.addAnalysis(a.board, 0, nil)
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
				db.addAnalysis(newBoard, dtm, m)
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
					db.addAnalysis(a.board, maxDTM+1, m)
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

func (db *EndGameDb) MaxDtm() int {
	return len(db.dtmDb)
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
func GenerateMoves(p *position) (list []*Move) {
	for _, m := range generateMoves(p) {
		b := p.board.DoMove(m)
		if !IsTheKingInCheck(NewPosition(b, WHITE)) {
			list = append(list, m)
		}
	}
	return list
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
func LoadEndGameDb() (db *EndGameDb, err error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return db, err
	}
	var data EndGameSave
	err = json.Unmarshal(b, &data)
	if err != nil {
		return db, err
	}
	db = &EndGameDb{
		Start:       time.Now(),
		AnalysisMap: make(map[string]*Analysis),
		dtmDb:       make([]map[string]bool, 0)}

	for i := 0; i < 34; i++ {
		db.dtmDb = append(db.dtmDb, make(map[string]bool))
	}

	for fen, v := range data.AnalysisMap {
		board := Fen2Board(fen)
		db.addPosition(board)
		dtms := DTMsFromString(v)
		for _, d := range dtms {
			db.addAnalysis(board, d.dtm, d.move.reverse())
		}
	}

	return db, err
}

const filename = "EndGameDb.json"

// SaveEndGameDb saves the an end game DB for KRK to file
func SaveEndGameDb(db *EndGameDb) error {
	fmt.Println("WriteDataToFile: ", filename)

	data := EndGameSave{AnalysisMap: make(map[string]string)}

	for p, a := range db.AnalysisMap {
		data.AnalysisMap[p] = fmt.Sprintf("%v", a.dtmWhite)
	}

	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filename, b, 0666)
}

// NewEndGameDb generates an end game DB for KRK
func NewEndGameDb() *EndGameDb {
	var err error

	endGames := &EndGameDb{
		Start:       time.Now(),
		AnalysisMap: make(map[string]*Analysis),
		dtmDb:       make([]map[string]bool, 0)}

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
	endGames.Duration = end.Sub(endGames.Start)
	if DEBUG {
		fmt.Printf("all positions %d\n", 64*63*62)
		fmt.Printf("endGames.positions() %d\n", endGames.positions())
		fmt.Printf("difference %d\n", 64*63*62-endGames.positions())
		fmt.Printf("duration %v\n", endGames.Duration)
	}
	endGames.retrogradeAnalysis()

	return endGames
}
