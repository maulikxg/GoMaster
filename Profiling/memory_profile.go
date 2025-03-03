package main

import (
	"fmt"
	"os"
	"runtime/pprof"
)

func allocateMemory() {
	_ = make([]byte, 50*1024*1024) // taking 10 mb
}
func main() {

	f, err := os.Create("mem.prof")
	if err != nil {
		fmt.Println("Error in creation.")
	}
	defer f.Close()

	allocateMemory()

	// Write memory profile
	pprof.WriteHeapProfile(f)

}
