package emil

import (
	"bytes"
	"encoding/gob"
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

func (db *PositionDb) addPositions(board *Board) {
	db.addPosition(NewPosition(board, WHITE))
	db.addPosition(NewPosition(board, BLACK))
}

func (db *PositionDb) addPosition(p *position) {
	if _, ok := db.Positions[p.key()]; ok {
		panic("key exsists in db " + string(p.key()))
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

// generate NextPositions
func (db *PositionDb) retrogradeAnalysisStep0(entry *PositionEntry) {
	moves := GenerateMoves(entry.Position)
	other := otherPlayer(entry.Position.Player)
	for _, move := range moves {
		nextBoard := entry.Position.Board.DoMove(move)
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
	board := NewBoard()

	for wk := A1; wk <= H8; wk++ {
		if err = board.Setup(WhiteKing, wk); err != nil {
			continue
		}

		start := time.Now()
		for bk := A1; bk <= H8; bk++ {
			if squaresDistances[wk][bk] <= 1 {
				continue
			}
			if err = board.Setup(BlackKing, bk); err != nil {
				continue
			}
			// no rock
			for wr := A1; wr <= H8; wr++ {
				if err = board.Setup(WhiteRock, wr); err != nil {
					continue
				}

				db.addPositions(board)
				board.Empty(wr)
			}
			db.addPositions(board)
			board.Empty(bk)

		}
		end := time.Now()
		if DEBUG {
			fmt.Printf("White king on %s, %v\n", BoardSquares[wk], end.Sub(start))
		}
		board.Empty(wk)
	}

}

// SaveEndGameDb saves the an end game DB for KRK to file
func (db *PositionDb) SavePositionDb(file string) error {
	var network bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&network) // Will write to network.

	fmt.Println("WriteDataToFile: ", file)

	start := time.Now()
	fmt.Printf("enc.Encode\n")

	// Encode (send) some values.
	err := enc.Encode(db)
	if err != nil {
		panic("encode error:" + err.Error())
	}

	end := time.Now()
	fmt.Printf("enc.Encode %v\n", end.Sub(start))

	start = time.Now()
	fmt.Printf("ioutil.WriteFile\n")
	err = ioutil.WriteFile(file, network.Bytes(), 0666)
	end = time.Now()
	fmt.Printf("ioutil.WriteFile %v, error=%v\n", end.Sub(start), err)
	return err
}
func LoadPositionDb(file string) (db *PositionDb, err error) {
	fmt.Println("LoadDataFromFile: ", file)

	start := time.Now()
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return db, err
	}
	end := time.Now()
	fmt.Printf("ioutil.ReadFile %v,b=%d, error=%v\n", end.Sub(start), len(b), err)

	data := NewPositionDB()
	start = time.Now()
	err = json.Unmarshal(b, data)
	if err != nil {
		return db, err
	}
	end = time.Now()
	fmt.Printf("json.Unmarshal%v\n", end.Sub(start))
	return data, err
}
