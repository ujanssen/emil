package main

import (
	"flag"
	"fmt"
	"github.com/ujanssen/emil"
)

func main() {
	fen := flag.String("fen", "8/8/8/8/8/8/8/8", "Forsyth–Edwards Notation")
	flag.Parse()
	fmt.Println(emil.Fen2Board(*fen).Picture())
}
