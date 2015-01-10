package emil

import (
	"errors"
	"fmt"
)

var errKingCapture = errors.New("king capture")

func generateMoveListWe(p *position) (list []*Move, err error) {
	var empty []*Move
	for src, piece := range p.Board.Squares {
		if isOwnPiece(p.Player, piece) {
			switch abs(piece) {
			case kingValue:
				for _, dst := range kingDestinationsFrom(src) {
					capture := p.Board.Squares[dst]
					if isOtherKing(p.Player, capture) {
						return empty, errKingCapture
					}
					if capture == Empty {
						list = append(list, newSilentMove(p.Player, piece, src, dst))
					} else if !isOwnPiece(p.Player, capture) {
						list = append(list, newCaptureMove(p.Player, piece, capture, src, dst))
					}
				}
			case rockValue:
				for _, dsts := range rockDestinationsFrom(src) {
					for _, dst := range dsts {
						capture := p.Board.Squares[dst]
						if isOtherKing(p.Player, capture) {
							return empty, errKingCapture
						}
						if capture == Empty {
							list = append(list, newSilentMove(p.Player, piece, src, dst))
						} else if !isOwnPiece(p.Player, capture) {
							list = append(list, newCaptureMove(p.Player, piece, capture, src, dst))
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
func CanPlayerCaptureKing(p *position) (kingInCheck bool) {
	_, kingCaptured := generateMoveListWe(p)
	if kingCaptured != nil {
		return true
	}
	return false
}

func isKingInCheck(p *position) (kingInCheck bool) {
	newPosition := NewPosition(p.Board, otherPlayer(p.Player))
	_, kingCaptured := generateMoveListWe(newPosition)
	if kingCaptured != nil {
		return true
	}
	return false

}
func filterKingCaptures(p *position, list []*Move) (result []*Move) {
	for _, m := range list {
		newBoard := p.Board.DoMove(m)
		newPosition := NewPosition(newBoard, p.Player)
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
	if p.Player == WHITE {
		bestScore = 2 * BlackKing
	} else {
		bestScore = 2 * WhiteKing
	}

	if deep == maxDeep {
		score := evaluate(p.Board)
		if DEBUG {
			fmt.Printf("deepSearch: %s deep:%d maxDeep Score %d\n", players[p.Player], deep, score)
		}
		return nil, score
	}

	list, err := generateMoveListWe(p)
	if err != nil {
		if DEBUG {
			fmt.Printf("deepSearch: %s deep:%d, err:%s\n", players[p.Player], deep, err)
		}
		return nil, 0
	}
	result := filterKingCaptures(p, list)

	if DEBUG {
		fmt.Printf("deepSearch: %s deep:%d, moves:%s\n", players[p.Player], deep, moveList(result))
	}

	for i, m := range result {
		if DEBUG {
			fmt.Printf("deepSearch: %s deep:%d, move[%d/%d]: %s\n", players[p.Player], deep, i+1, len(result), m)
		}
		newBoard := p.Board.DoMove(m)
		newPosition := NewPosition(newBoard, otherPlayer(p.Player))

		_, score := deepSearch(newPosition, deep+1, maxDeep)

		if (p.Player == WHITE && score > bestScore) || (p.Player == BLACK && score < bestScore) {
			bestScore = score
			bestMove = m
		}
	}
	if DEBUG {
		fmt.Printf("deepSearch: %s deep:%d, bestMove %s, bestScore %d\n",
			players[p.Player], deep, bestMove, bestScore)
	}
	return bestMove, bestScore
}

func evaluate(b *Board) (score int) {
	score = 0
	for _, p := range b.Squares {
		score += p
	}
	return score
}
