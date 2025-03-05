# 📌 Go Concurrency Safety & Error Handling (`defer`, `panic`, `recover`)

## 🔹 **Concurrency Safety Table**

| Data Type        | Read Safe? | Write Safe? | Concurrent Read & Write Safe? | Notes |
|-----------------|-----------|------------|------------------------------|-------|
| **Slices (`[]T`)** | ✅ Yes | ❌ No | ❌ No | Writing or resizing a slice concurrently causes data races |
| **Maps (`map[K]V`)** | ✅ Yes | ❌ No | ❌ No | Maps are not thread-safe for concurrent read & write |
| **Channels (`chan T`)** | ✅ Yes | ✅ Yes | ✅ Yes | Channels are designed for safe concurrent use |
| **Arrays (`[N]T`)** | ✅ Yes | ❌ No | ❌ No | Arrays are not dynamically resizable like slices |
| **Sync Map (`sync.Map`)** | ✅ Yes | ✅ Yes | ✅ Yes | A thread-safe alternative to `map[K]V` |
| **Mutex (`sync.Mutex`)** | ✅ Yes | ✅ Yes | ✅ Yes | Used to safely modify shared resources |
| **Atomic (`sync/atomic`)** | ✅ Yes | ✅ Yes | ✅ Yes | Provides low-level atomic operations |

---

## ⚠️ **`panic` in Go**

### 📌 **What is `panic`?**
- `panic` is used to **terminate the program immediately** when an unexpected error occurs.
- It **stops normal execution** and begins **unwinding the stack**, executing `defer` statements.

### ✅ **Example of `panic`**:
```go
package main

import "fmt"

func main() {
    fmt.Println("Before Panic")
    panic("Something went wrong!") // This will stop execution
    fmt.Println("After Panic") // This will NOT execute
}
```

### 🔥 **Key Points About `panic`**
1. **Stops normal execution** and starts stack unwinding.
2. **Deferred functions still execute before program termination**.
3. **Use only in truly exceptional cases** (not for regular errors).

---

## 🚀 **`defer` in Go**

### 📌 **What is `defer`?**
- `defer` postpones the execution of a function **until the surrounding function returns**.
- Multiple `defer` calls execute **in Last-In-First-Out (LIFO) order**.

### ✅ **Example of `defer` Execution Order**:
```go
package main

import "fmt"

func main() {
    defer fmt.Println("First Deferred")  
    defer fmt.Println("Second Deferred") 
    defer fmt.Println("Third Deferred")  

    fmt.Println("Main Function Execution")
}
```

### 🔥 **Expected Output**:
```
Main Function Execution
Third Deferred
Second Deferred
First Deferred
```

### ✅ **Defer in Panic Situations**:
```go
package main

import "fmt"

func main() {
    defer fmt.Println("This will execute before panic ends")

    fmt.Println("Before Panic")
    panic("Something bad happened!")
}
```

### 🔥 **Key Points About `defer`**
1. **Executes in LIFO order**.
2. **Always runs before a function exits** (even in case of `panic`).
3. **Useful for cleanup tasks (closing files, releasing locks, etc.)**.

---

## 🔄 **`recover` in Go**

### 📌 **What is `recover`?**
- `recover()` **catches panics** and prevents the program from crashing.
- Can **only be used inside a `defer` function**.

### ✅ **Example of `recover` Handling a Panic**:
```go
package main

import "fmt"

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()

    fmt.Println("Before Panic")
    panic("Something went wrong!")
    fmt.Println("After Panic") // This won't execute
}
```

### 🔥 **Expected Output**:
```
Before Panic
Recovered from panic: Something went wrong!
```

### 🔥 **Key Points About `recover`**
1. **Must be called inside a `defer` function**.
2. **If `recover` is not used, `panic` will crash the program**.
3. **Only catches panics from the same goroutine**.

---

## 🏆 **Best Practices for Concurrency Safety**

✅ **Use `sync.Mutex` to protect shared resources**
```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	m := make(map[int]int)

	go func() {
		mu.Lock()
		m[1] = 10
		mu.Unlock()
	}()

	go func() {
		mu.Lock()
		fmt.Println(m[1])
		mu.Unlock()
	}()
}
```

✅ **Use `sync.Map` for concurrent access to maps**
```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map

	m.Store("key", "value")

	go func() {
		m.Store("newKey", "newValue")
	}()

	go func() {
		if val, ok := m.Load("key"); ok {
			fmt.Println(val)
		}
	}()
}
```

✅ **Use `recover()` to prevent crashes in goroutines**
```go
package main

import "fmt"

func main() {
	go safeFunction()
	select {} // Keep the program running
}

func safeFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	panic("Goroutine Panic!")
}
```

---

## 🚀 **Summary**
| Concept | Description |
|---------|-------------|
| **`panic`** | Used to terminate a program in case of an unrecoverable error |
| **`defer`** | Delays function execution until the surrounding function exits |
| **`recover`** | Handles panics and prevents program crashes |
| **Slices & Maps** | Not thread-safe; must use locks or `sync.Map` |
| **Channels** | Safe for concurrent use |

---

### 📚 **Final Notes**:
- **Use `panic` only for fatal errors**.
- **Use `recover` to handle unexpected crashes, but don’t misuse it**.
- **Avoid modifying maps and slices from multiple goroutines without synchronization**.

