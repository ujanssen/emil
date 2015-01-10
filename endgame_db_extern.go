package emil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

const filename = "AnalysisMap.json"

var IN_TEST = false

type EndGameSave struct {
	AnalysisMap map[string]string `json:"analysis"`
}

// EndGameDb to query for mate in 1,2, etc.
type EndGameDb struct {
	Start       time.Time
	Duration    time.Duration
	AnalysisMap map[string]*Analysis
	AnalysisStr map[string]string
}

func (db *EndGameDb) Find(p *position) (bestMove *Move) {
	if DEBUG {
		fmt.Printf("Find:\n%s\n", p.Board)
	}
	a := db.AnalysisMap[p.Board.String()]
	if DEBUG {
		fmt.Printf("Found: AnalysisMap with dtms %v\n", a.DTMs(p.Player))
	}
	return a.BestMove(p.Player)
}

func (db *EndGameDb) FindMatesIn(dtm int) (as []*Analysis) {
	for _, a := range db.AnalysisMap {
		if a.dtm == dtm {
			as = append(as, a)
		}
	}
	return as
}

func (db *EndGameDb) FindMates() (as []*Analysis) {
	return db.FindMatesIn(0)
}

func (db *EndGameDb) FindMate(piece, square int) (boards []*Board) {
	for _, a := range db.AnalysisMap {
		if a.dtm == 0 {
			if a.board.Squares[square] == piece {
				boards = append(boards, a.board)
			}
		}
	}
	return boards
}

func (db *EndGameDb) CreateAnalysisStr() {
	db.AnalysisStr = make(map[string]string)
	for k, a := range db.AnalysisMap {
		db.AnalysisStr[k] = fmt.Sprintf("%v.%v", a.dtmWhite, a.dtmBlack)
	}
}

func GenerateMoves(p *position) (list []*Move) {
	for _, m := range generateMoves(p) {
		b := p.Board.DoMove(m)
		if !isKingInCheck(NewPosition(b, p.Player)) {
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
		AnalysisMap: make(map[string]*Analysis)}

	for fen, v := range data.AnalysisMap {
		board := Fen2Board(fen)
		db.addAnalysisFromStr(board, v)
	}
	db.retrogradeAnalysis()

	return db, err
}

// SaveEndGameDb saves the an end game DB for KRK to file
func SaveEndGameDb(file string, analysisStr map[string]string) error {
	fmt.Println("WriteDataToFile: ", file)

	data := EndGameSave{}
	data.AnalysisMap = analysisStr
	start := time.Now()
	fmt.Printf("json.MarshalIndent\n")
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	end := time.Now()
	fmt.Printf("json.MarshalIndent %v\n", end.Sub(start))

	start = time.Now()
	fmt.Printf("ioutil.WriteFile\n")
	err = ioutil.WriteFile(file, b, 0666)
	end = time.Now()
	fmt.Printf("ioutil.WriteFile %v, error=%v\n", end.Sub(start), err)
	return err
}

// NewEndGameDb generates an end game DB for KRK
func NewEndGameDb() *EndGameDb {
	var err error

	db := &EndGameDb{
		Start:       time.Now(),
		AnalysisMap: make(map[string]*Analysis),
		AnalysisStr: make(map[string]string)}

	if DEBUG {
		fmt.Printf("create all position and moves\n")
	}
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

				if squaresDistances[wk][bk] <= 1 {
					continue
				}
				db.addAnalysis(board)
			}
		}
	}
	end := time.Now()
	db.Duration = end.Sub(db.Start)
	if DEBUG {
		fmt.Printf("create all position and moves duration %v\n", db.Duration)
	}

	return db
}
