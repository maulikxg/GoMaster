##  Go Routine Scheduler 

Go’s runtime includes a **user-space scheduler** to efficiently manage **goroutines**. Instead of relying on the OS scheduler for threading, Go has its own lightweight scheduling mechanism that helps in achieving **massive concurrency with low overhead**.

---

## 📌 **1. Understanding Goroutine Scheduling**

### 🔹 **What is a Scheduler?**
A **scheduler** decides **which task (goroutine) runs on which thread (OS thread)** at any given time.

Unlike traditional thread-based models, Go **multiplexes thousands of goroutines over a small number of OS threads** to optimize CPU utilization and avoid excessive context switching.

### 🔹 **Key Features of Go's Scheduler**
✅ **M:N Scheduling** – Go maps **M** goroutines onto **N** OS threads.  
✅ **Work Stealing** – Load balancing across processor threads.  
✅ **Preemptive Scheduling** – Long-running goroutines are interrupted to prevent starvation.  
✅ **Parallel Execution** – Runs goroutines on multiple CPU cores using **GOMAXPROCS**.

---

## 📌 **2. Go’s Scheduler Model: G-P-M Architecture**

Go’s scheduler is based on a **G-P-M Model**:

| Component | Description |
|-----------|------------|
| **G (Goroutine)** | Represents a goroutine. Lightweight thread managed by Go runtime. |
| **P (Processor)** | A logical processor that schedules goroutines. Controls access to the OS thread. |
| **M (Machine)** | Represents an OS thread running goroutines. |

### 🔹 **How G-P-M Works?**
1️⃣ **Goroutines (`G`) are assigned to available Processors (`P`).**  
2️⃣ **Each Processor (`P`) picks goroutines (`G`) from its local run queue.**  
3️⃣ **If a Processor (`P`) is idle, it steals work from another `P`.**  
4️⃣ **Each Processor (`P`) is bound to a single OS thread (`M`) at a time.**

🚀 **Goroutines switch between OS threads as needed but stay within their assigned `P`.**

#### **Example: Understanding G-P-M**
```go
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(2) // Set 2 logical processors

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Goroutine 1 executing")
			time.Sleep(time.Millisecond * 100)
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Goroutine 2 executing")
			time.Sleep(time.Millisecond * 100)
		}
	}()

	time.Sleep(time.Second)
}
```
🔹 **Output:** Goroutines will execute in an interleaved manner across 2 processors.

---

## 📌 **3. Work Stealing Mechanism**

Go’s **Work Stealing Algorithm** prevents **Processor (P) starvation**.

### 🔹 **How Work Stealing Works?**
- Each `P` has a **local run queue** of goroutines.
- When a `P` finishes its tasks, it **steals goroutines** from another `P`'s queue.
- This prevents idle processors and ensures better CPU utilization.

---

## 📌 **4. Scheduling Strategies in Go**

### 🔹 **1. Cooperative Scheduling (Before Go 1.14)**
- Goroutines **yield voluntarily** at function calls (e.g., `runtime.Gosched()`).
- Long-running goroutines could block others.

#### **Example: Yielding with `runtime.Gosched()`**
```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Goroutine executing")
			runtime.Gosched() // Yield execution
		}
	}()

	fmt.Println("Main function executed")
}
```
🔹 **Output:** The scheduler allows the main goroutine to execute.

### 🔹 **2. Preemptive Scheduling (Go 1.14+)**
- The **GC & runtime actively preempts long-running goroutines**.
- If a goroutine **does not yield**, the scheduler **forcibly stops it** to avoid starvation.

#### **Example: Infinite Loop Gets Preempted**
```go
package main

import "fmt"

func main() {
	go func() {
		for {
			fmt.Println("Running infinitely...")
		}
	}()
}
```
🔹 **Output:** The infinite loop won't block other goroutines.

---

## 📌 **5. Goroutine Lifecycle & States**

A **goroutine can be in multiple states** during execution:

| State       | Description |
|------------|------------|
| **Runnable**  | Ready to execute but waiting for a CPU. |
| **Running**  | Currently executing on a CPU core. |
| **Waiting**  | Blocked (I/O, sleep, channel operation). |
| **Dead**  | Completed execution. |

🔹 **Goroutines move between these states as the scheduler manages execution.**

---

## 📌 **6. Goroutine Scheduling Challenges & Optimizations**

### 🔹 **1. Blocking System Calls & OS Threads**
- If a goroutine calls **a blocking system call**, the scheduler **creates a new OS thread** to keep other goroutines running.
- This can increase OS thread usage.

#### **Example: Blocking System Call (Simulated with Sleep)**
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		time.Sleep(2 * time.Second) // Blocks execution
		fmt.Println("Done Sleeping")
	}()
	fmt.Println("Main finished")
	time.Sleep(3 * time.Second) // Prevents main from exiting
}
```
🔹 **Output:** The main function finishes first, while the goroutine is still waiting.

---

### 🔹 **2. Starvation & Preemption**
- **Starvation** occurs if a goroutine **never releases control**.
- **Preemptive scheduling** fixes this by forcing scheduling points.

---

### 🔹 **3. Controlling Parallelism with `GOMAXPROCS`**
`GOMAXPROCS` **controls how many OS threads run goroutines in parallel**.

#### **Example: Setting Parallelism**
```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2) // Limit parallelism to 2 threads
	fmt.Println("Max Procs:", runtime.GOMAXPROCS(0)) // Get current value
}
```

---

## 📌 **7. Optimizing Goroutine Scheduling**

### 🔹 **Best Practices for Efficient Scheduling**
✅ **Limit OS Thread Creation:** Avoid excessive blocking system calls.  
✅ **Use Worker Pools:** Prevent too many active goroutines.  
✅ **Profile Performance:** Use `pprof` to analyze goroutine behavior.  
✅ **Set `GOMAXPROCS` Wisely:** Tune it based on CPU cores.

---

## 📌 **8. Debugging the Scheduler**

### 🔹 **Using `runtime.NumGoroutine()`**
```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Active Goroutines:", runtime.NumGoroutine())
}
```
🔹 **Output:** Shows the current number of running goroutines.

---

## 🚀 **Final Thoughts**
🔹 **Go’s goroutine scheduler is highly optimized for concurrency.**  
🔹 **It uses G-P-M scheduling with work stealing to balance execution.**  
🔹 **Preemptive scheduling ensures fairness & prevents goroutine starvation.**  
🔹 **Optimizing `GOMAXPROCS` and using worker pools improves performance.**

