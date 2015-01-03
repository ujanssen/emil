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
	kingValue   = 10000
	queenValue  = 900
	rockValue   = 500
	bishopValue = 300
	knightValue = 301
	pawnValue   = 100

	Empty = 0

	WhiteKing = kingValue
	BlackKing = BLACK * kingValue

	WhiteQueen = queenValue
	BlackQueen = BLACK * queenValue

	WhiteRock = rockValue
	BlackRock = BLACK * rockValue

	WhiteBishop = bishopValue
	BlackBishop = BLACK * bishopValue

	WhiteKnight = knightValue
	BlackKnight = BLACK * knightValue

	WhitePawn = pawnValue
	BlackPawn = BLACK * pawnValue
)

// Pieces symbols
var (
	Pieces = map[int]string{
		WhiteKing:   "K",
		BlackKing:   "k",
		WhiteQueen:  "Q",
		BlackQueen:  "q",
		WhiteRock:   "R",
		BlackRock:   "r",
		WhiteBishop: "B",
		BlackBishop: "b",
		WhiteKnight: "N",
		BlackKnight: "n",
		WhitePawn:   "P",
		BlackPawn:   "p",
		Empty:       " ",
	}
	Symbols = map[string]int{
		"K": WhiteKing,
		"k": BlackKing,
		"Q": WhiteQueen,
		"q": BlackQueen,
		"R": WhiteRock,
		"r": BlackRock,
		"B": WhiteBishop,
		"b": BlackBishop,
		"N": WhiteKnight,
		"n": BlackKnight,
		"P": WhitePawn,
		"p": BlackPawn,
		" ": Empty,
	}
)

var (
	//FirstSquares of rank a to print the board
	FirstSquares = [...]int{A8, A7, A6, A5, A4, A3, A2, A1}

	//BoardSquares an array of *Square of the board
	BoardSquares [SQUARES]*Square

	squareMap map[string]int

	squaresDistances [SQUARES][SQUARES]int
)
