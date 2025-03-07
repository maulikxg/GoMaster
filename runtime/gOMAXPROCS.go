package main

import (
	"fmt"
	"runtime"
)

func main() {

	oldProcs := runtime.GOMAXPROCS(0)
	fmt.Printf("Current GOMAXPROCS: %d\n", oldProcs)

	//prev := runtime.GOMAXPROCS(2)
	//fmt.Printf("Old GOMAXPROCS: %d, New GOMAXPROCS: %d\n", prev, runtime.GOMAXPROCS(0))
}
