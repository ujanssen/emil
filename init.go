package emil

func init() {
	// define BoardSquares
	for j := 0; j < 8; j++ {
		for i, file := range FILES {
			index := i + j*8
			BoardSquares[index] = newSquare(string(file), j+1, index)
		}
	}

	// define squaresDistances
	for _, s := range BoardSquares {
		for _, r := range FirstSquares {
			for f := 0; f < 8; f++ {
				squaresDistances[s.index][r+f] = s.distance(BoardSquares[r+f])
			}
		}
	}

	// compute piece moves
	for i := A1; i <= H8; i++ {
		kingMoves[i] = kingDestinationsFrom(i)
		rockMoves[i] = rockDestinationsFrom(i)
	}
}
