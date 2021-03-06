package emil_test

import (
	"github.com/ujanssen/emil"
	"testing"
)

/*
  a b c d e f g h
8   R         k   8
7                 7
6           K     6
5                 5
4                 4
3                 3
2                 2
1                 1
  a b c d e f g h
*/
func TestForTheOnlyBlackMove(t *testing.T) {
	board := emil.NewBoard()
	board.Setup(emil.WhiteKing, emil.F6)
	board.Setup(emil.BlackKing, emil.G8)
	board.Setup(emil.WhiteRock, emil.B8)
	p := emil.NewPosition(board, emil.BLACK)

	want := "kg8h7"
	got := emil.Search(p).String()
	if got != want {
		t.Errorf("the moves should be %s, got %s", want, got)
	}
}

/*
  a b c d e f g h
8         k     R 8
7                 7
6         K       6
5                 5
4                 4
3                 3
2                 2
1                 1
  a b c d e f g h
*/
func TestNoBlackMove(t *testing.T) {
	board := emil.NewBoard()
	board.Setup(emil.WhiteKing, emil.E6)
	board.Setup(emil.BlackKing, emil.E8)
	board.Setup(emil.WhiteRock, emil.H8)
	p := emil.NewPosition(board, emil.BLACK)

	var want *emil.Move
	got := emil.Search(p)
	if got != want {
		t.Errorf("the moves should be %s, got %s", want, got)
	}
}

/*
  a b c d e f g h
8             k   8
7                 7
6         K       6
5                 5
4                 4
3                 3
2                 2
1 r R             1
  a b c d e f g h
*/
func TestRockCapture(t *testing.T) {
	board := emil.NewBoard()
	board.Setup(emil.WhiteKing, emil.F6)
	board.Setup(emil.BlackKing, emil.G8)
	board.Setup(emil.WhiteRock, emil.B1)
	board.Setup(emil.BlackRock, emil.A1)
	p := emil.NewPosition(board, emil.WHITE)

	want := "Rb1xa1"
	got := emil.Search(p).String()
	if got != want {
		t.Errorf("the moves should be %s, got %s", want, got)
	}
}
