# Go Garbage Collection (GC) 

## 📌 Overview of Go's Garbage Collector
Go uses a **concurrent, tri-color mark-and-sweep garbage collector** to manage memory efficiently.  
It is designed to **minimize pause times** and run in parallel with goroutines.

---

## 1️⃣ GC Cycles & Its Phases

### 🔹 **Understanding GC Cycles**  
Garbage Collection in Go happens **in cycles** where the memory is analyzed and cleaned up.  
Each cycle consists of **3 main phases**:

### **1️⃣ Mark Phase (Finding Live Objects)**
- The **GC scans all objects** starting from **roots** (stack, globals, heap references).  
- It follows references to find **reachable** objects and **marks them as live**.

### **2️⃣ Sweep Phase (Reclaiming Unused Memory)**
- Unmarked objects are **considered garbage** and their memory is **reclaimed**.  
- This process is done **incrementally** to reduce pause times.

### **3️⃣ Scavenge Phase (Returning Unused Memory to OS)**
- If large blocks of memory remain unused **for a long time**, they are returned to the **OS**.

---

## 2️⃣ The Tri-Color Marking Algorithm (How Marking Works)

Go's GC uses a **Tri-Color Mark & Sweep Algorithm** to track object reachability.  
It categorizes objects into **three colors**:

| Color | Meaning |
|-------|---------|
| 🖤 **Black** | Live objects (reachable, fully scanned) |
| ⚪ **White** | Garbage (unreachable, will be deleted) |
| ⚫ **Gray** | Objects that are **reachable but not yet fully scanned** |

### 🔹 **How the Tri-Color Algorithm Works?**
1. **Start with Roots** → Move **roots (globals, stacks)** to the **gray set**.
2. **Scan Gray Objects** → Move **referenced objects** to **gray** and mark current as **black**.
3. **Repeat Until Gray is Empty** → Everything left in white is garbage.
4. **Sweep Phase** → Free memory of **white** objects.

#### **Example: Understanding Marking in Action**
```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println("Heap Allocation Before GC:", m.HeapAlloc)

	// Force GC
	runtime.GC()

	runtime.ReadMemStats(&m)
	fmt.Println("Heap Allocation After GC:", m.HeapAlloc)
}
```
**🔍 Output:** Shows how much memory was freed after GC.

---

## 3️⃣ Finalizers & Memory Limits

### 🔹 **Finalizers in Go**
A **finalizer** runs a function before an object is garbage collected.  
🚨 **Finalizers should not be used for critical cleanup!** (use `defer` instead).

#### **Example: Using `runtime.SetFinalizer`**
```go
package main

import (
	"fmt"
	"runtime"
)

type Resource struct{}

func cleanup(r *Resource) {
	fmt.Println("Finalizer executed! Cleaning up.")
}

func main() {
	r := &Resource{}
	runtime.SetFinalizer(r, cleanup)
}
```
🔹 **Use case:** Logging when an object is freed.

---

## 4️⃣ GOGC (Garbage Collection Goal)
`GOGC` **controls GC frequency** by adjusting heap growth.

### 🔹 **GOGC Behavior**
| GOGC Value | Behavior |
|------------|----------|
| `100` (default) | GC runs when heap **doubles** |
| `200` | GC runs **less often** (heap triples) |
| `50` | GC runs **more frequently** (heap grows 50%) |
| `-1` | **Disables GC** |

#### **Example: Changing GOGC**
```go
package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	fmt.Println("Default GOGC:", debug.SetGCPercent(-1)) // Get current value
	debug.SetGCPercent(200) // Increase GOGC (less frequent GC)
	fmt.Println("Updated GOGC:", debug.SetGCPercent(-1))
}
```

---

## 5️⃣ GOMEMLIMIT (Memory Limits in Go)
`GOMEMLIMIT` **sets a hard memory limit** for heap usage.

#### **Example: Setting a Memory Limit**
```go
package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	debug.SetMemoryLimit(200 * 1024 * 1024) // 200MB limit
	fmt.Println("Memory limit set to 200MB")
}
```

---

## 6️⃣ Monitoring GC with `runtime.MemStats`

### 🔹 **Why Use `runtime.MemStats`?**
✅ **Detect memory leaks**  
✅ **Monitor GC frequency**  
✅ **Track heap usage**  

#### **Example: Reading Memory Stats**
```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Println("Heap Allocation:", m.HeapAlloc)
	fmt.Println("Total Allocation:", m.TotalAlloc)
	fmt.Println("Heap System:", m.HeapSys)
	fmt.Println("GC Runs:", m.NumGC)
}
```

---

## 7️⃣ Optimization Guide for GC Performance

### 🔹 Best Practices for Optimizing GC in Go
✅ **Reduce GC pressure:** Use **object pooling** (`sync.Pool`).  
✅ **Optimize GOGC:** Experiment with `GOGC=50` or `GOGC=200`.  
✅ **Limit memory usage:** Use `GOMEMLIMIT`.  
✅ **Profile memory usage:** Use `runtime.MemStats`.  

#### **Example: Using `sync.Pool` to Reduce Allocations**
```go
package main

import (
	"fmt"
	"sync"
)

var pool = sync.Pool{
	New: func() interface{} { return new(int) },
}

func main() {
	num := pool.Get().(*int) // Reuse object
	*num = 42
	fmt.Println("Value:", *num)
	pool.Put(num) // Return to pool
}
```

---

## 📌 Summary Table

| Feature         | Purpose                                    | Default |
|----------------|--------------------------------------------|---------|
| **GOGC**       | Controls **GC frequency**                 | `100`   |
| **GOMEMLIMIT** | Sets **hard memory limit**                | No Limit |
| **Finalizers** | Runs cleanup code before object deletion  | N/A     |
| **runtime.MemStats** | Monitors memory and GC usage        | N/A     |

---

## 🚀 When to Use What?
| Situation | Use `GOGC`? | Use `GOMEMLIMIT`? | Use `runtime.MemStats`? |
|--------------|---------------|-----------------|----------------|
| **App uses too much memory?** | ✅ Increase GOGC | ✅ Set GOMEMLIMIT | ✅ Track HeapAlloc |
| **Reduce GC pauses?** | ✅ Increase GOGC | ❌ | ✅ Check NumGC |
| **Memory leak suspected?** | ❌ | ❌ | ✅ Monitor HeapAlloc & TotalAlloc |
| **Prevent OOM crashes?** | ❌ | ✅ Set limit | ✅ Check HeapSys |

---



