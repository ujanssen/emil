package emil

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

var errPositionNotFound = errors.New("Position not found")

type PositionKey string

// PositionEntry is an entry in the PositionDb
type PositionEntry struct {
	Position      *Position
	Dtm           int
	PrevPositions map[PositionKey]*Move
	NextPositions map[PositionKey]*Move
}

// NewPositionEntry ceates a new *PositionEntry
func NewPositionEntry(p *Position) *PositionEntry {
	return &PositionEntry{
		Position:      p,
		Dtm:           initial,
		PrevPositions: make(map[PositionKey]*Move),
		NextPositions: make(map[PositionKey]*Move)}
}

// PositionDb to query for mate in 1,2, etc.
type PositionDb struct {
	Positions map[PositionKey]*PositionEntry
}

func (db *PositionDb) addPositions(board *Board) {
	db.addPosition(board, WHITE)
	db.addPosition(board, BLACK)
}

func (db *PositionDb) addPosition(board *Board, player int) {
	p := NewPosition(board, player)
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
		Positions: make(map[PositionKey]*PositionEntry)}
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
	dataFile, err := os.Open(file)
	if err != nil {
		panic("decode error " + err.Error())
	}

	dec := gob.NewDecoder(dataFile)
	data := NewPositionDB()
	if err = dec.Decode(data); err != nil {
		panic("decode error " + err.Error())
	}
	dataFile.Close()
	return data, err
}

// FindWhitePosition serves with rpc
func (db *PositionDb) FindWhitePosition(fen string, result *PositionEntry) error {
	board := Fen2Board(fen)
	p := NewPosition(board, WHITE)
	pe, ok := db.Positions[p.key()]
	if !ok {
		log.Printf("Position %s not found\n", fen)
		return errPositionNotFound
	}
	*result = *pe
	return nil
}
