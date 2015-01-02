package emil

// Players
const (
	WHITE = 1
	BLACK = -1
)

var (
	players = map[int]string{
		WHITE: "white",
		BLACK: "black",
	}
)

// DEBUG set to write cpu.po
var DEBUG = false

// PROFILE set to true for log messages
var PROFILE = false

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

	Empty     = 0
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
	Symbols = map[string]int{
		"K": WhiteKing,
		"k": BlackKing,
		"R": WhiteRock,
		"r": BlackRock,
	}
)

var (
	//FirstSquares of rank a to print the board
	FirstSquares = [...]int{A8, A7, A6, A5, A4, A3, A2, A1}

	//BoardSquares an array of *Square of the board
	BoardSquares [SQUARES]*Square

        squareMap map[string]*Square

	squaresDistances [SQUARES][SQUARES]int
)
