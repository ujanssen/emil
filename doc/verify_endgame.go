package main

import (
	"fmt"
	"github.com/ujanssen/emil"
	"time"
)

func main() {
	emil.DEBUG = false

	start := time.Now()
	db := emil.NewEndGameDb()
	end := time.Now()

	for dtm := 0; dtm < db.MaxDtm(); dtm++ {
		fmt.Println("db.FindMatesIn", dtm)
		boards := db.FindMatesIn(dtm)

		for i, b := range boards {
			fmt.Printf("%d\n%s\n\n", i+1, b)
		}
		fmt.Printf("\n\n\n")
	}

	for dtm := 0; dtm < db.MaxDtm(); dtm++ {
		boards := db.FindMatesIn(dtm)
		fmt.Printf("db.FindMatesIn %d: %d boards\n", dtm, len(boards))
	}

	fmt.Printf("\n\n\nduration %v\n\n\n", end.Sub(start))
}
