package emil

import (
	"errors"
	"fmt"
)

var errKingCapture = errors.New("king capture")

func generateMoveList(b *Board, player int) (list []*Move, err error) {
	var empty []*Move
	for src, piece := range b.squares {
		if isOwnPiece(player, piece) {
			switch abs(piece) {
			case kingValue:
				for _, dst := range kingDestinationsFrom(src) {
					capture := b.squares[dst]
					if isKing(capture) {
						return empty, errKingCapture
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
						if isKing(capture) {
							return empty, errKingCapture
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
	return list, err
}

func filterKingCaptures(b *Board, player int, list []*Move) (result []*Move) {
	for _, m := range list {
		if DEBUG {
			println("TEST move", m.String())
		}
		b.doMove(m)
		_, kingCaptured := generateMoveList(b, otherPlayer(m.player))
		if kingCaptured == nil {
			result = append(result, m)
		} else {
			if DEBUG {
				println("KingCaptured")
			}
		}
		b.undoMove(m)
		if DEBUG {
			fmt.Printf("\n\n\n")
		}
	}
	return result
}

//Search best move for player on board
func Search(b *Board, player int) (bestMove *Move) {
	bestMove, _ = deepSearch(b, player, 0, 1)
	return bestMove
}

//Search best move for player on board
func deepSearch(b *Board, player, deep, maxDeep int) (bestMove *Move, bestScore int) {
	if DEBUG {
		fmt.Printf("deepSearch: %s deep:%d %d\n", players[player], deep, maxDeep)
		fmt.Printf("%s\n", b)
	}
	if player == WHITE {
		bestScore = 2 * BlackKing
	} else {
		bestScore = 2 * WhiteKing
	}

	if deep == maxDeep {
		score := evaluate(b)
		if DEBUG {
			println("maxDeep Score:", score)
		}
		return nil, score
	}

	list, err := generateMoveList(b, player)
	if err != nil {
		return nil, 0
	}

	result := filterKingCaptures(b, player, list)

	for _, m := range result {
		b.doMove(m)
		_, score := deepSearch(b, otherPlayer(player), deep+1, maxDeep)
		b.undoMove(m)

		if (player == WHITE && score > bestScore) || (player == BLACK && score < bestScore) {
			bestScore = score
			bestMove = m
		}
	}
	return bestMove, bestScore
}

func evaluate(b *Board) (score int) {
	score = 0
	for _, p := range b.squares {
		score += p
	}
	return score
}
