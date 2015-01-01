package main

import (
	"fmt"
	"github.com/ujanssen/emil"
	"time"
)

func main() {
	emil.IN_TEST = true
	emil.DEBUG = true

	start := time.Now()
	db := emil.NewEndGameDb()
	end := time.Now()

	for dtm := -1; dtm < db.MaxDtm(); dtm++ {
		as := db.FindMatesIn(dtm)
		fmt.Printf("db.FindMatesIn %d: %d boards\n", dtm, len(as))
	}

	fmt.Printf("\n\n\ncreate duration %v\n\n\n", end.Sub(start))

	start = time.Now()
	emil.SaveEndGameDb(db)
	end = time.Now()
	fmt.Printf("\n\n\nsave duration %v\n\n\n", end.Sub(start))
}
