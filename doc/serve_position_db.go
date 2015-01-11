package main

import (
	"github.com/ujanssen/emil"
	"log"
	"net"
	"net/rpc"
	"time"
)

func main() {
	emil.DEBUG = true

	log.Printf("load Possitions\n")
	start := time.Now()
	db, _ := emil.LoadPositionDb("Positions.gob")
	end := time.Now()
	log.Printf("load duration %v\n", end.Sub(start))
	log.Printf("db.Positions %v\n", len(db.Positions))

	rpc.Register(db)
	listener, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	for {
		if conn, err := listener.Accept(); err != nil {
			log.Fatal("accept error: " + err.Error())
		} else {
			log.Printf("new connection established\n")
			go rpc.ServeConn(conn)
		}
	}
}
