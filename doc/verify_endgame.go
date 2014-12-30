package main

import (
	"fmt"
	"github.com/ujanssen/emil"
	"time"
)

func main() {
	emil.DEBUG = true

	start := time.Now()
	db := emil.NewEndGameDb()
	end := time.Now()

	for dtm := -1; dtm < db.MaxDtm(); dtm++ {
		fmt.Println("db.FindMatesIn", dtm)
		as := db.FindMatesIn(dtm)

		for i, a := range as {
			fmt.Printf("%d %s\n%s\n\n", i+1, a.Move(), a.Board())
		}
		fmt.Printf("\n\n\n")
	}

	for dtm := -1; dtm < db.MaxDtm(); dtm++ {
		as := db.FindMatesIn(dtm)
		fmt.Printf("db.FindMatesIn %d: %d boards\n", dtm, len(as))
	}

	fmt.Printf("\n\n\nduration %v\n\n\n", end.Sub(start))
}
