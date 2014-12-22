package emil_test

import (
	"github.com/ujanssen/emil"
	"testing"
)

func TestForTheOnlyBlackMove(t *testing.T) {
	board := emil.NewBoard()
	board.Setup(emil.WhiteKing, emil.F6)
	board.Setup(emil.BlackKing, emil.G8)
	board.Setup(emil.WhiteRock, emil.B8)

	want := "[kg8h7]"
	got, _ := emil.Search(board, emil.BLACK, false)
	if got != want {
		t.Errorf("the moves should be %s, got %s", want, got)
	}
}
