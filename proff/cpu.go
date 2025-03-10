package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

func main() {

	f, err := os.Create("cpu.prof")
	if err != nil {
		fmt.Printf("Failed to create CPU profile: %s\n", err)
	}
	defer f.Close()

	// starting the cpu profiling
	if err := pprof.StartCPUProfile(f); err != nil {
		fmt.Printf("Failed to start CPU profile: %s\n", err)
	}
	defer pprof.StopCPUProfile()

	for i := 0; i < 5; i++ {
		heavyComputation()
	}

	time.Sleep(2 * time.Second)

}

func heavyComputation() {

	for i := 0; i < 1e8; i++ {
		_ = i * i
	}
}
