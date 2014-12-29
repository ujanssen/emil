package main

import (
	"fmt"
	"github.com/ujanssen/emil"
)

func main() {
	emil.DEBUG = true

	db := emil.NewEndGameDb()

	for dtm := 0; dtm <= 2; dtm++ {
		fmt.Println("db.FindMatesIn", dtm)
		boards := db.FindMatesIn(dtm)
		fmt.Printf("found %d boards\n\n", len(boards))

		for i, b := range boards {
			fmt.Printf("%d\n%s\n\n", i+1, b)
		}
		fmt.Printf("\n\n\n")
	}
}
