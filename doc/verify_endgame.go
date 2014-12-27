package main

import (
	"fmt"
	"github.com/ujanssen/emil"
)

func main() {
	emil.DEBUG = true

	db := emil.NewEndGameDb()

	fmt.Println("db.FindMate(emil.BlackKing, emil.A8")

	boards := db.FindMate(emil.BlackKing, emil.A8)
	fmt.Printf("found %d boards\n\n", len(boards))

	for _, b := range boards {
		fmt.Printf("%s\n\n", b)
	}
}
