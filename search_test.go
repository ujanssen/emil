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

	want := "[kg8h7]"
	got, _ := emil.Search(board, emil.BLACK)
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

	want := "[]"
	got, _ := emil.Search(board, emil.BLACK)
	if got != want {
		t.Errorf("the moves should be %s, got %s", want, got)
	}
}
