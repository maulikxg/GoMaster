package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"
	"time"
)

var mu sync.Mutex

func slowFunction() {
	mu.Lock()
	defer mu.Unlock()
	time.Sleep(3 * time.Second) // Simulate delay
}

func main() {
	// Enable block profiling
	runtime.SetBlockProfileRate(1)

	// Create profile file
	f, err := os.Create("block.prof")
	if err != nil {
		fmt.Println("Could not create block profile:", err)
		return
	}
	defer f.Close()

	// Start multiple goroutines
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		slowFunction()
	}()
	go func() {
		defer wg.Done()
		slowFunction()
	}()
	go func() {
		defer wg.Done()
		slowFunction()
	}()

	wg.Wait()

	// Write block profile data to file
	pprof.Lookup("block").WriteTo(f, 0)
}
