package emil

// Directions in terms of board index
const (
	North     = 8
	South     = -8
	West      = -1
	East      = 1
	NorthWest = 7
	NorthEast = 9
	SouthWest = -9
	SouthEast = -7
)

var (
	kingDirections = [...]int{North, South, West, East, NorthWest, NorthEast, SouthWest, SouthEast}
	rookDirections = [...]int{North, South, West, East}

	kingMoves [SQUARES][]int
	rockMoves [SQUARES][][]int
)

func destinations(piece, square int) string {
	switch piece {
	case WhiteKing:
		return squareList(kingDestinationsFrom(square))
	case BlackKing:
		return squareList(kingDestinationsFrom(square))
	case WhiteRock:
		return squareLists(rockDestinationsFrom(square))
	case BlackRock:
		return squareLists(rockDestinationsFrom(square))
	default:
		panic("yet not implemented")
	}
}

func kingDestinationsFrom(source int) []int {
	var list []int
	for _, d := range kingDirections {
		dst := source + d
		if validIndex(dst) && squaresDistances[source][dst] == 1 {
			list = append(list, dst)
		}
	}
	return list
}
func rockDestinationsFrom(source int) [][]int {
	var list [][]int
	for _, d := range rookDirections {
		var dstList []int
		for step := 1; step < 8; step++ {
			dst := source + (step * d)
			if validIndex(dst) && squaresDistances[source][dst] == step && BoardSquares[source].isSameRankOrFile(BoardSquares[dst]) {
				dstList = append(dstList, dst)
			} else {
				break
			}
		}
		if len(dstList) > 0 {
			list = append(list, dstList)
		}
	}
	return list
}

func squareList(list []int) string {
	r := "["
	for i, s := range list {
		if i > 0 {
			r += ", "
		}
		r += BoardSquares[s].name
	}
	r += "]"
	return r
}

func squareLists(lists [][]int) string {
	r := "["
	for i, list := range lists {
		if i > 0 {
			r += ", "
		}
		r += squareList(list)
	}
	r += "]"
	return r
}
