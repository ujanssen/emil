package emil_test

import (
	"github.com/ujanssen/emil"
	"testing"
)

var db *emil.EndGameDb

func init() {
	db = emil.NewEndGameDb()
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

/*
  a b c d e f g h
8               k 8
7           K     7
6         R       6
5                 5
4                 4
3                 3
2                 2
1                 1
  a b c d e f g h
*/
func TestFindMoveRe6h6(t *testing.T) {
	board := emil.NewBoard()
	board.Setup(emil.WhiteKing, emil.F7)
	board.Setup(emil.BlackKing, emil.H8)
	board.Setup(emil.WhiteRock, emil.E6)

	want := "Re6h6"
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
