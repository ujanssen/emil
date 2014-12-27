package main

import (
	"fmt"
	"github.com/ujanssen/emil"
)

func main() {
	emil.DEBUG = true

	db := emil.NewEndGameDb()

	fmt.Println("db.FindMate(emil.BlackKing, emil.E8")

	for _, b := range db.FindMate(emil.BlackKing, emil.E8) {
		fmt.Printf("%s\n\n", b)
	}
}
