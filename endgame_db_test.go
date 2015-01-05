package emil_test

import (
	"fmt"
	"github.com/ujanssen/emil"
	"testing"
)

var db *emil.EndGameDb

/*
  a b c d e f g h
8   k             8
7                 7
6   K             6
5                 5
4                 4
3                 3
2                 2
1         R       1
  a b c d e f g h
FEN: 1k6/8/1K6/8/8/8/8/4R3
*/
func TestFindMoveRh1h8(t *testing.T) {
	board := emil.Fen2Board("1k6/8/1K6/8/8/8/8/4R3")
	p := emil.NewPosition(board, emil.WHITE)
	list := emil.GenerateMoves(p)

	got := fmt.Sprintf("%v", list)
	fmt.Printf("moves %v\n", list)

	want := "[Re1e2 Re1e3 Re1e4 Re1e5 Re1e6 Re1e7 Re1e8 Re1d1 Re1c1 Re1b1 Re1a1 Re1f1 Re1g1 Re1h1 Kb6b5 Kb6a6 Kb6c6 Kb6a5 Kb6c5]"
	if got != want {
		t.Errorf("the moves should be %s, got %s", want, got)
	}

}
