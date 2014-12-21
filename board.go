package emil

import (
	"fmt"
)

// Board with an array of field values, representing pieces
type Board struct {
	squares []int
}

// NewBoard creates a new Board
func NewBoard() *Board {
	return &Board{squares: make([]int, SQUARES)}
}
func (b *Board) String() string {
	files := "a b c d e f g h "
	s := "  " + files + " \n"
	for _, r := range FirstSquares {
		s += fmt.Sprintf("%d ", BoardSquares[r].rank)
		for f := 0; f < 8; f++ {
			s += fmt.Sprintf("%s ", symbol(b.squares[r+f]))
		}
		s += fmt.Sprintf("%d\n", BoardSquares[r].rank)
	}
	s += "  " + files + " \n"
	return s
}

func symbol(piece int) string {
	switch piece {
	case WhiteKing:
		return "K"
	case BlackKing:
		return "k"
	case WhiteRock:
		return "R"
	case BlackRock:
		return "r"
	default:
		return " "
	}
}

//Setup a piece on a square
func (b *Board) Setup(piece, square int) {
	b.squares[square] = piece
}

//Moves prints all moves for a piece on square
func (b *Board) Moves(piece, square int) string {
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

//Square of a chess board
type Square struct {
	name  string
	file  string
	rank  int
	index int
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

func newSquare(file string, rank, index int) *Square {
	name := fmt.Sprintf("%v%d", file, rank)

	return &Square{
		name:  name,
		file:  file,
		rank:  rank,
		index: index}
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
