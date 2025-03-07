package main

import (
	"fmt"
	"runtime"
	"time"
)

func worke(id int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Worker %d: iteration %d\n", id, i)
		time.Sleep(time.Millisecond * 500)
	}
}

func profileGoroutines() {
	// First, find out how many goroutines exist
	numGoroutines := runtime.NumGoroutine()

	// Create a slice to hold stack records
	profile := make([]runtime.StackRecord, numGoroutines)

	// Collect the goroutine profile
	n, ok := runtime.GoroutineProfile(profile)
	if !ok {
		fmt.Println("Failed to get goroutine profile")
		return
	}

	// Print collected profiles
	fmt.Printf("Total goroutines: %d\n", n)
	for i, record := range profile[:n] {
		fmt.Printf("Goroutine %d: %+v\n", i+1, record)
	}
}

func main() {
	// Start some worker goroutines
	go worke(1)
	go worke(2)
	go worke(3)

	// Allow some time for goroutines to run
	time.Sleep(time.Second * 2)

	// Profile the goroutines
	profileGoroutines()
}
