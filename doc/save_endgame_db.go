package main

import (
	"fmt"
	"github.com/ujanssen/emil"
	"time"
)

func main() {
	emil.IN_TEST = !true
	emil.DEBUG = true

	start := time.Now()
	db := emil.NewEndGameDb()
	end := time.Now()

	for dtm := -1; dtm < 30; dtm++ {
		as := db.FindMatesIn(dtm)
		if len(as) > 0 {
			fmt.Printf("db.FindMatesIn %d: %d boards\n", dtm, len(as))
		}
	}

	fmt.Printf("\n\n\ncreate duration %v\n\n\n", end.Sub(start))

	start = time.Now()
	err := emil.SaveEndGameDb(db)
	end = time.Now()
	fmt.Printf("\n\n\nsave duration %v\nerr %v\n\n", end.Sub(start), err)
}
