package emil_test

import (
	"github.com/ujanssen/emil"
	"testing"
)

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
func TestFindCheckmateMove(t *testing.T) {
	board := emil.NewBoard()
	board.Setup(emil.WhiteKing, emil.E6)
	board.Setup(emil.BlackKing, emil.E8)
	board.Setup(emil.WhiteRock, emil.H1)

	db := emil.NewEndGameDb()

	want := "Rh1h8"
	move := db.Find(board)
	if move == nil {
		t.Errorf("the move should be %s, got nil", want)
	} else {
		got := move.String()
		if got != want {
			t.Errorf("the move should be %s, got %s", want, got)
		}
	}
}
