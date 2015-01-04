package emil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type positionKey string

// PositionEntry is an entry in the PositionDb
type PositionEntry struct {
	Position      *position
	Dtm           int
	PrevPositions map[positionKey]*Move
	NextPositions map[positionKey]*Move
}

// NewPositionEntry ceates a new *PositionEntry
func NewPositionEntry(p *position) *PositionEntry {
	return &PositionEntry{
		Position:      p,
		Dtm:           initial,
		PrevPositions: make(map[positionKey]*Move),
		NextPositions: make(map[positionKey]*Move)}
}

func (entry *PositionEntry) addMoveToNextPosition(next *position, m *Move) {
}

// PositionDb to query for mate in 1,2, etc.
type PositionDb struct {
	Positions map[positionKey]*PositionEntry
}

func (db *PositionDb) addPosition(p *position) {
	if _, ok := db.Positions[p.key()]; ok {
		panic("key exsists in db " + p.key())
	}
	entry := NewPositionEntry(p)
	db.retrogradeAnalysisStep0(entry)
	db.Positions[p.key()] = entry
}

func (db *PositionDb) AddPrevPositions() {
	for key, entry := range db.Positions {
		for nextKey, moveToNext := range entry.NextPositions {
			nextPosition := PositionFromKey(string(nextKey))
			nextEntry, ok := db.Positions[nextPosition.key()]
			if ok {
				nextEntry.PrevPositions[key] = moveToNext
			}
		}
	}
}

// generate all moves
func (db *PositionDb) retrogradeAnalysisStep0(entry *PositionEntry) {
	moves := GenerateMoves(entry.Position)
	other := otherPlayer(entry.Position.player)
	for _, move := range moves {
		nextBoard := entry.Position.board.DoMove(move)
		nextPosition := NewPosition(nextBoard, other)
		entry.NextPositions[nextPosition.key()] = move
	}
}

// NewPositionDB creates a new *PositionDB
func NewPositionDB() *PositionDb {
	return &PositionDb{
		Positions: make(map[positionKey]*PositionEntry)}
}

func (db *PositionDb) FillWithKRKPositions() {
	var err error

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

				db.addPosition(NewPosition(board, WHITE))
				db.addPosition(NewPosition(board, BLACK))
			}
		}
	}
}

// SaveEndGameDb saves the an end game DB for KRK to file
func (db *PositionDb) SavePositionDb(file string) error {
	fmt.Println("WriteDataToFile: ", file)

	start := time.Now()
	fmt.Printf("json.MarshalIndent\n")
	b, err := json.MarshalIndent(db, "", "  ")
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