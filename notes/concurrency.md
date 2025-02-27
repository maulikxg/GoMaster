
# Go Concurrency Patterns 


---

## 📌 Introduction to Concurrency in Go
Concurrency in Go is based on the **Communicating Sequential Processes (CSP)** model, where independent processes (goroutines) communicate via **channels** instead of sharing memory.

### ✅ Why Use Concurrency?
- Improves **performance** by utilizing multi-core CPUs.
- Handles **I/O-bound** and **CPU-bound** tasks efficiently.
- Enables **parallel execution** without complex thread management.

### ⚡ Key Concurrency Features in Go:
1. **Goroutines** – Lightweight managed threads.
2. **Channels** – Communication between goroutines.
3. **select Statement** – Multi-channel operations.
4. **Synchronization Primitives** – `sync.Mutex`, `sync.WaitGroup`, `sync.Cond`.

---

## 🏃‍♂️ 1. Goroutines – Lightweight Threads

A **goroutine** is a function executing independently in the background.

### 🔹 Creating a Goroutine
```go
package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello, Go!")
}

func main() {
	go sayHello() // Starts a goroutine
	time.Sleep(time.Second) // Prevents main from exiting immediately
}
```
**Key Points:**
- `go sayHello()` runs `sayHello` concurrently.
- The main function **must not exit** before goroutines finish.
- `time.Sleep` is used here for simplicity, but **use proper synchronization**.

### ⚠️ Goroutine Pitfalls:
- Goroutines do **not** return values directly.
- Without synchronization, they may exit before completing.

---

## 🔄 2. Channels – Safe Communication Between Goroutines

A **channel** enables safe data transfer between goroutines.

### 🔹 Creating a Channel
```go
ch := make(chan int) // Unbuffered channel
```

### 🔹 Sending and Receiving Data
```go
package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		ch <- 42 // Send data
	}()

	val := <-ch // Receive data
	fmt.Println(val) // Output: 42
}
```
**Key Takeaways:**
- Sending and receiving are **blocking operations**.
- Data exchange happens only when both sender and receiver are ready.

### 📦 Buffered Channels – Asynchronous Communication
Buffered channels allow multiple values to be stored **without blocking**.

```go
ch := make(chan int, 3) // Capacity of 3
ch <- 1
ch <- 2
ch <- 3
fmt.Println(<-ch) // 1 (oldest value)
```
- If the buffer is full, **sending blocks** until space is available.
- If the buffer is empty, **receiving blocks** until data is available.

---

## ⚙️ 3. Concurrency Patterns in Go

### 📌 3.1 Generator Pattern – Streaming Data

A **generator function** produces values and sends them through a channel.

```go
package main

import "fmt"

func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out) // Prevents deadlock
	}()
	return out
}

func main() {
	ch := generator(2, 3, 5, 7)
	for num := range ch {
		fmt.Println(num) // 2, 3, 5, 7
	}
}
```
**Advantages:**
- Separates **data production** from **processing**.
- Avoids **tight coupling** between producer and consumer.

---

### 📌 3.2 Fan-Out and Fan-In – Parallel Processing

#### **Fan-Out** – Distribute Work Across Multiple Goroutines
Multiple goroutines consume tasks from a single channel.

```go
package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(time.Second) // Simulate work
		results <- job * 2
	}
}

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results) // Three concurrent workers
	}

	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	for k := 1; k <= 5; k++ {
		fmt.Println(<-results) // Processed results
	}
}
```
- **Increases throughput** by distributing tasks.
- Ensures **parallel execution** of independent tasks.

#### **Fan-In** – Merging Multiple Channels into One
```go
func merge(ch1, ch2 <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for val := range ch1 {
			out <- val
		}
		for val := range ch2 {
			out <- val
		}
		close(out)
	}()
	return out
}
```
- Combines **multiple producers** into a **single output**.

---

### 📌 3.3 Timeout with `select`
Prevent blocking by setting a **timeout**.

```go
select {
case msg := <-ch:
    fmt.Println("Received:", msg)
case <-time.After(time.Second):
    fmt.Println("Timeout!")
}
```
- **Useful for network calls, database queries, etc.**

---

### 📌 3.4 Cancellation Using a Done Channel
Gracefully stop goroutines.

```go
func worker(done chan bool) {
	fmt.Println("Working...")
	time.Sleep(time.Second)
	done <- true
}

func main() {
	done := make(chan bool)
	go worker(done)
	<-done // Wait for completion
}
```
- Allows **controlled shutdown** of goroutines.

---

## 🔐 4. Synchronization Mechanisms

### ✅ `sync.WaitGroup` – Waiting for Multiple Goroutines
```go
var wg sync.WaitGroup

wg.Add(3) // Number of goroutines

go func() {
	defer wg.Done()
	fmt.Println("Task 1 completed")
}()

go func() {
	defer wg.Done()
	fmt.Println("Task 2 completed")
}()

wg.Wait() // Block until all goroutines finish
```

### ✅ `sync.Mutex` – Protect Shared Resources
```go
var mu sync.Mutex

mu.Lock()
// Critical section
mu.Unlock()
```
- Prevents **race conditions** in concurrent programs.

---

## 🚀 Best Practices for Go Concurrency
✅ **Prefer Channels Over Shared Memory** – Avoids complex locking.  
✅ **Use Buffered Channels for Performance** – Reduces unnecessary blocking.  
✅ **Always Close Channels Properly** – Prevents **goroutine leaks**.  
✅ **Use `select` for Multi-Channel Operations** – Enables **non-blocking execution**.  
✅ **Avoid Global Variables in Concurrent Code** – Reduces **race conditions**.

---

