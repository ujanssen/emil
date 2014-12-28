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
func isKingInCheck(p *position) (kingInCheck bool) {
	_, kingCaptured := generateMoveList(p.board, otherPlayer(p.player))
	if kingCaptured != nil {
		return true
	}
	return false

}
func filterKingCaptures(p *position, list []*Move) (result []*Move) {
	for _, m := range list {
		newBoard := p.board.doMove(m)
		newPosition := NewPosition(newBoard, p.player)
		if !isKingInCheck(newPosition) {
			result = append(result, m)
		}
	}
	return result
}

//Search best move for player on board
func Search(p *position) (bestMove *Move) {
	bestMove, _ = deepSearch(p, 0, 1)
	return bestMove
}

//Search best move for player on board
func deepSearch(p *position, deep, maxDeep int) (bestMove *Move, bestScore int) {
	if p.player == WHITE {
		bestScore = 2 * BlackKing
	} else {
		bestScore = 2 * WhiteKing
	}

	if deep == maxDeep {
		score := evaluate(p.board)
		if DEBUG {
			fmt.Printf("deepSearch: %s deep:%d maxDeep Score %d\n", players[p.player], deep, score)
		}
		return nil, score
	}

	list, err := generateMoveList(p.board, p.player)
	if err != nil {
		if DEBUG {
			fmt.Printf("deepSearch: %s deep:%d, err:%s\n", players[p.player], deep, err)
		}
		return nil, 0
	}
	result := filterKingCaptures(p, list)

	if DEBUG {
		fmt.Printf("deepSearch: %s deep:%d, moves:%s\n", players[p.player], deep, moveList(result))
	}

	for i, m := range result {
		if DEBUG {
			fmt.Printf("deepSearch: %s deep:%d, move[%d/%d]: %s\n", players[p.player], deep, i+1, len(result), m)
		}
		newBoard := p.board.doMove(m)
		newPosition := NewPosition(newBoard, otherPlayer(p.player))

		_, score := deepSearch(newPosition, deep+1, maxDeep)

		if (p.player == WHITE && score > bestScore) || (p.player == BLACK && score < bestScore) {
			bestScore = score
			bestMove = m
		}
	}
	if DEBUG {
		fmt.Printf("deepSearch: %s deep:%d, bestMove %s, bestScore %d\n",
			players[p.player], deep, bestMove, bestScore)
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
