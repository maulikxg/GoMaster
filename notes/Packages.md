
## **1️⃣ `runtime` Package – Go's Low-Level Powerhouse**

The `runtime` package provides access to Go's runtime system, including:
- Goroutine scheduling (`G-P-M` model)
- Memory management and garbage collection (`GOGC`, `GOMEMLIMIT`)
- Stack growth and preemption
- Debugging and performance monitoring

### **1.1 Controlling Goroutines**

#### **Yielding Execution (`Gosched()`)**
This function allows a goroutine to yield execution so that other goroutines can run.

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
			runtime.Gosched() // Yields to other goroutines
		}
	}()

	fmt.Println("Main function executing")
}
```
This ensures the scheduler can switch between goroutines.

#### **Getting Goroutine Count (`NumGoroutine()`)**
To check how many goroutines are currently running:

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Goroutines running:", runtime.NumGoroutine())
}
```

---

### **1.2 Memory Management and GC**

#### **Setting Garbage Collection Aggressiveness (`GOGC`)**
`GOGC` controls how often the garbage collector runs.

- Default is `100` (GC runs when heap grows by 100%).
- `GOGC=50` makes GC run more frequently.
- `GOGC=200` makes GC run less frequently.
- `GOGC=off` disables GC.

Example:

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Default GOGC:", runtime.GOGC()) // Get current value
	runtime.GC() // Force garbage collection
	fmt.Println("Garbage collection triggered")
}
```

#### **Memory Limit (`GOMEMLIMIT`)**
Introduced in Go 1.19, this sets an upper bound on memory usage. If memory usage exceeds the limit, GC becomes more aggressive.

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Current Memory Limit:", runtime.MemStats{}.HeapAlloc)
	runtime.GC()
}
```

---

### **1.3 Debugging and Performance Monitoring**

#### **Inspecting Memory Usage (`MemStats`)**
`runtime.MemStats` provides detailed memory usage statistics.

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println("Heap Alloc:", m.HeapAlloc)
	fmt.Println("Total Alloc:", m.TotalAlloc)
	fmt.Println("Sys:", m.Sys)
	fmt.Println("NumGC:", m.NumGC)
}
```

#### **Stack Growth and Preemption**
Go dynamically increases goroutine stack sizes. To monitor stack usage:

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Stack size:", runtime.Stack(nil, false))
}
```

---

# **2️⃣ Deep Dive into `reflect` Package**

The `reflect` package in Go allows inspection and manipulation of types at runtime. It enables dynamic programming techniques such as generic operations, serialization, and deep copying.

## **2.1 Why Use `reflect`?**
Go is a statically typed language, meaning types are determined at compile time. However, sometimes you need to work with unknown types dynamically (e.g., in frameworks, testing, or serialization). The `reflect` package provides this capability by allowing:
- **Type Inspection** – Determine the type and kind of a value.
- **Value Manipulation** – Modify values dynamically.
- **Struct Field Access** – Read and write fields of structs dynamically.
- **Calling Methods Dynamically** – Invoke methods on unknown types.

---

## **2.2 Key Types in `reflect`**
### **2.2.1 `reflect.Type` – Describes a Go Type**
`reflect.Type` allows introspection of a type at runtime.

```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num int = 10
	t := reflect.TypeOf(num)

	fmt.Println("Type:", t)                  // int
	fmt.Println("Kind:", t.Kind())            // int
	fmt.Println("Size:", t.Size(), "bytes")   // Platform dependent
	fmt.Println("Align:", t.Align())          // Memory alignment
}
```

**Key Functions:**
- `reflect.TypeOf(x)` → Returns the `reflect.Type` of `x`.
- `t.Kind()` → Returns the **kind** (basic classification) of the type.
- `t.Name()` → Returns the **name** of the type.

---

### **2.2.2 `reflect.Value` – Holds a Value and Allows Manipulation**
`reflect.Value` allows you to retrieve and modify the value dynamically.

```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var num int = 42
	v := reflect.ValueOf(num)

	fmt.Println("Value:", v)             // 42
	fmt.Println("Type:", v.Type())        // int
	fmt.Println("Can Set:", v.CanSet())   // false (because it's not addressable)
}
```
**Important:**
- `reflect.ValueOf(x)` → Returns the `reflect.Value` of `x`.
- `.CanSet()` → Checks if the value is modifiable.

To modify a value, you need to pass a pointer:

```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	num := 10
	v := reflect.ValueOf(&num).Elem() // Get the actual value

	fmt.Println("Before:", num)  
	v.SetInt(50) // Modifying value
	fmt.Println("After:", num)  
}
```

---

## **2.3 Deep Dive into Common Operations**

### **2.3.1 Comparing Values – `reflect.DeepEqual`**
Go does not allow direct comparison of slices, maps, or structs with unexported fields. `reflect.DeepEqual` enables deep comparison.

```go
package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	c := []int{1, 2, 4}

	fmt.Println(reflect.DeepEqual(a, b)) // true
	fmt.Println(reflect.DeepEqual(a, c)) // false
}
```

---

### **2.3.2 Working with Structs Dynamically**
You can access and modify struct fields at runtime.

```go
package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{"Alice", 25}
	t := reflect.TypeOf(p)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("Field %d: %s (%s)\n", i, field.Name, field.Type)
	}
}
```

**Modifying Struct Fields Dynamically**  
To modify fields, you need to pass a pointer:

```go
package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{"Alice", 25}
	v := reflect.ValueOf(&p).Elem() // Get the actual struct

	v.FieldByName("Age").SetInt(30) // Modify field
	fmt.Println(p)                  // {Alice 30}
}
```

---

### **2.3.3 Calling Methods Dynamically**
If a type has methods, you can invoke them dynamically.

```go
package main

import (
	"fmt"
	"reflect"
)

type Math struct{}

func (Math) Add(a, b int) int {
	return a + b
}

func main() {
	m := Math{}
	v := reflect.ValueOf(m)
	method := v.MethodByName("Add")

	result := method.Call([]reflect.Value{reflect.ValueOf(3), reflect.ValueOf(5)})
	fmt.Println("Result:", result[0].Int()) // 8
}
```

---

## **2.4 Edge Cases and Best Practices**

### **Unexported Fields Are Inaccessible**
Trying to modify an unexported field will cause a panic.

```go
type secret struct {
	value int
}

// This will panic
// v.FieldByName("value").SetInt(50)
```
Use `unsafe` to bypass this (covered later).

### **Performance Considerations**
- `reflect` is **slow**. Avoid it in performance-critical code.
- Prefer interfaces instead of using `reflect` where possible.
- Use caching to minimize repetitive reflection calls.

---
# **3️⃣ Deep Dive into `strconv` Package**

The `strconv` package provides functions for converting between strings and basic data types like integers, floats, and booleans.

---

## **3.1 Why Use `strconv`?**
- Go does not allow implicit type conversion (e.g., `int` to `string`).
- Converting numbers to strings and vice versa is a common task in programming.
- `fmt.Sprintf()` can also be used, but `strconv` is more efficient.

---

## **3.2 Converting Between Strings and Numbers**

### **3.2.1 String to Integer (`strconv.Atoi`, `strconv.ParseInt`)**

#### **Using `Atoi` (String → Int, Base 10)**
`strconv.Atoi` converts a string to an integer but only supports base 10.

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	num, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Integer:", num) // 123
}
```
- Returns `int` and an `error` (in case the string is not a valid number).
- Only works for base-10 numbers.

#### **Using `ParseInt` (Supports Different Bases)**
If you need to parse numbers in different bases (binary, octal, hex), use `strconv.ParseInt`.

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	num, err := strconv.ParseInt("1101", 2, 64) // Binary (base 2)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Parsed Integer:", num) // 13
}
```
- The second parameter is the **base** (2 for binary, 10 for decimal, 16 for hexadecimal, etc.).
- The third parameter is the **bit size** (0, 8, 16, 32, 64).

---

### **3.2.2 Integer to String (`strconv.Itoa`, `strconv.FormatInt`)**

#### **Using `Itoa` (Int → String, Base 10)**
Converts an integer to a string.

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := strconv.Itoa(456)
	fmt.Println("String:", str) // "456"
}
```

#### **Using `FormatInt` (Supports Different Bases)**
For other bases, use `strconv.FormatInt`.

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := strconv.FormatInt(255, 16) // Convert to Hexadecimal
	fmt.Println("Hex:", str) // "ff"
}
```

---

## **3.3 Converting Between Strings and Floats**

### **String to Float (`strconv.ParseFloat`)**
Parses a string into a floating-point number.

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	num, err := strconv.ParseFloat("3.1415", 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Parsed Float:", num) // 3.1415
}
```
- The second argument specifies precision (`32` for `float32`, `64` for `float64`).

---

### **Float to String (`strconv.FormatFloat`)**
Converts a float to a string with formatting options.

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := strconv.FormatFloat(3.1415, 'f', 2, 64)
	fmt.Println("Formatted Float:", str) // "3.14"
}
```
- `'f'` → Use fixed-point notation.
- `'e'` → Use scientific notation.
- `'g'` → Automatically choose based on precision.
- The third parameter (`2`) specifies the number of decimal places.

---

## **3.4 Converting Between Strings and Booleans**

### **String to Boolean (`strconv.ParseBool`)**
Parses a string into a boolean (`true` or `false`).

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	b, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Boolean:", b) // true
}
```
Accepts `"1"`, `"t"`, `"T"`, `"true"`, `"TRUE"` as `true`.  
Accepts `"0"`, `"f"`, `"F"`, `"false"`, `"FALSE"` as `false`.

---

### **Boolean to String (`strconv.FormatBool`)**
Converts a boolean to a string.

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := strconv.FormatBool(false)
	fmt.Println("String:", str) // "false"
}
```

---

## **3.5 Edge Cases and Best Practices**

### **Handling Invalid Inputs**
If a string is not a valid number, `strconv` returns an error.

```go
num, err := strconv.Atoi("abc")
if err != nil {
	fmt.Println("Conversion failed:", err)
}
```

### **Using `fmt.Sprintf` vs `strconv`**
`fmt.Sprintf("%d", num)` can also convert numbers to strings but is slower than `strconv.Itoa()`. Use `strconv` for performance-critical code.

---
# **4️⃣ Deep Dive into `unsafe` Package**

The `unsafe` package in Go provides low-level memory manipulation capabilities, allowing developers to bypass Go’s type safety rules. It is primarily used for:
- Direct memory access.
- Converting between different types without explicit conversion.
- Working with raw pointers and structs.
- Optimizing performance in special cases.

⚠️ **Warning:** `unsafe` is powerful but dangerous. It should only be used when absolutely necessary, as it can lead to memory corruption, undefined behavior, and non-portable code.

---

## **4.1 Key Types and Functions in `unsafe`**

| **Function / Type**        | **Description** |
|---------------------------|----------------|
| `unsafe.Pointer`          | A generic pointer type that can be converted between different pointer types. |
| `uintptr`                 | An integer type that can hold a memory address (useful for pointer arithmetic). |
| `unsafe.Sizeof(x)`        | Returns the size of `x` in bytes. |
| `unsafe.Alignof(x)`       | Returns the alignment requirement of `x`. |
| `unsafe.Offsetof(field)`  | Returns the byte offset of a field within a struct. |

---

## **4.2 `unsafe.Pointer` – The Universal Pointer**

Unlike `*T` (typed pointers), `unsafe.Pointer` can be used to convert between different pointer types.

### **Converting Between Pointer Types**
```go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var num int = 42
	var ptr *int = &num
	var uptr unsafe.Pointer = unsafe.Pointer(ptr)  // Convert *int → unsafe.Pointer
	var floatPtr *float64 = (*float64)(uptr)       // Convert unsafe.Pointer → *float64

	fmt.Println(*floatPtr) // Undefined behavior (memory layout mismatch)
}
```
⚠️ **Caution:** `floatPtr` interprets an integer as a float, leading to garbage values.

---

## **4.3 `uintptr` – Storing Memory Addresses**

`uintptr` is an integer type that can store a memory address. It enables pointer arithmetic but should be used carefully.

### **Example: Using `uintptr` for Offset Calculation**
```go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	arr := [4]int{10, 20, 30, 40}
	ptr := unsafe.Pointer(&arr[0])      // Pointer to first element
	nextPtr := unsafe.Pointer(uintptr(ptr) + unsafe.Sizeof(arr[0])) // Move to next element

	fmt.Println(*(*int)(nextPtr)) // 20
}
```
- We use `uintptr(ptr) + unsafe.Sizeof(arr[0])` to move the pointer forward by the size of an `int`.
- The result is cast back to `*int` to retrieve the next value.

---

## **4.4 Inspecting Memory Layout with `unsafe.Sizeof`, `unsafe.Alignof`, `unsafe.Offsetof`**

### **Getting Size and Alignment**
```go
package main

import (
	"fmt"
	"unsafe"
)

type Data struct {
	A int8
	B int64
	C int16
}

func main() {
	var d Data
	fmt.Println("Size of Data:", unsafe.Sizeof(d))  // 16 or more (depends on padding)
	fmt.Println("Alignment of B:", unsafe.Alignof(d.B)) // 8 (on most platforms)
}
```
- **`Sizeof(x)`** → Returns the memory occupied by `x`.
- **`Alignof(x)`** → Returns the required alignment for `x`.
- **Padding occurs** due to alignment rules, leading to extra memory usage.

### **Finding Struct Field Offsets**
```go
package main

import (
	"fmt"
	"unsafe"
)

type Example struct {
	X int8
	Y int64
	Z int16
}

func main() {
	var e Example
	fmt.Println("Offset of X:", unsafe.Offsetof(e.X)) // 0
	fmt.Println("Offset of Y:", unsafe.Offsetof(e.Y)) // 8 (due to alignment)
	fmt.Println("Offset of Z:", unsafe.Offsetof(e.Z)) // 16 (padded for alignment)
}
```
- `Y` starts at offset `8` instead of `1`, due to memory alignment.
- `Z` is pushed to `16` due to padding.

---

## **4.5 Accessing Unexported Struct Fields**

In Go, unexported struct fields cannot be accessed directly outside their package. `unsafe` allows bypassing this restriction.

### **Example: Accessing Private Fields**
```go
package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

type secret struct {
	private int
}

func main() {
	s := secret{private: 42}
	ptr := unsafe.Pointer(&s)                         // Get pointer to struct
	fieldPtr := (*int)(unsafe.Pointer(uintptr(ptr)))  // Convert to int pointer
	*fieldPtr = 99                                    // Modify the field

	fmt.Println(s) // {99}
}
```
⚠️ **Caution:** This breaks Go’s encapsulation rules and should be avoided in production.

---

## **4.6 Using `unsafe` for Performance Optimization**

### **Avoiding Slice Reallocation in High-Performance Code**
A slice’s underlying array may be reallocated if its capacity is exceeded. Using `unsafe`, we can reuse the memory efficiently.

```go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	data := []int{1, 2, 3}
	header := (*[3]int)(unsafe.Pointer(&data[0])) // Create a fixed-size array view

	fmt.Println(header[1]) // 2
}
```
- This technique avoids extra allocations but should be used cautiously.

---

## **4.7 Best Practices and When to Use `unsafe`**

✅ **Use `unsafe` when:**
- Interfacing with low-level system calls.
- Optimizing memory layout in high-performance applications.
- Accessing raw memory in special cases (e.g., embedded systems, networking).

❌ **Avoid `unsafe` when:**
- Writing standard Go applications.
- Handling unexported fields (violates encapsulation).
- Working with arbitrary memory locations (can cause crashes).

---

# **5️⃣ Deep Dive into `os` Package**

The `os` package in Go provides a platform-independent way to interact with the underlying operating system. It allows handling:
- File and directory manipulation
- Environment variables
- Process management
- Standard input/output (I/O) operations

---

## **5.1 File Handling with `os` Package**

### **5.1.1 Creating a File (`os.Create`)**

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close() // Ensure file is closed properly

	fmt.Println("File created successfully:", file.Name())
}
```
- `os.Create()` creates a new file or truncates an existing file.
- `defer file.Close()` ensures the file is properly closed after execution.

---

### **5.1.2 Writing to a File (`os.Write`, `os.WriteString`)**

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("Hello, Go!")
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}
}
```
- `file.WriteString("text")` writes a string to the file.
- `_` ignores the byte count returned by `WriteString()`.

---

### **5.1.3 Reading a File (`os.ReadFile`)**

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("File contents:", string(data))
}
```
- `os.ReadFile("filename")` reads the entire file into memory.

---

### **5.1.4 Appending to a File (`os.OpenFile` with `O_APPEND`)**

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("example.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("\nAppended text.")
	if err != nil {
		fmt.Println("Error appending to file:", err)
	}
}
```
- `os.O_APPEND` opens the file for appending.
- `os.O_WRONLY` ensures the file is opened in write mode.

---

### **5.1.5 Deleting a File (`os.Remove`)**

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Remove("example.txt")
	if err != nil {
		fmt.Println("Error deleting file:", err)
		return
	}
	fmt.Println("File deleted successfully")
}
```
- `os.Remove("filename")` deletes the specified file.

---

## **5.2 Directory Handling**

### **5.2.1 Creating a Directory (`os.Mkdir` & `os.MkdirAll`)**

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Mkdir("mydir", 0755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}
	fmt.Println("Directory created successfully")
}
```
- `0755` sets **permissions** (read/write/execute for the owner, read/execute for others).

For nested directories, use `os.MkdirAll()`:
```go
os.MkdirAll("parent/child", 0755)
```

---

### **5.2.2 Listing Files in a Directory (`os.ReadDir`)**

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	files, err := os.ReadDir(".") // List current directory files
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
```
- `os.ReadDir(".")` reads the contents of the current directory.

---

### **5.2.3 Removing a Directory (`os.RemoveAll`)**

```go
os.RemoveAll("mydir") // Deletes the directory and its contents
```
⚠️ **Caution:** This permanently deletes everything inside the directory.

---

## **5.3 Environment Variables**

### **5.3.1 Getting an Environment Variable (`os.Getenv`)**
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	path := os.Getenv("PATH")
	fmt.Println("PATH:", path)
}
```

---

### **5.3.2 Setting an Environment Variable (`os.Setenv`)**

```go
os.Setenv("MY_VAR", "GoLang")
fmt.Println(os.Getenv("MY_VAR")) // GoLang
```

---

### **5.3.3 Listing All Environment Variables (`os.Environ`)**

```go
for _, env := range os.Environ() {
	fmt.Println(env)
}
```
- `os.Environ()` returns a slice of `"KEY=VALUE"` strings.

---

## **5.4 Process Management**

### **5.4.1 Getting the Process ID (`os.Getpid`)**

```go
fmt.Println("Process ID:", os.Getpid()) // Prints current process ID
```

---

### **5.4.2 Executing External Commands (`os.StartProcess`)**

```go
package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("echo", "Hello, Go!")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(output)) // "Hello, Go!"
}
```
- `exec.Command("cmd", "arg1", "arg2")` executes a system command.

---

### **5.4.3 Exiting the Program (`os.Exit`)**
```go
os.Exit(1) // Exits the program with status code 1
```

---

## **5.5 File Locking with `os/sync`**

For concurrent access, the `os/sync` package provides `sync.Mutex` for locking.

```go
package main

import (
	"fmt"
	"os"
	"sync"
)

var lock sync.Mutex

func writeFile() {
	lock.Lock()
	defer lock.Unlock()

	file, err := os.OpenFile("data.txt", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("New Data\n")
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func main() {
	go writeFile()
	go writeFile()
}
```
- `sync.Mutex` prevents race conditions when multiple Goroutines write to the same file.

---

## **5.6 Summary: Why Use `os`?**

| Feature | Function |
|---------|----------|
| **File Handling** | `os.Create`, `os.WriteFile`, `os.ReadFile`, `os.Remove` |
| **Directory Handling** | `os.Mkdir`, `os.ReadDir`, `os.RemoveAll` |
| **Environment Variables** | `os.Getenv`, `os.Setenv`, `os.Environ` |
| **Process Management** | `os.Getpid`, `exec.Command`, `os.Exit` |
| **Concurrency Control** | `sync.Mutex` (from `os/sync`) |

---
# **6️⃣ Deep Dive into `sync` Package**

The `sync` package in Go provides primitives for synchronization, allowing safe concurrent programming. It includes:
- **Mutexes** (`sync.Mutex`, `sync.RWMutex`)
- **WaitGroups** (`sync.WaitGroup`)
- **Once** (`sync.Once`)
- **Condition Variables** (`sync.Cond`)
- **Map** (`sync.Map`)
- **Pools** (`sync.Pool`)

---

## **6.1 Mutex: Preventing Race Conditions**

### **6.1.1 `sync.Mutex` (Mutual Exclusion Lock)**

A `Mutex` ensures that only **one Goroutine** can access a critical section at a time.

#### **Example: Without Mutex (Race Condition)**

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var counter = 0

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		counter++ // Multiple Goroutines modifying the counter at the same time (race condition)
	}
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increment(&wg)
	}
	wg.Wait()
	fmt.Println("Final Counter Value:", counter) // Unreliable output due to race condition
}
```
🔴 **Issue**: The counter variable is accessed by multiple Goroutines **without synchronization**, leading to inconsistent results.

---

#### **Example: With `sync.Mutex` (Safe Access)**

```go
var mu sync.Mutex

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		mu.Lock()
		counter++
		mu.Unlock()
	}
}
```
✅ **Fix**: `mu.Lock()` ensures **only one Goroutine** modifies `counter` at a time.

---

### **6.1.2 `sync.RWMutex` (Read/Write Mutex)**

- `sync.RWMutex` allows **multiple readers** but only **one writer** at a time.
- Useful for scenarios where **reading is more frequent than writing**.

#### **Example: Read vs Write Lock**

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var rw sync.RWMutex
var data = 0

func reader(id int) {
	rw.RLock()
	fmt.Printf("Reader %d reads data: %d\n", id, data)
	time.Sleep(100 * time.Millisecond)
	rw.RUnlock()
}

func writer(id int) {
	rw.Lock()
	data++
	fmt.Printf("Writer %d updates data to: %d\n", id, data)
	time.Sleep(200 * time.Millisecond)
	rw.Unlock()
}

func main() {
	for i := 0; i < 3; i++ {
		go reader(i)
	}
	go writer(1)

	time.Sleep(1 * time.Second)
}
```
✅ **Benefit**: Allows multiple readers **without blocking**, but **only one writer** at a time.

---

## **6.2 `sync.WaitGroup`: Waiting for Goroutines**

`sync.WaitGroup` ensures the **main Goroutine waits** until all spawned Goroutines complete execution.

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d is working\n", id)
	time.Sleep(1 * time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait() // Wait for all workers to finish
	fmt.Println("All workers done")
}
```
✅ **Key Points**:
- `wg.Add(n)`: Adds `n` Goroutines to wait for.
- `wg.Done()`: Marks one Goroutine as completed.
- `wg.Wait()`: Blocks until all Goroutines call `Done()`.

---

## **6.3 `sync.Once`: Ensuring One-Time Execution**

`sync.Once` ensures a function runs **only once**, even if called from multiple Goroutines.

```go
package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func initialize() {
	fmt.Println("Initializing system...")
}

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	once.Do(initialize) // Will execute only once
	fmt.Println("Worker running")
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
}
```
✅ **Use Case**: Useful for **one-time database connection setup, configuration loading, etc.**

---

## **6.4 `sync.Cond`: Signaling Between Goroutines**

A `sync.Cond` is used for **Goroutine coordination**, where one Goroutine signals another to proceed.

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var ready = false
var cond = sync.NewCond(&sync.Mutex{})

func waiter() {
	cond.L.Lock()
	for !ready {
		cond.Wait() // Wait until signaled
	}
	fmt.Println("Goroutine received signal!")
	cond.L.Unlock()
}

func signaler() {
	time.Sleep(2 * time.Second)
	cond.L.Lock()
	ready = true
	cond.Signal() // Send signal to one waiting Goroutine
	cond.L.Unlock()
}

func main() {
	go waiter()
	go signaler()
	time.Sleep(3 * time.Second)
}
```
✅ **Use Case**: Useful when **one Goroutine must wait** for another to complete some work.

---

## **6.5 `sync.Map`: Concurrent Safe Map**

`sync.Map` is a **thread-safe alternative** to Go’s built-in `map`.

### **6.5.1 Writing and Reading from `sync.Map`**

```go
package main

import (
	"fmt"
	"sync"
)

var m sync.Map

func main() {
	m.Store("name", "Go")
	m.Store("age", 10)

	val, ok := m.Load("name")
	if ok {
		fmt.Println("Name:", val) // Name: Go
	}

	m.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":", value)
		return true
	})
}
```
✅ **Benefit**: **No need for manual locking** when reading/writing to `sync.Map`.

---

## **6.6 `sync.Pool`: Object Caching**

`sync.Pool` is used for **reusing objects**, reducing memory allocations.

### **6.6.1 Using `sync.Pool` for Reusing Objects**

```go
package main

import (
	"fmt"
	"sync"
)

var pool = sync.Pool{
	New: func() interface{} {
		return "New Object"
	},
}

func main() {
	obj := pool.Get()
	fmt.Println(obj) // New Object

	pool.Put("Reused Object")

	obj = pool.Get()
	fmt.Println(obj) // Reused Object
}
```
✅ **Use Case**: **Optimizing memory** by reusing objects instead of reallocating them.

---

## **6.7 Summary: Why Use `sync`?**

| Feature | Function |
|---------|----------|
| **Mutex** | `sync.Mutex`, `sync.RWMutex` (prevents race conditions) |
| **Waiting** | `sync.WaitGroup` (waits for Goroutines to finish) |
| **One-time Execution** | `sync.Once` (ensures function runs only once) |
| **Condition Variables** | `sync.Cond` (signals between Goroutines) |
| **Concurrent Map** | `sync.Map` (thread-safe alternative to `map`) |
| **Object Pooling** | `sync.Pool` (reuses objects to optimize memory) |

---

