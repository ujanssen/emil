package emil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

const filename = "EndGameDb.json"

var IN_TEST = false

type EndGameSave struct {
	AnalysisMap map[string]string `json:"analysis"`
}

// EndGameDb to query for mate in 1,2, etc.
type EndGameDb struct {
	Start       time.Time
	Duration    time.Duration
	AnalysisMap map[string]*Analysis
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

func (db *EndGameDb) MaxDtm() int {
	return len(db.dtmDb)
}

func GenerateMoves(p *position) (list []*Move) {
	for _, m := range generateMoves(p) {
		b := p.board.DoMove(m)
		if !isKingInCheck(NewPosition(b, p.player)) {
			list = append(list, m)
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
		db.addAnalysis(board)
		dtms := DTMsFromString(v)
		for _, d := range dtms {
			db.addDTMToAnalysis(board, d.dtm, d.move.reverse())
		}
	}

	return db, err
}

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

	db := &EndGameDb{
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

				db.addAnalysis(board)
			}
		}
	}
	end := time.Now()
	db.Duration = end.Sub(db.Start)
	if DEBUG {
		positions := len(db.AnalysisMap)

		fmt.Printf("all positions %d\n", 64*63*62)
		fmt.Printf("endGames.positions() %d\n", positions)
		fmt.Printf("difference %d\n", 64*63*62-positions)
		fmt.Printf("duration %v\n", db.Duration)
	}
	db.retrogradeAnalysis()

	return db
}
