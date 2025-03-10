package main

import (
	"log"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
	"time"
)

func processData(id int, wg *sync.WaitGroup, ch chan<- int, mu *sync.Mutex) {
	defer wg.Done()

	// Simulate CPU-intensive work with memory allocation
	data := make([]int, 0, 10000)
	for i := 0; i < 10000; i++ {
		data = append(data, i*i) // Dynamic growth
		if i%1000 == 0 {
			// Simulate contention
			mu.Lock()
			time.Sleep(1 * time.Millisecond) // Contention point
			mu.Unlock()
		}
	}

	// Force GC pressure with temporary allocations
	for j := 0; j < 100; j++ {
		tmp := make([]int, 1000)
		_ = tmp // Prevent optimization
	}

	// Send result
	ch <- len(data)
}

func main() {
	// Create trace file
	f, err := os.Create("trace Intensive.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Start tracing
	if err := trace.Start(f); err != nil {
		log.Fatal(err)
	}
	defer trace.Stop()

	// Configure runtime for more visibility
	runtime.GOMAXPROCS(4)              // Limit to 4 CPUs for scheduling pressure
	runtime.SetMutexProfileFraction(1) // Enable mutex profiling
	runtime.SetBlockProfileRate(1)     // Enable block profiling

	// Simulate heavy workload
	ch := make(chan int, 100) // Buffered channel
	var wg sync.WaitGroup
	mu := &sync.Mutex{}

	// Launch 1000 goroutines
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go processData(i, &wg, ch, mu)
	}

	// Collect results
	go func() {
		for range ch {
			// Drain channel
		}
		wg.Wait()
	}()

	// Keep running to capture events
	time.Sleep(5 * time.Second) // Allow tracing to capture full activity
}
