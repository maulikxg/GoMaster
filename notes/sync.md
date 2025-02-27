

# **Deep Dive into `sync.Cond`, `sync.Pool`, and `sync.Map` in Go**

## **1. `sync.Cond`: Synchronization with Conditional Variables**
### **What is `sync.Cond`?**
`sync.Cond` is a **condition variable** used for **goroutine synchronization**. It helps goroutines **wait** until a **certain condition** is met.

### **Key Features of `sync.Cond`**
| Feature           | Description |
|------------------|------------|
| **Used for event-based waiting** | Goroutines wait until a condition is met. |
| **Built on `sync.Mutex` or `sync.RWMutex`** | Requires a lock for synchronization. |
| **Three key methods** | `Wait()`, `Signal()`, and `Broadcast()`. |

### **Methods of `sync.Cond`**
| Method | Description |
|--------|------------|
| `Wait()` | The goroutine **waits** until a condition is met (requires lock). |
| `Signal()` | Wakes up **one** waiting goroutine. |
| `Broadcast()` | Wakes up **all** waiting goroutines. |

---

### **Example: Goroutine Synchronization using `sync.Cond`**
#### **Scenario:** Workers must wait until a job is assigned.

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var cond = sync.NewCond(&sync.Mutex{}) // Initialize condition variable
var ready = false // Shared state

func worker(id int) {
	cond.L.Lock()
	for !ready { // Wait until condition is true
		fmt.Printf("Worker %d is waiting...\n", id)
		cond.Wait() // Release lock & wait for Signal/Broadcast
	}
	fmt.Printf("Worker %d received the job!\n", id)
	cond.L.Unlock()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	// Start 3 worker goroutines
	for i := 1; i <= 3; i++ {
		go func(id int) {
			defer wg.Done()
			worker(id)
		}(i)
	}

	// Simulate some delay before assigning jobs
	time.Sleep(2 * time.Second)

	// Assign job and wake up all workers
	cond.L.Lock()
	ready = true
	cond.Broadcast() // Wake up all workers
	cond.L.Unlock()

	wg.Wait()
}
```

### **Output**
```
Worker 1 is waiting...
Worker 2 is waiting...
Worker 3 is waiting...
Worker 1 received the job!
Worker 2 received the job!
Worker 3 received the job!
```
---
## **2. `sync.Pool`: Object Reuse & Memory Optimization**
### **What is `sync.Pool`?**
`sync.Pool` is a **memory optimization mechanism** that caches **reusable objects** to reduce **garbage collection (GC) overhead**.

### **Key Features of `sync.Pool`**
| Feature | Description |
|---------|------------|
| **Thread-safe** | Multiple goroutines can access it concurrently. |
| **Reduces memory allocation overhead** | Avoids frequent memory allocations. |
| **Cleared during GC** | Objects in the pool are cleared when GC runs. |

### **Methods of `sync.Pool`**
| Method | Description |
|--------|------------|
| `Get()` | Fetches an object from the pool (creates a new one if empty). |
| `Put(obj)` | Returns an object to the pool for reuse. |

---

### **Example: Buffer Reuse using `sync.Pool`**
#### **Scenario:** Reduce memory allocations for buffers.

```go
package main

import (
	"bytes"
	"fmt"
	"sync"
)

var bufferPool = sync.Pool{
	New: func() interface{} {
		fmt.Println("Allocating new buffer")
		return new(bytes.Buffer)
	},
}

func useBuffer(id int) {
	buf := bufferPool.Get().(*bytes.Buffer)
	buf.Reset() // Clear buffer before use
	buf.WriteString(fmt.Sprintf("Worker %d using buffer\n", id))
	fmt.Print(buf.String())
	bufferPool.Put(buf) // Return buffer to pool
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	for i := 1; i <= 3; i++ {
		go func(id int) {
			defer wg.Done()
			useBuffer(id)
		}(i)
	}

	wg.Wait()
}
```

### **Output**
```
Allocating new buffer
Worker 1 using buffer
Allocating new buffer
Worker 2 using buffer
Worker 3 using buffer
```
- Some buffers are **reused**, reducing memory allocation.

---
## **3. `sync.Map`: Thread-Safe Concurrent Map**
### **What is `sync.Map`?**
`sync.Map` is a **concurrent map** optimized for **read-heavy** workloads.

### **Key Features of `sync.Map`**
| Feature | Description |
|---------|------------|
| **Thread-safe** | No need for explicit locking (`sync.Mutex`). |
| **Optimized for high read concurrency** | Works well when reads > writes. |
| **Copy-on-write for updates** | Ensures efficient reads with minimal contention. |

### **Methods of `sync.Map`**
| Method | Description |
|--------|------------|
| `Store(key, value)` | Inserts or updates a key-value pair. |
| `Load(key)` | Retrieves the value for a given key. |
| `LoadOrStore(key, value)` | Retrieves the value if present; otherwise, stores and returns it. |
| `Delete(key)` | Removes a key from the map. |
| `Range(fn)` | Iterates over all key-value pairs. |

---

### **Example: Managing Active User Sessions with `sync.Map`**
#### **Scenario:** Track user login sessions.

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var sessionStore sync.Map

func userLogin(userID string) {
	sessionStore.Store(userID, time.Now()) // Store login time
	fmt.Printf("User %s logged in.\n", userID)
}

func checkSession(userID string) {
	value, exists := sessionStore.Load(userID)
	if exists {
		fmt.Printf("User %s has an active session: %v\n", userID, value)
	} else {
		fmt.Printf("User %s has no active session.\n", userID)
	}
}

func userLogout(userID string) {
	sessionStore.Delete(userID)
	fmt.Printf("User %s logged out.\n", userID)
}

func listActiveSessions() {
	fmt.Println("Active Sessions:")
	sessionStore.Range(func(key, value interface{}) bool {
		fmt.Printf("User: %s, Login Time: %v\n", key, value)
		return true
	})
}

func main() {
	userLogin("Alice")
	userLogin("Bob")

	checkSession("Alice")
	checkSession("Charlie") // Not logged in

	listActiveSessions()

	userLogout("Alice")

	listActiveSessions()
}
```

### **Output**
```
User Alice logged in.
User Bob logged in.
User Alice has an active session: 2025-02-27 12:34:56 +0000 UTC
User Charlie has no active session.
Active Sessions:
User: Alice, Login Time: 2025-02-27 12:34:56 +0000 UTC
User: Bob, Login Time: 2025-02-27 12:35:10 +0000 UTC
User Alice logged out.
Active Sessions:
User: Bob, Login Time: 2025-02-27 12:35:10 +0000 UTC
```

---
# **Final Summary**
| Feature | `sync.Cond` | `sync.Pool` | `sync.Map` |
|---------|------------|------------|------------|
| **Purpose** | Synchronize goroutines waiting for a condition | Reuse objects to reduce GC overhead | Thread-safe map for concurrent access |
| **Best for** | Event-based waiting | Memory optimization | Read-heavy workloads |
| **Key Methods** | `Wait()`, `Signal()`, `Broadcast()` | `Get()`, `Put()` | `Load()`, `Store()`, `Delete()`, `Range()` |





