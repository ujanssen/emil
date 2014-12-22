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

// Piece Values
const (
	kingValue = 10000
	rockValue = 500

	WhiteKing = kingValue
	BlackKing = BLACK * kingValue
	WhiteRock = rockValue
	BlackRock = BLACK * rockValue
)

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
	//FirstSquares of rank a to print the board
	FirstSquares = [...]int{A8, A7, A6, A5, A4, A3, A2, A1}

	//BoardSquares an array of *Square of the board
	BoardSquares [SQUARES]*Square

	kingDirections = [...]int{North, South, West, East, NorthWest, NorthEast, SouthWest, SouthEast}
	rookDirections = [...]int{North, South, West, East}

	kingMoves [SQUARES][]int
	rockMoves [SQUARES][][]int

	squaresDistances [SQUARES][SQUARES]int
)
