package emil

import (
	"fmt"
)

type analysis struct {
	mateIn int
	board  *Board
}

type endGameDb struct {
	positionDb map[string]*analysis
}

const unknown = -1

func (db *endGameDb) addPosition(board *Board) {
	a := &analysis{mateIn: unknown, board: board}
	db.positionDb[board.String()] = a
}

func (db *endGameDb) positions() int {
	return len(db.positionDb)
}

func NewEndGameDb() {
	var err error
	fmt.Printf("Generating all possible positions for KRK\n")

	endGames := &endGameDb{positionDb: make(map[string]*analysis)}

	DEBUG = false
	for wk := A1; wk <= H8; wk++ {
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

				err = board.KingsToClose(wk, bk)
				if err != nil {
					continue
				}

				endGames.addPosition(board)
			}
		}
	}
	fmt.Printf("all positions %d\n", 64*63*62)
	fmt.Printf("endGames.positions() %d\n", endGames.positions())
	fmt.Printf("difference %d\n", 64*63*62-endGames.positions())

}
