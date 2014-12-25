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

	isBorder bool
}

func newSquare(file string, rank, index int) *Square {
	name := fmt.Sprintf("%v%d", file, rank)

	return &Square{
		name:     name,
		file:     file,
		rank:     rank,
		index:    index,
		isBorder: isBorderSquare(file, rank)}
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

func isBorderSquare(file string, rank int) bool {
	if rank == 1 || rank == 8 {
		return true
	}
	if file == "a" || file == "h" {
		return true
	}
	return false
}
