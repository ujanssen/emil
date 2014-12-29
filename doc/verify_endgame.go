package main

import (
	"fmt"
	"github.com/ujanssen/emil"
)

func main() {
	emil.DEBUG = true

	db := emil.NewEndGameDb()

	for dtm := 0; dtm <= 4; dtm++ {
		fmt.Println("db.FindMatesIn", dtm)
		boards := db.FindMatesIn(dtm)

		for i, b := range boards {
			fmt.Printf("%d\n%s\n\n", i+1, b)
		}
		fmt.Printf("\n\n\n")
	}

	for dtm := 0; dtm <= 4; dtm++ {
		boards := db.FindMatesIn(dtm)
		fmt.Printf("db.FindMatesIn %d: %d boards\n\n", dtm, len(boards))
	}
}
