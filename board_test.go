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
	want := "8/8/8/8/8/8/8/8"
	got := emil.Fen2Board(want).String()
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
func TestPrintFEN(t *testing.T) {
	want := "4k3/8/4K3/8/8/8/8/7R"
	got := emil.Fen2Board(want).String()
	if got != want {
		t.Errorf("the move should be %s, got %s", want, got)
	}
}
