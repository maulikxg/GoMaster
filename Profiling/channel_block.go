package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

func blockedChannel() {
	ch := make(chan int)
	go func() {
		time.Sleep(3 * time.Second) // Simulating delay
		ch <- 10
	}()
	<-ch // Blocks here
}

func main() {
	// Enable block profiling
	runtime.SetBlockProfileRate(1)

	// Create profile file
	f, err := os.Create("blockk.prof")
	if err != nil {
		fmt.Println("Could not create block profile:", err)
		return
	}
	defer f.Close()

	// Run the function
	blockedChannel()

	// Write block profile data to file
	pprof.Lookup("block").WriteTo(f, 0)
}
