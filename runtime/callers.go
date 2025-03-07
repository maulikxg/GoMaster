package main

import (
	"fmt"
	"runtime"
)

func main() {

	printCallerInfo()

	exampleFunction()

}

func exampleFunction() {
	trace := make([]uintptr, 10) // Store 10 stack frames
	count := runtime.Callers(3, trace)

	fmt.Printf("Captured %d stack frames\n", count)
	for i := 0; i < count; i++ {
		fn := runtime.FuncForPC(trace[i])
		if fn != nil {
			file, line := fn.FileLine(trace[i])
			fmt.Printf("Frame %d: %s in %s:%d\n", i, fn.Name(), file, line)
		}
	}
}

func printCallerInfo() {
	pc, file, line, ok := runtime.Caller(3) // Skip 1 to get the caller of this function
	if !ok {
		fmt.Println("Could not get caller info")
		return
	}

	// Get function name from program counter
	fn := runtime.FuncForPC(pc)
	fmt.Printf("Caller: %s, File: %s, Line: %d\n", fn.Name(), file, line)
}
