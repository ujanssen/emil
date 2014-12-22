package emil

import (
	"fmt"
)

//Search prints all moves for a piece on square
func Search(b *Board, player int, testKingCapture bool) (string, bool) {
	empty := ""
	var result, list []*Move
	for src, piece := range b.squares {
		if isOwnPiece(player, piece) {
			switch abs(piece) {
			case kingValue:
				for _, dst := range kingDestinationsFrom(src) {
					capture := b.squares[dst]
					if isKing(capture) {
						return empty, true
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
							return empty, true
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

	if testKingCapture {
		return empty, false
	}

	for _, m := range list {
		println("TEST move", m.String())
		b.doMove(m)
		if !b.isKingCapturedAfter(m) {
			result = append(result, m)
		} else {
			println("KingCaptured")
		}
		b.undoMove(m)
		fmt.Printf("\n\n\n")
	}

	return moveList(result), false
}
