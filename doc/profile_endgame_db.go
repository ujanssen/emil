package main

import (
	"github.com/ujanssen/emil"
	"runtime/pprof"
)

func main() {
	emil.PROFILE = true
	emil.NewEndGameDb()

	defer pprof.StopCPUProfile()
}
