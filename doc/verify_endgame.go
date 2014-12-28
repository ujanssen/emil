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

	fmt.Println("---------------\n\n")
	fmt.Println("db.FindMatesIn2")

	boards = db.FindMatesIn(2)
	fmt.Printf("found %d boards\n\n", len(boards))

	for i, b := range boards {
		fmt.Printf("%d\n%s\n\n", i+1, b)
	}
}
