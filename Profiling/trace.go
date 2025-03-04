package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
	"time"
)

// Simulates goroutines performing blocking and concurrent tasks
func worker(id int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 100) // Simulate some work
		ch <- id * i
	}
}

func main() {
	// Create a trace file
	f, err := os.Create("trace.out")
	if err != nil {
		fmt.Println("Failed to create trace file:", err)
		return
	}
	defer f.Close()

	// Start tracing
	trace.Start(f)
	defer trace.Stop()

	// Simulate workload
	var wg sync.WaitGroup
	ch := make(chan int, 10)

	// Launch multiple workers
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go worker(i, &wg, ch)
	}

	// Read from the channel
	go func() {
		for val := range ch {
			fmt.Println("Received:", val)
		}
	}()

	wg.Wait()
	close(ch)
	fmt.Println("Tracing complete. Run 'go tool trace trace.out' to analyze.")
}
