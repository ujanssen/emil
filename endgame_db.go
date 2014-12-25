package emil

import (
	"fmt"
	"time"
)

type analysis struct {
	mateIn int
	board  *Board
}

type endGameDb struct {
	positionDb map[string]*analysis

	mateIn0 int
	pattIn0 int

	searchedPositions int
}

const unknown = -1
const patt = -2

func (db *endGameDb) addPosition(board *Board) {
	a := &analysis{mateIn: unknown, board: board}
	db.positionDb[board.String()] = a
}

func (db *endGameDb) positions() int {
	return len(db.positionDb)
}

func (db *endGameDb) retrogradeAnalysis() {
	// find positions where black is checkmate

	start := time.Now()

	player := BLACK
	for boardStr, analysis := range db.positionDb {
		if analysis.mateIn > unknown {
			continue
		}
		// mate only on border square
		blackKingSquare := BoardSquares[analysis.board.blackKing]
		if !blackKingSquare.isBorder {
			continue
		}
		// mate only with help from king
		if squaresDistances[analysis.board.blackKing][analysis.board.whiteKing] > 2 {
			continue
		}

		move := Search(analysis.board, player)
		db.searchedPositions++
		if move == nil {
			if isKingInCheck(analysis.board, player) {
				analysis.mateIn = 0
				db.mateIn0++
				if DEBUG {
					fmt.Printf("mate:\n%s\n", boardStr)
				}
			} else {
				analysis.mateIn = patt
				db.pattIn0++
				if DEBUG {
					fmt.Printf("patt:\n%s\n", boardStr)
				}
			}
		}
	}
	end := time.Now()

	fmt.Printf("searchedPositions %d\n", db.searchedPositions)
	fmt.Printf("found mate in 0 %d\n", db.mateIn0)
	fmt.Printf("found patt in 0 %d\n", db.pattIn0)

	fmt.Printf("duration %v\n", end.Sub(start))
}

// NewEndGameDb generates an end game DB for KRK
func NewEndGameDb() {
	var err error
	start := time.Now()
	fmt.Printf("Generating all possible positions for KRK\n")

	endGames := &endGameDb{positionDb: make(map[string]*analysis)}

	DEBUG = true
	for wk := A1; wk <= H8; wk++ {
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
	fmt.Printf("all positions %d\n", 64*63*62)
	fmt.Printf("endGames.positions() %d\n", endGames.positions())
	fmt.Printf("difference %d\n", 64*63*62-endGames.positions())
	fmt.Printf("duration %v\n", end.Sub(start))

	endGames.retrogradeAnalysis()
}
