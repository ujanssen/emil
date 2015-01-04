package emil

import (
	"fmt"
	"time"
)

type positionKey string

// PositionEntry is an entry in the PositionDb
type PositionEntry struct {
	Position      *position
	Dtm           int
	PrevPositions map[positionKey]*PositionEntry
	NextPositions map[positionKey]*PositionEntry
}

// NewPositionEntry ceates a new *PositionEntry
func NewPositionEntry(p *position) *PositionEntry {
	return &PositionEntry{
		Position:      p,
		Dtm:           initial,
		PrevPositions: make(map[positionKey]*PositionEntry),
		NextPositions: make(map[positionKey]*PositionEntry)}
}

// PositionDb to query for mate in 1,2, etc.
type PositionDb struct {
	positions map[positionKey]*PositionEntry
}

func (db *PositionDb) Positions() int {
	return len(db.positions)
}

func (db *PositionDb) addPosition(p *position) {
	if _, ok := db.positions[p.key()]; ok {
		panic("key exsists in db " + p.key())
	}
	db.positions[p.key()] = NewPositionEntry(p)
}

// NewPositionDB creates a new *PositionDB
func NewPositionDB() *PositionDb {
	return &PositionDb{
		positions: make(map[positionKey]*PositionEntry)}
}

func (db *PositionDb) FillWithKRKPositions() {
	var err error
	start := time.Now()

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
	end := time.Now()
	duration := end.Sub(start)
	if DEBUG {
		fmt.Printf("create all position and moves duration %v\n", duration)
	}
}
