package main

import (
	"flag"
	"fmt"
	"github.com/ujanssen/emil"
	"log"
	"net/rpc"
)

func main() {
	emil.DEBUG = true
	fen := flag.String("fen", "8/8/8/8/8/8/8/8", "Forsyth–Edwards Notation")
	flag.Parse()
	fmt.Println(emil.Fen2Board(*fen).Picture())

	client, err := rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	// Synchronous call
	var pe emil.PositionEntry
	err = client.Call("PositionDb.FindWhitePosition", fen, &pe)
	if err != nil {
		log.Fatal("db error:", err)
	} else {
		fmt.Printf("Position: %v\n", pe.Position)
		fmt.Printf("DTM: %v\n", pe.Dtm)
		fmt.Printf("NextPositions: %v\n", pe.NextPositions)
		fmt.Printf("PrevPositions: %v\n", pe.PrevPositions)
	}
}
