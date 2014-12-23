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
	var bestScore int

	if player == WHITE {
		bestScore = 2 * BlackKing
	} else {
		bestScore = 2 * WhiteKing
	}

	list, err := generateMoveList(b, player)
	if err != nil {
		return nil
	}

	result := filterKingCaptures(b, player, list)
	for _, m := range result {
		score := evaluate(b, player, m)
		if (player == WHITE && score > bestScore) || (player == BLACK && score < bestScore) {
			bestScore = score
			bestMove = m
		}
	}
	return bestMove
}

func evaluate(b *Board, player int, m *Move) (score int) {
	score = 0
	b.doMove(m)
	for _, p := range b.squares {
		score += p
	}
	if DEBUG {
		println("Score:", score)
	}
	b.undoMove(m)
	return score
}
