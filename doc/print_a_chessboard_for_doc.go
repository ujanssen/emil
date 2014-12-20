package main

import (
	"fmt"
	"github.com/ujanssen/emil"
)

func printSqares() string {
	s := ""
	for _, index := range emil.FirstSquares {
		for file := 0; file < 8; file++ {
			s += fmt.Sprintf("%v ", emil.BoardSquares[index+file])
		}
		s += fmt.Sprintf("\n")
	}
	return s
}

func printBoard() string {
	files := "a  b  c  d  e  f  g  h "
	s := "   " + files + " \n"
	rank := 8
	for _, r := range emil.FirstSquares {
		s += fmt.Sprintf("%d ", rank)
		for f := 0; f < 8; f++ {
			s += fmt.Sprintf("%2d ", r+f)
		}
		s += fmt.Sprintf("%d\n", rank)
		rank--
	}
	s += "   " + files + " \n"
	return s
}
func main() {
	fmt.Printf("A chess board\n")
	fmt.Printf("\n")
	fmt.Printf("%s\n", printSqares())
	fmt.Printf("\n")
	fmt.Printf("%s\n", printBoard())
}
