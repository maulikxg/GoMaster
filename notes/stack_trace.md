# Go Stack Traces with `runtime/debug.Stack()` and `runtime.Stack()`

Stack traces in Go are crucial for debugging runtime issues, especially in concurrent applications. The `runtime/debug` and `runtime` packages provide methods to capture stack traces efficiently.

## **1. `runtime/debug.Stack()`**
### **Overview**
- Captures a **formatted** stack trace of the current goroutine.
- Provides a **human-readable** output with function calls and line numbers.
- Useful for debugging without crashing the program.

### **Example: Capturing Stack Trace in a Goroutine**
```go
package main

import (
	"fmt"
	"runtime/debug"
	"time"
)

func worker() {
	fmt.Println("Worker running...")
	fmt.Println(string(debug.Stack()))
}

func main() {
	go worker()
	time.Sleep(time.Second) // Allow goroutine to complete
}
```
### **Output Example**
```
goroutine 6 [running]:
runtime/debug.Stack(0xc000048768, 0x0, 0x0)
	/usr/local/go/src/runtime/debug/stack.go:24 +0x9f
main.worker()
	/tmp/prog.go:10 +0x65
created by main.main
	/tmp/prog.go:14 +0x45
```
### **Key Takeaways**
- Captures only the **current goroutine**.
- Shows function calls, file names, and line numbers.
- Ideal for debugging inside a specific function.

---

## **2. `runtime.Stack()`**
### **Overview**
- Captures **raw** stack traces for the current or all goroutines.
- Returns a **byte slice**, requiring conversion to a string.
- More **performance-efficient** than `debug.Stack()`.

### **Example: Capturing All Goroutines**
```go
package main

import (
	"fmt"
	"runtime"
	"time"
)

func worker() {
	for i := 0; i < 3; i++ {
		fmt.Printf("Worker %d running...\n", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go worker()
	go worker()
	time.Sleep(1 * time.Second)

	buf := make([]byte, 4096) // Allocate buffer for stack trace
	n := runtime.Stack(buf, true)
	fmt.Println("Stack trace for all goroutines:\n", string(buf[:n]))
}
```
### **Output Example**
```
goroutine 1 [running]:
main.main()
	/tmp/prog.go:18 +0x65
goroutine 6 [sleep]:
time.Sleep()
	/usr/local/go/src/runtime/time.go:188 +0xbf
main.worker()
	/tmp/prog.go:10 +0x85
goroutine 7 [sleep]:
time.Sleep()
	/usr/local/go/src/runtime/time.go:188 +0xbf
main.worker()
	/tmp/prog.go:10 +0x85
```
### **Key Takeaways**
- Captures stack traces of **all goroutines** when `true` is passed as the second argument.
- Helps debug concurrency issues, such as goroutine leaks or deadlocks.
- Requires manual buffer allocation.

---

