package emil

import (
	"fmt"
	"time"
)

type analysis struct {
	dtm   int // Depth to mate
	board *Board
	moves []*Move
}

type endGameDb struct {
	positionDb map[string]*analysis

	retros []map[string]*analysis

	pattIn0 int

	searchedPositions int
}

const unknown = -1
const patt = -2

func (db *endGameDb) addPosition(board *Board) {
	a := &analysis{
		dtm:   unknown,
		board: board,
		moves: make([]*Move, 0)}
	db.positionDb[board.String()] = a
}

func (db *endGameDb) addAnalysis(a *analysis) {
	db.retros[a.dtm][a.board.String()] = a
}

func (db *endGameDb) positions() int {
	return len(db.positionDb)
}

func (db *endGameDb) retrogradeAnalysis() {
	// find positions where black is checkmate
	db.retros = append(db.retros, make(map[string]*analysis))

	start := time.Now()

	player := BLACK
	for boardStr, a := range db.positionDb {
		if a.dtm > unknown {
			continue
		}
		// mate only on border square
		blackKingSquare := BoardSquares[a.board.blackKing]
		if !blackKingSquare.isBorder {
			continue
		}
		// mate only with help from king
		if squaresDistances[a.board.blackKing][a.board.whiteKing] > 2 {
			continue
		}

		move := Search(a.board, player)
		db.searchedPositions++
		if move == nil {
			if isKingInCheck(a.board, player) {
				a.dtm = 0
				db.addAnalysis(a)
				if DEBUG {
					fmt.Printf("mate:\n%s\n", boardStr)
				}
			} else {
				a.dtm = patt
				db.pattIn0++
				if DEBUG {
					fmt.Printf("patt:\n%s\n", boardStr)
				}
			}
		}
	}
	end := time.Now()

	fmt.Printf("searchedPositions %d\n", db.searchedPositions)
	fmt.Printf("db.retros[0] %d\n", len(db.retros[0]))
	fmt.Printf("found patt in 0 %d\n", db.pattIn0)

	fmt.Printf("duration %v\n", end.Sub(start))

	//Suche alle Stellungen, bei denen Weiß am Zug ist und
	//Weiß mindestens einen Zug hat, der zu einer Stellung unter 1. führt.
	//Das sind alle Stellungen, in denen Weiß mit einem Zug matt setzen kann.
	//Markiere diese Stellungen in der Datei.
	start = time.Now()
	db.retros = append(db.retros, make(map[string]*analysis))
	player = WHITE
	for _, a := range db.retros[0] {
		list := generateMoves(a.board, player)
		moves := filterKingCaptures(a.board, player, list)
		for i, m := range moves {
			a.board.doMove(m)
			newBoardStr := a.board.String()
			if _, ok := db.retros[0][newBoardStr]; ok {
				fmt.Printf("move[%d/%d]: %s found in db.retros[0]\n", i+1, len(moves), m.String())
				continue // new position is checkmate
			}
			a.dtm = 1
			a.moves = append(a.moves, m)
			db.retros[1][newBoardStr] = a
			fmt.Printf("move[%d/%d]: %s added to db.retros[1]\n", i+1, len(moves), m.String())

			a.board.undoMove(m)
		}
	}
	end = time.Now()

	fmt.Printf("db.retros[1] %d\n", len(db.retros[1]))
	fmt.Printf("duration %v\n", end.Sub(start))
}

func generateMoves(b *Board, player int) (list []*Move) {
	for src, piece := range b.squares {
		if isOwnPiece(player, piece) {
			switch abs(piece) {
			case kingValue:
				for _, dst := range kingDestinationsFrom(src) {
					capture := b.squares[dst]
					if isOtherKing(player, capture) {
						continue
					}
					if capture == Empty {
						list = append(list, newSilentMove(player, piece, src, dst))
					} else if !isOwnPiece(player, capture) {
						list = append(list, newCaptureMove(player, piece, capture, src, dst))
					}
				}
			case rockValue:
				for _, dsts := range rockDestinationsFrom(src) {
					for _, dst := range dsts {
						capture := b.squares[dst]
						if isOtherKing(player, capture) {
							break
						}
						if capture == Empty {
							list = append(list, newSilentMove(player, piece, src, dst))
						} else if !isOwnPiece(player, capture) {
							list = append(list, newCaptureMove(player, piece, capture, src, dst))
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
func NewEndGameDb() {
	var err error
	start := time.Now()
	fmt.Printf("Generating all possible positions for KRK\n")

	endGames := &endGameDb{
		positionDb: make(map[string]*analysis),
		retros:     make([]map[string]*analysis, 0)}

	DEBUG = true
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
	fmt.Printf("all positions %d\n", 64*63*62)
	fmt.Printf("endGames.positions() %d\n", endGames.positions())
	fmt.Printf("difference %d\n", 64*63*62-endGames.positions())
	fmt.Printf("duration %v\n", end.Sub(start))

	endGames.retrogradeAnalysis()
}
