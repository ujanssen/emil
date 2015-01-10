package main

import (
	"fmt"
	"github.com/ujanssen/emil"
	"time"
)

func main() {
	emil.DEBUG = true

	start := time.Now()
	db, _ := emil.LoadPositionDb("Positions.json")
	end := time.Now()
	fmt.Printf("load duration %v\n", end.Sub(start))
	fmt.Printf("db.Positions %v\n", len(db.Positions))
}
