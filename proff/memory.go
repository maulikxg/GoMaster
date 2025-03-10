package main

import (
	"log"
	"os"
	"runtime/pprof"
	"time"
)

func allocateMemory() {
	// Simulate memory allocation
	slice := make([]int, 1e7) // 10 million integers
	for i := range slice {
		slice[i] = i // Fill the slice
	}
	time.Sleep(1 * time.Second) // Keep it alive briefly
}

func main() {
	// Open file to save memory profile
	f, err := os.Create("mem.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Run some work to profile
	for i := 0; i < 5; i++ {
		allocateMemory()
	}

	// Write heap profile to file
	if err := pprof.WriteHeapProfile(f); err != nil {
		log.Fatal(err)
	}
}
