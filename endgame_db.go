package emil

import (
	"fmt"
	"time"
)

type analysis struct {
	dtm   int // Depth to mate
	board *Board
	move  *Move
}

// EndGameDb to query for mate in 1,2, etc.
type EndGameDb struct {
	positionDb map[string]*analysis

	retros []map[string]*analysis

	pattIn0 int

	searchedPositions int
}

const unknown = -1
const patt = -2

func (db *EndGameDb) Find(board *Board) (bestMove *Move) {
	if DEBUG {
		fmt.Printf("Find:\n%s\n", board.String())
	}
	if a, ok := db.retros[0][board.String()]; ok {
		if DEBUG {
			fmt.Printf("Found: retros with dtm %d\n", a.dtm)
		}
		if a.move != nil {
			return a.move
		}
	}
	if a, ok := db.positionDb[board.String()]; ok {
		if DEBUG {
			fmt.Printf("Found: positionDb with dtm %d\n", a.dtm)
		}
		if a.move != nil {
			return a.move
		}
	}
	return nil
}

func (db *EndGameDb) addPosition(board *Board) {
	a := &analysis{
		dtm:   unknown,
		board: board}
	db.positionDb[board.String()] = a
}

func (db *EndGameDb) addAnalysis(board *Board, dtm int, move *Move) {
	a := &analysis{
		dtm:   dtm,
		board: board}
	if move != nil {
		a.move = move.reverse()
	}
	db.positionDb[a.board.String()] = a
	db.retros[dtm][a.board.String()] = a
}

func (db *EndGameDb) positions() int {
	return len(db.positionDb)
}

func (db *EndGameDb) retrogradeAnalysis() {
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
				db.addAnalysis(a.board, a.dtm, nil)
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
	if DEBUG {
		fmt.Printf("searchedPositions %d\n", db.searchedPositions)
		fmt.Printf("db.retros[0] %d\n", len(db.retros[0]))
		fmt.Printf("found patt in 0 %d\n", db.pattIn0)

		fmt.Printf("duration %v\n", end.Sub(start))
	}
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
		moves = filterKingCaptures(a.board, otherPlayer(player), list)
		for i, m := range moves {
			newBoard := a.board.doMove(m)
			newBoardStr := newBoard.String()
			if _, ok := db.retros[0][newBoardStr]; ok {
				if DEBUG {
					fmt.Printf("move[%d/%d]: %s found in db.retros[0]\n", i+1, len(moves), m.String())
				}
				continue // new position is checkmate
			}
			db.addAnalysis(newBoard, 1, m)

			if DEBUG {
				fmt.Printf("move[%d/%d]: %s added to db.retros[1]\n", i+1, len(moves), m.String())
			}
		}
	}
	end = time.Now()

	if DEBUG {
		fmt.Printf("db.retros[1] %d\n", len(db.retros[1]))
		fmt.Printf("duration %v\n", end.Sub(start))
	}
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
func NewEndGameDb() *EndGameDb {
	var err error
	start := time.Now()

	endGames := &EndGameDb{
		positionDb: make(map[string]*analysis),
		retros:     make([]map[string]*analysis, 0)}

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
	if DEBUG {
		fmt.Printf("all positions %d\n", 64*63*62)
		fmt.Printf("endGames.positions() %d\n", endGames.positions())
		fmt.Printf("difference %d\n", 64*63*62-endGames.positions())
		fmt.Printf("duration %v\n", end.Sub(start))
	}
	endGames.retrogradeAnalysis()

	return endGames
}
