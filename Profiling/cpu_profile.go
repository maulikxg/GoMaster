package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

func main() {

	f, err := os.Create("cpu.pprof")
	if err != nil {
		fmt.Println("Error creating CPU profile: ", err)
	}

	defer f.Close()

	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	slower()

	time.Sleep(time.Second)

	for i := 0; i < 10000000000; i++ {
		_ = i + 1
	}

}

func slower() {

	for i := 0; i < 10000000000; i++ {
		_ = i * i
	}
}
