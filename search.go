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
					if isOtherKing(player, capture) {
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
						if isOtherKing(player, capture) {
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
func isKingInCheck(b *Board, player int) (kingInCheck bool) {
	_, kingCaptured := generateMoveList(b, otherPlayer(player))
	if kingCaptured != nil {
		return true
	}
	return false

}
func filterKingCaptures(b *Board, player int, list []*Move) (result []*Move) {
	for _, m := range list {
		b.doMove(m)
		if !isKingInCheck(b, player) {
			result = append(result, m)
		}
		b.undoMove(m)
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
	if player == WHITE {
		bestScore = 2 * BlackKing
	} else {
		bestScore = 2 * WhiteKing
	}

	if deep == maxDeep {
		score := evaluate(b)
		if DEBUG {
			fmt.Printf("deepSearch: %s deep:%d maxDeep Score %d\n", players[player], deep, score)
		}
		return nil, score
	}

	list, err := generateMoveList(b, player)
	if err != nil {
		if DEBUG {
			fmt.Printf("deepSearch: %s deep:%d, err:%s\n", players[player], deep, err)
		}
		return nil, 0
	}

	result := filterKingCaptures(b, player, list)

	if DEBUG {
		fmt.Printf("deepSearch: %s deep:%d, moves:%s\n", players[player], deep, moveList(result))
	}

	for i, m := range result {
		if DEBUG {
			fmt.Printf("deepSearch: %s deep:%d, move[%d/%d]: %s\n", players[player], deep, i+1, len(result), m)
		}
		b.doMove(m)
		_, score := deepSearch(b, otherPlayer(player), deep+1, maxDeep)
		b.undoMove(m)

		if (player == WHITE && score > bestScore) || (player == BLACK && score < bestScore) {
			bestScore = score
			bestMove = m
		}
	}
	if DEBUG {
		fmt.Printf("deepSearch: %s deep:%d, bestMove %s, bestScore %d\n",
			players[player], deep, bestMove, bestScore)
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
