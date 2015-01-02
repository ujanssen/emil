package emil_test

import (
	"github.com/ujanssen/emil"
	"testing"
)

func TestMoveFromString(t *testing.T) {
	want := "Ke1e2"
	move := emil.MoveFromString(want)
	got := move.String()
	if got != want {
		t.Errorf("the move should be %s, got %s", want, got)
	}
	if p := move.Piece(); p != emil.WhiteKing {
		t.Errorf("the piece should be %d, got %d", emil.WhiteKing, p)
	}
	if p := move.Source(); p != emil.E1 {
		t.Errorf("the piece should be %d, got %d", emil.E1, p)
	}
	if p := move.Destination(); p != emil.E2 {
		t.Errorf("the piece should be %d, got %d", emil.E2, p)
	}

}
