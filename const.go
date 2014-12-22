package emil

// Players
const (
	WHITE = 1
	BLACK = -1
)

const (
	//FILES the files of a board
	FILES = "abcdefgh"

	//SQUARES the number of quares
	SQUARES = 64
)

// Pieces values
const (
	kingValue = 10000
	rockValue = 500

	WhiteKing = kingValue
	BlackKing = BLACK * kingValue
	WhiteRock = rockValue
	BlackRock = BLACK * rockValue
)

// Pieces symbols
var (
	Pieces = map[int]string{
		WhiteKing: "K",
		BlackKing: "k",
		WhiteRock: "R",
		BlackRock: "r",
	}
)

var (
	//FirstSquares of rank a to print the board
	FirstSquares = [...]int{A8, A7, A6, A5, A4, A3, A2, A1}

	//BoardSquares an array of *Square of the board
	BoardSquares [SQUARES]*Square

	squaresDistances [SQUARES][SQUARES]int
)
