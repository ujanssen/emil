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
	//	err := emil.SaveEndGameDb("AnalysisMap.json", db.AnalysisStr)
	end := time.Now()
	db.Positions()
	fmt.Printf("db.Positions() %d\n\n", db.Positions())
	fmt.Printf("\n\n\nsave duration %v\n\n", end.Sub(start))
}
