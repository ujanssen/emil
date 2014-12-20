package main

import (
	"fmt"
	"github.com/ujanssen/emil"
)

func printSqares() string {
	files := "abcdefgh"
	s := ""
	for rank := 8; rank > 0; rank-- {
		for file := 0; file < len(files); file++ {
			s += fmt.Sprintf("%s%d ", string(files[file]), rank)
		}
		s += fmt.Sprintf("\n")
	}
	return s
}

func printBoard() string {
	files := "a  b  c  d  e  f  g  h "
	s := "  " + files + " \n"
	rank := 8
	for _, r := range emil.FirstSquares {
		s += fmt.Sprintf("%d ", rank)
		for f := 0; f < 8; f++ {
			s += fmt.Sprintf("%2d ", r+f)
		}
		s += fmt.Sprintf("%d\n", rank)
		rank--
	}
	s += "  " + files + " \n"
	return s
}
func main() {
	fmt.Printf("A chess board\n")
	fmt.Printf("\n")
	fmt.Printf("%s\n", printSqares())
	fmt.Printf("\n")
	fmt.Printf("%s\n", printBoard())
}
