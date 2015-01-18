package main

import (
	"fmt"
	"github.com/ujanssen/emil"
	"time"
)

func main() {
	emil.IN_TEST = !true
	emil.DEBUG = true

	db := emil.NewPositionDB()
	start := time.Now()
	db.FillWithKRKPositions()
	end := time.Now()
	fmt.Printf("db.Positions() %d\n\n", len(db.Positions))
	fmt.Printf("FillWithKRKPositions %v\n\n", end.Sub(start))

	start = time.Now()
	db.AddPrevPositions()
	end = time.Now()
	fmt.Printf("db.Positions() %d\n\n", len(db.Positions))
	fmt.Printf("AddPrevPositions %v\n\n", end.Sub(start))

	db.SavePositionDb("Positions.gob")

}
