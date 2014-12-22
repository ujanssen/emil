package emil

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
