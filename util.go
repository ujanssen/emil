package emil

func isOtherKing(player, capture int) bool {
	return (player == WHITE && capture == BlackKing) ||
		(player == BLACK && capture == WhiteKing)
}

func isOwnPiece(player, capture int) bool {
	return (player == WHITE && capture > 0) ||
		(player == BLACK && capture < 0)
}

func otherPlayer(player int) int {
	if player == WHITE {
		return BLACK
	}
	return WHITE
}

func validIndex(i int) bool {
	if i >= A1 && i <= H8 {
		return true
	}
	return false
}

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
