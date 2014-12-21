package emil

import (
	"fmt"
)

//Square of a chess board
type Square struct {
	name  string
	file  string
	rank  int
	index int
}

func newSquare(file string, rank, index int) *Square {
	name := fmt.Sprintf("%v%d", file, rank)

	return &Square{
		name:  name,
		file:  file,
		rank:  rank,
		index: index}
}

func (s *Square) String() string {
	return s.name
}
func (s *Square) distance(o *Square) int {
	return max(
		abs(s.rank-o.rank),
		abs(int(s.file[0])-int(o.file[0])))
}

func (s *Square) isSameRankOrFile(o *Square) bool {
	return s.rank == o.rank || s.file == o.file
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
