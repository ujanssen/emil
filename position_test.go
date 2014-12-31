package emil_test

import (
	"github.com/ujanssen/emil"
	"testing"
)

/*
  a b c d e f g h
8                 8
7                 7
6                 6
5                 5
4                 4
3                 3
2                 2
1                 1
  a b c d e f g h
*/
func TestEmptyBoard(t *testing.T) {
	board := emil.NewBoard()

	want := "8/8/8/8/8/8/8/8 w"
	got := emil.NewPosition(board, emil.WHITE).String()
	if got != want {
		t.Errorf("the move should be %s, got %s", want, got)
	}
}

/*
  a b c d e f g h
8         k       8
7                 7
6         K       6
5                 5
4                 4
3                 3
2                 2
1               R 1
  a b c d e f g h
*/
func TestFindMoveRh1h8(t *testing.T) {
	board := emil.NewBoard()
	board.Setup(emil.WhiteKing, emil.E6)
	board.Setup(emil.BlackKing, emil.E8)
	board.Setup(emil.WhiteRock, emil.H1)

	want := "4k3/8/4K3/8/8/8/8/7R w"
	got := emil.NewPosition(board, emil.WHITE).String()
	if got != want {
		t.Errorf("the move should be %s, got %s", want, got)
	}
}
