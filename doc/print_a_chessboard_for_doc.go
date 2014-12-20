package main

import (
	"fmt"
)

const (
	a1 = 0
	b1 = 1
	c1 = 2
	d1 = 3
	e1 = 4
	f1 = 5
	g1 = 6
	h1 = 7
	a2 = 8
	b2 = 9
	c2 = 10
	d2 = 11
	e2 = 12
	f2 = 13
	g2 = 14
	h2 = 15
	a3 = 16
	b3 = 17
	c3 = 18
	d3 = 19
	e3 = 20
	f3 = 21
	g3 = 22
	h3 = 23
	a4 = 24
	b4 = 25
	c4 = 26
	d4 = 27
	e4 = 28
	f4 = 29
	g4 = 30
	h4 = 31
	a5 = 32
	b5 = 33
	c5 = 34
	d5 = 35
	e5 = 36
	f5 = 37
	g5 = 38
	h5 = 39
	a6 = 40
	b6 = 41
	c6 = 42
	d6 = 43
	e6 = 44
	f6 = 45
	g6 = 46
	h6 = 47
	a7 = 48
	b7 = 49
	c7 = 50
	d7 = 51
	e7 = 52
	f7 = 53
	g7 = 54
	h7 = 55
	a8 = 56
	b8 = 57
	c8 = 58
	d8 = 59
	e8 = 60
	f8 = 61
	g8 = 62
	h8 = 63
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
	firstSquares := [...]int{a8, a7, a6, a5, a4, a3, a2, a1}
	s := "  " + files + " \n"
	rank := 8
	for _, r := range firstSquares {
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
