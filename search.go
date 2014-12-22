package emil

import (
	"errors"
	"fmt"
)

var errKingCapture = errors.New("king capture")

//Search prints all moves for a piece on square
func Search(b *Board, player int, onlyTestKingCapture bool) (empty string, err error) {
	empty = ""
	var result, list []*Move
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

	if onlyTestKingCapture {
		return empty, nil
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

	return moveList(result), nil
}
