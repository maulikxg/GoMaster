


# **Go Concurrency Patterns **

---

## **1. Introduction to Concurrency in Go**
### **Why Concurrency?**
Concurrency enables **efficient execution of multiple tasks**, improving application responsiveness and performance. Go achieves this using:
- **Goroutines** - Lightweight threads (`go func()`)
- **Channels** - Safe communication between goroutines (`chan int`)
- **WaitGroups** - Synchronizing multiple goroutines (`sync.WaitGroup`)
- **Mutexes** - Protecting shared resources (`sync.Mutex`)

---

## **2. Worker Pool Pattern**
### **Why Use a Worker Pool?**
- Controls the number of concurrent workers.
- Prevents resource exhaustion.
- Improves task execution efficiency.

### **Implementation**
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

const numWorkers = 3

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(time.Second) // Simulate work
		results <- job * 2 // Process the job
	}
}

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)
	var wg sync.WaitGroup

	// Create worker pool
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Send jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Wait for workers to finish
	wg.Wait()
	close(results)

	// Collect results
	for res := range results {
		fmt.Println("Result:", res)
	}
}
```
### **Key Takeaways**
✔ **Workers process jobs concurrently**  
✔ **WaitGroup ensures all workers complete before exiting**  
✔ **Buffered channels handle job distribution efficiently**

---

## **3. Fan-Out, Fan-In Pattern**
### **Why Fan-Out, Fan-In?**
- **Fan-Out:** Distributes work across multiple workers.
- **Fan-In:** Aggregates results into a single channel.

### **Implementation**
```go
package main

import (
	"fmt"
	"sync"
)

func producer(out chan<- int) {
	for i := 1; i <= 5; i++ {
		out <- i
	}
	close(out)
}

func worker(id int, in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range in {
		fmt.Printf("Worker %d processing %d\n", id, num)
		out <- num * num
	}
}

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)
	var wg sync.WaitGroup

	// Start producer
	go producer(jobs)

	// Fan-Out: Multiple workers
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Close results once all workers are done
	go func() {
		wg.Wait()
		close(results)
	}()

	// Fan-In: Collect results
	for res := range results {
		fmt.Println("Result:", res)
	}
}
```
### **Key Takeaways**
✔ **Multiple workers process jobs concurrently (Fan-Out).**  
✔ **Single collector gathers results (Fan-In).**

---

## **4. Pipeline Pattern**
### **Why Pipelines?**
- **Sequential data processing** in multiple stages.
- **Improves readability** and modularity.

### **Implementation**
```go
package main

import (
	"fmt"
)

func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	numbers := generator(1, 2, 3, 4, 5)
	squares := square(numbers)

	for sq := range squares {
		fmt.Println("Square:", sq)
	}
}
```
### **Key Takeaways**
✔ **Each stage processes data and passes it forward.**  
✔ **Channels link multiple processing steps.**

---

## **5. Cancellation with Context**
### **Why Context?**
- **Graceful shutdown** of goroutines.
- **Timeouts and deadlines** for long-running operations.

### **Implementation**
```go
package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker", id, "stopping")
			return
		default:
			fmt.Println("Worker", id, "working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	for i := 1; i <= 3; i++ {
		go worker(ctx, i)
	}

	time.Sleep(3 * time.Second) // Wait before main exits
}
```
### **Key Takeaways**
✔ **`context.WithTimeout` cancels goroutines after 2 seconds.**  
✔ **Ensures controlled shutdown.**

---

## **6. Select Statement for Multiple Channels**
### **Why Use `select`?**
- Listens to **multiple channels**.
- Prevents **blocking** on a single channel.

### **Implementation**
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Message from channel 1"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Message from channel 2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received:", msg2)
		}
	}
}
```
### **Key Takeaways**
✔ **Handles multiple channel inputs without blocking.**  
✔ **Ensures efficient communication between goroutines.**

---

## **7. Final Summary**
| **Pattern** | **Use Case** |
|------------|-------------|
| **Worker Pool** | Efficient task execution with limited workers |
| **Fan-Out, Fan-In** | Distribute workload, then collect results |
| **Pipeline** | Step-by-step data processing |
| **Context Cancellation** | Graceful shutdown of goroutines |
| **Select Statement** | Manage multiple channels |

---

### **Visualization (Worker-Pool)**
```
+---------+     +---------+     +---------+
| Worker1 |     | Worker2 |     | Worker3 |
+---------+     +---------+     +---------+
|               |               |
+---------+    +---------+    +---------+
|  Job1   |    |  Job2   |    |  Job3   |
+---------+    +---------+    +---------+
```


### **Visualization (Fan-In & Fan-Out)**
```
Producer
   |
   V
+-----------+
|  Channel  |
+-----------+
   |   |   |
   V   V   V
Worker1  Worker2  Worker3  (Fan-Out)
   |       |       |
   V       V       V
+-----------+
|  Results  |
+-----------+
        |
        V
    Aggregator (Fan-In)
```


### **Visualization (Pipeline)**
```
+---------+    +---------+    +---------+
| Input   | -> | Stage 1 | -> | Stage 2 | -> Output
+---------+    +---------+    +---------+
```


### **Visualization**
```
Main Goroutine
       |
       V
Worker1    Worker2    Worker3
       |       |       |
       V       V       V
    +---------------------+
    | Context Cancellation |
    +---------------------+
```



### **Visualization (Select)**
```
+-------------------+
|      Select       |
+-------------------+
   |          |
   V          V
Channel 1   Channel 2
   |          |
   V          V
Response   Response
```