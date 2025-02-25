

# Go Composite Types: The Deep Dive

---

## 1. Arrays
**Definition**: Fixed-length, contiguous sequence of elements of a single type.

#### Memory Allocation
- **Where**: **Stack** (usually), unless it “escapes” to the heap.
- **Why**: Fixed size is known at compile time, so the compiler can allocate it inline. If an array escapes (e.g., returned from a function), it goes to the heap.
- **Size**: `len * sizeof(element)` bytes.

#### Default Passing Mechanism
- **By Value**: Arrays are copied when passed to functions—every element gets duplicated.
- **Impact**: Expensive for big arrays (memory and CPU hit).

#### Internal Representation
```
Array: [3]int{1, 2, 3}
Memory (Stack):
[ 1 | 2 | 3 ]  // 3 * 8 bytes = 24 bytes (assuming 64-bit int)
```

#### Example
```go
package main

import "fmt"

func processArray(arr [3]int) {
    arr[0] = 99 // Copy modified, original untouched
}

func main() {
    data := [3]int{1, 2, 3}
    processArray(data)
    fmt.Println(data) // [1 2 3]—original unchanged
}
```

#### Performance Impact
- **Pros**: Predictable memory, no runtime resizing.
- **Cons**: Copying large arrays kills performance—O(n) time and space.

---

## 2. Slices

**Definition**: Dynamic, flexible view over an underlying array.

#### Memory Allocation
- **Where**: **Heap** (underlying array) + **Stack** (slice header).
- **Why**: The slice header (`struct { ptr, len, cap }`) is small and lives on the stack, but the backing array is heap-allocated for dynamic resizing.
- **Size**: Header = 24 bytes (64-bit: 8 for pointer, 8 for `len`, 8 for `cap`) + backing array size.

#### Default Passing Mechanism
- **By Value**: The slice header is copied when passed, but it points to the same backing array.
- **Impact**: Copying is cheap (24 bytes), but mutations to the array affect the original—reference-like behavior with a twist.

#### Internal Representation
```
Slice: []int{1, 2, 3}
Slice Header (Stack):
[ ptr -> | len: 3 | cap: 3 ]
Backing Array (Heap):
[ 1 | 2 | 3 ]
```

#### Slice Internals: Capacity & Resizing
- **Structure**: `struct { ptr *T; len int; cap int }`.
- **Growth**: When `append` exceeds `cap`, Go allocates a new, larger array (often 2x), copies the data, and updates the pointer—O(n) cost.
- **Example**: `make([]int, 2, 4)`—room for 2 elements, capacity for 4.

#### Your Code: Step-by-Step Breakdown
Here’s your code integrated, followed by what’s happening under the hood:

```go
package main

import "fmt"

func modifySlice(s []int) []int {
    fmt.Println("\nInside modifySlice function:")
    fmt.Println("Before modification:", s)
    s[0] = 500         // Modify first element
    s = append(s, 999) // Append a new element
    fmt.Println("After modification:", s)
    return s
}

func main() {
    // Step 1: Create an array and slice it
    arr := [5]int{10, 20, 30, 40, 50}
    slice := arr[1:4] // slice -> {20, 30, 40}

    fmt.Println("Initial Array:", arr)
    fmt.Println("Initial Slice:", slice)

    // Step 2: Pass slice to function
    slice = modifySlice(slice)

    fmt.Println("\nAfter function call:")
    fmt.Println("Array:", arr)
    fmt.Println("Slice:", slice)
}
```

**Output**:
```
Initial Array: [10 20 30 40 50]
Initial Slice: [20 30 40]

Inside modifySlice function:
Before modification: [20 30 40]
After modification: [500 30 40 999]

After function call:
Array: [10 500 30 40 50]
Slice: [500 30 40 999]
```

#### Under-the-Hood Explanation
Let’s dissect this with memory states and mechanics:

1. **Step 1: Array Creation (`arr`)**
    - `arr := [5]int{10, 20, 30, 40, 50}` creates a fixed 5-element array.
    - **Memory (Stack)**:
      ```
      arr: [ 10 | 20 | 30 | 40 | 50 ] // 5 * 8 bytes = 40 bytes
      ```
    - Fixed size, stack-allocated (unless it escapes).

2. **Step 2: Slice Creation (`slice := arr[1:4]`)**
    - `slice` is a view of `arr` from index 1 to 3 (exclusive 4): `{20, 30, 40}`.
    - **Slice Header (Stack)**:
      ```
      slice: [ ptr -> arr[1] | len: 3 | cap: 4 ]
      ```
    - **Backing Array (Still Stack)**:
      ```
      arr: [ 10 | 20 | 30 | 40 | 50 ]
             ^-- ptr points here
      ```
    - `len = 3` (elements visible), `cap = 4` (remaining space from index 1 to end).

3. **Step 3: Passing to `modifySlice(slice)`**
    - Slice header is copied by value to `s`—new stack frame, same backing array.
    - **Main Stack**: `slice: [ ptr -> arr[1] | len: 3 | cap: 4 ]`
    - **Function Stack**: `s: [ ptr -> arr[1] | len: 3 | cap: 4 ]`
    - **Memory**: No new allocation—both point to `arr`.

4. **Step 4: Modify First Element (`s[0] = 500`)**
    - `s[0]` accesses the first element of the slice (offset 1 in `arr`).
    - Updates `arr[1]` from 20 to 500.
    - **Backing Array (Stack)**:
      ```
      arr: [ 10 | 500 | 30 | 40 | 50 ]
      ```
    - `s` and `slice` still share this array—mutation is visible everywhere.

5. **Step 5: Append (`s = append(s, 999)`)**
    - Current `len = 3`, `cap = 4`—room for one more element without resizing.
    - `append` adds 999 at index 3 (offset 4 in `arr`), updates `len` to 4.
    - **Backing Array (Stack)**:
      ```
      arr: [ 10 | 500 | 30 | 40 | 999 ]
      ```
    - **Updated `s` Header (Function Stack)**:
      ```
      s: [ ptr -> arr[1] | len: 4 | cap: 4 ]
      ```
    - No new allocation—fits within `cap`.

6. **Step 6: Return and Assign (`slice = modifySlice(slice)`)**
    - `s` (updated header) is copied back to `slice` in `main`.
    - **Main Stack**:
      ```
      slice: [ ptr -> arr[1] | len: 4 | cap: 4 ]
      ```
    - Final `slice` is `[500, 30, 40, 999]`, and `arr` reflects the changes.

7. **What If `cap` Was Exceeded?**
    - If `slice` had `cap = 3` (e.g., `arr[1:4:4]` with max capacity set), `append` would:
        - Allocate a new heap array (e.g., 6 elements).
        - Copy `{500, 30, 40}` to it.
        - Add 999.
        - **New Heap Array**: `[500, 30, 40, 999, ?, ?]`.
        - `s` would point there, leaving `arr` as `[10, 500, 30, 40, 50]`.

#### Key Takeaways
- **Mutation**: `s[0] = 500` modifies the shared backing array—visible in `arr` and `slice`.
- **Append**: Stays in-place if `cap` allows; otherwise, new heap allocation.
- **Passing**: Header copy is cheap, but array sharing means side effects.

#### Performance Impact
- **Pros**: Flexible resizing, cheap header passing.
- **Cons**: Over-capacity appends trigger heap allocations—plan `cap` to avoid this.

#### Optimization Tip
- Use `make([]int, len, cap)` or `arr[start:end:cap]` to pre-set capacity:
```go
slice := arr[1:4:5] // Explicit cap = 5
```

---


## 3. Maps

Maps in Go are dynamic, unordered key-value stores—think dictionaries or hash maps. They’re perfect for fast lookups, like mapping device names to CPU usage in your monitoring system. Let’s unpack everything under the hood.

---

#### Internal Implementation
- **Structure**: Go maps are hash tables with buckets—arrays of slots holding key-value pairs.
- **Core Type**: Internally, a map is a pointer to a `runtime.hmap` struct:
  ```go
  type hmap struct {
      count     int       // Number of entries (len)
      B         uint8     // Log2 of bucket count (e.g., B=3 -> 8 buckets)
      buckets   unsafe.Pointer // Array of buckets
      oldbuckets unsafe.Pointer // Old buckets during resize
      // Plus more fields for resizing, overflow, etc.
  }
  ```
- **Bucket**: Each bucket (~8 key-value pairs) is a `bmap` struct, with arrays for keys, values, and overflow pointers.

#### Memory Allocation
- **Where**: **Heap**.
- **Why**: Dynamic size—grows/shrinks at runtime, managed by the garbage collector (GC).
- **Size**: Starts small (e.g., 8 buckets), scales with entries. Each bucket + overhead ≈ dozens of bytes.

#### Hashing Mechanism
- **Hash Function**: Go uses a custom hash (e.g., `siphash` for strings) based on key type—fast and collision-resistant.
- **Bucket Selection**: `hash(key) & (bucketCount-1)` picks a bucket (bitwise AND for power-of-2 sizing).
- **Collisions**: Handled with chaining—overflow buckets linked to the main bucket.

#### Bucket Structure
```
Map: map[string]int{"a": 1, "b": 2}
Header (Stack): [ ptr -> hmap ]
Heap (hmap):
[ count: 2 | B: 1 | buckets -> ]
Buckets:
[ bucket0: ("a", 1) | bucket1: ("b", 2) ]
```

---

### Declaring & Initializing Maps
Maps need explicit creation—here’s every way to do it:

1. **Using `make` (Preferred for Preallocation)**:
   ```go
   m := make(map[string]int)        // Empty, no preallocation
   m2 := make(map[string]int, 10)   // Hint: 10 entries—fewer resizes
   ```
    - **When**: General use, or when you know the size upfront.

2. **Map Literal (Convenient for Small Maps)**:
   ```go
   m := map[string]int{
       "a": 1,
       "b": 2, // Trailing comma required
   }
   ```
    - **When**: Quick setup with initial data.

3. **Var Declaration (Nil Map)**:
   ```go
   var m map[string]int // nil—can’t assign yet!
   ```
    - **When**: Placeholder—initialize later with `make`.

- **Under the Hood**: `make` allocates the `hmap` + initial buckets on the heap. Literals call `make` + populate.

---

### Insertions, Deletions, Lookups
- **Insert**: `m[key] = value`
    - Hash key → find bucket → store key-value pair.
    - **Perf**: O(1) average, O(n) worst-case with collisions.
- **Delete**: `delete(m, key)`
    - Hash key → mark slot as free (tombstone or shift).
    - **Perf**: O(1) average.
- **Lookup**: `value, ok := m[key]`
    - Hash key → check bucket → return value or zero if missing.
    - **Perf**: O(1) average.

#### Resizing
- **Trigger**: Load factor (~6.5 entries/bucket) exceeded or too many overflow buckets.
- **Process**: Doubles buckets (e.g., 8 → 16), rehashes keys—O(n) operation.
- **Memory**:
  ```
  Before: [ b0: 6 pairs | b1: 5 pairs ]
  After:  [ b0 | b1 | b2 | b3 ] // Rehashed
  ```

---

### Passing Maps to Functions
- **Mechanism**: **By Value**—copies the `hmap` pointer (8 bytes).
- **Behavior**: Reference-like—mutations affect the original map.
- **Impact**: Cheap copy, shared data.

#### Example
```go
package main

import "fmt"

func modifyMap(m map[string]int) {
    m["a"] = 99 // Modifies original
    delete(m, "b")
}

func main() {
    m := map[string]int{"a": 1, "b": 2}
    modifyMap(m)
    fmt.Println(m) // map[a:99]
}
```

#### Memory Diagram
```
Before Call:
Stack (main): [ m: ptr -> ]
Heap: [ hmap: count=2 | buckets -> [ ("a", 1) | ("b", 2) ] ]

Function Stack:
Stack (modifyMap): [ m: ptr -> ] // Same hmap
Heap: [ hmap: count=1 | buckets -> [ ("a", 99) ] ]
```

---

### Advanced Topics
1. **Struct Keys**:
    - Must be **comparable** (no slices, maps, funcs).
    - **Example**:
      ```go
      type Key struct {
          ID   int
          Name string
      }
      m := map[Key]int{{1, "bro"}: 42}
      ```

2. **Custom Hash Functions**: Not directly supported—Go’s runtime picks the hash. Use structs with comparable fields.

3. **Efficient Iteration**:
    - `for k, v := range m`—order is randomized (since Go 1.0).
    - **Tip**: Preallocate slices for keys/values if collecting them.

---

### Built-in Functions
- **make(map[K]V, [hint])**: Creates a map, optional size hint.
- **delete(m, key)**: Removes a key-value pair.
- **len(m)**: Returns number of entries.
- **range**: Iterates over key-value pairs.

---

## Go 1.21+ `maps` Package
From Go 1.21, the `maps` package adds goodies (import `"maps"`):


The `maps` package provides **utility functions** to manipulate maps efficiently. These functions are available from **Go 1.21+**.

### **1. Clearing All Keys from a Map (`Clear`)**
This is a built-in method to **remove all keys from a map**.

#### **Syntax:**
```go
func Clear[M ~map[K]V, K comparable, V any](m M)
```
#### **Example:**
```go
package main

import (
	"fmt"
	"maps"
)

func main() {
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("Before clearing:", m)

	maps.Clear(m) // Removes all keys

	fmt.Println("After clearing:", m) // Output: map[]
}
```
✅ **More efficient** than manually deleting keys.

---

### **2. Cloning a Map (`Clone`)**
Creates an **exact copy** of a map.

#### **Syntax:**
```go
func Clone[M ~map[K]V, K comparable, V any](m M) M
```
#### **Example:**
```go
m1 := map[string]int{"apple": 10, "banana": 20}
m2 := maps.Clone(m1) // Deep copy

fmt.Println(m1) // Output: map[apple:10 banana:20]
fmt.Println(m2) // Output: map[apple:10 banana:20]
```
✅ `m2` is a separate copy of `m1`, modifying one won’t affect the other.

---

### **3. Copying One Map to Another (`Copy`)**
Copies **all key-value pairs** from one map to another.

#### **Syntax:**
```go
func Copy[M1 ~map[K]V, M2 ~map[K]V, K comparable, V any](dst M1, src M2)
```
#### **Example:**
```go
m1 := map[string]int{"a": 1}
m2 := map[string]int{"b": 2, "c": 3}

maps.Copy(m1, m2) // Copies all key-value pairs from m2 to m1

fmt.Println(m1) // Output: map[a:1 b:2 c:3]
```
✅ `dst` map retains its original keys while new keys from `src` are added.

---

### **4. Deleting Keys with a Condition (`DeleteFunc`)**
Deletes keys **based on a condition function**.

#### **Syntax:**
```go
func DeleteFunc[M ~map[K]V, K comparable, V any](m M, del func(K, V) bool)
```
#### **Example:**
```go
m := map[string]int{"a": 5, "b": 10, "c": 15}

// Remove entries where value > 7
maps.DeleteFunc(m, func(k string, v int) bool {
	return v > 7
})

fmt.Println(m) // Output: map[a:5]
```
✅ **Useful for conditional deletions**.

---

### **5. Comparing Two Maps for Equality (`Equal`)**
Checks if **two maps are identical**.

#### **Syntax:**
```go
func Equal[M1, M2 ~map[K]V, K, V comparable](m1 M1, m2 M2) bool
```
#### **Example:**
```go
m1 := map[string]int{"x": 1, "y": 2}
m2 := map[string]int{"x": 1, "y": 2}

fmt.Println(maps.Equal(m1, m2)) // Output: true
```
✅ **Returns `true` if both maps contain the same key-value pairs.**

---

### **6. Comparing Maps with Custom Logic (`EqualFunc`)**
Compares **values** using a **custom equality function**.

#### **Syntax:**
```go
func EqualFunc[M1 ~map[K]V1, M2 ~map[K]V2, K comparable, V1, V2 any](m1 M1, m2 M2, eq func(V1, V2) bool) bool
```
#### **Example:**
```go
m1 := map[string]string{"a": "hello"}
m2 := map[string]string{"a": "HELLO"}

// Case-insensitive comparison
isEqual := maps.EqualFunc(m1, m2, func(v1, v2 string) bool {
	return strings.EqualFold(v1, v2)
})

fmt.Println(isEqual) // Output: true
```
✅ Allows **custom comparison logic**.

---

### **7. Extracting All Keys (`Keys`)**
Returns a **slice of all map keys**.

#### **Syntax:**
```go
func Keys[M ~map[K]V, K comparable, V any](m M) []K
```
#### **Example:**
```go
m := map[string]int{"one": 1, "two": 2, "three": 3}
keys := maps.Keys(m)

fmt.Println(keys) // Output: [one two three] (order is random)
```
✅ **Useful for iteration & sorting.**

---

### **8. Extracting All Values (`Values`)**
Returns a **slice of all values** from a map.

#### **Syntax:**
```go
func Values[M ~map[K]V, K comparable, V any](m M) []V
```
#### **Example:**
```go
m := map[string]int{"one": 1, "two": 2, "three": 3}
values := maps.Values(m)

fmt.Println(values) // Output: [1 2 3] (order is random)
```
✅ Extracts only values, making it **easy to process**.

---

## **🔥 Key Takeaways:**
- `maps.Clear(m)`: **Efficiently removes all keys**.
- `maps.Clone(m)`: **Creates a deep copy**.
- `maps.Copy(dst, src)`: **Merges two maps**.
- `maps.DeleteFunc(m, condition)`: **Deletes based on a function**.
- `maps.Equal(m1, m2)`: **Checks if maps are identical**.
- `maps.EqualFunc(m1, m2, customFunc)`: **Custom equality check**.
- `maps.Keys(m)`: **Extracts all keys**.
- `maps.Values(m)`: **Extracts all values**.

---

### Pitfalls & Edge Cases
1. **Nil Maps**:
    - Can’t assign to a `nil` map—panics or silent failure.
   ```go
   var m map[string]int
   // m["a"] = 1 // Panic!
   m = make(map[string]int) // Fix it
   ```

2. **Non-Comparable Keys**:
    - Slices, maps, funcs as keys = compile error.
   ```go
   // map[[]int]int{} // Nope!
   ```

3. **Concurrency Issues**:
    - Maps aren’t thread-safe—concurrent writes cause race conditions.
    - **Fix**: Use `sync.Map` or mutexes:
      ```go
      var m = make(map[string]int)
      var mu sync.Mutex
      mu.Lock()
      m["a"] = 1
      mu.Unlock()
      ```

---

### Performance Trade-Offs & Optimizations
- **Insert/Lookup**: O(1) average, but collisions degrade to O(n).
- **Resizing**: O(n)—preallocate with `make(map[K]V, size)`:
  ```go
  m := make(map[string]int, 1000) // Fewer resizes
  ```
- **Iteration**: O(n), but random order adds CPU overhead.
- **Benchmark Example**:
  ```go
  package main

  import "testing"

  func BenchmarkMapInsert(b *testing.B) {
      m := make(map[int]int, b.N)
      for i := 0; i < b.N; i++ {
          m[i] = i
      }
  }
  ```

---


### 4. Structs
**Definition**: Custom type bundling fields.

#### Memory Allocation
- **Where**: **Stack** (if local and doesn’t escape) or **Heap** (if returned or stored globally).
- **Why**: Fixed size, but escape analysis decides placement.
- **Size**: Sum of fields + padding for alignment (more on this below).

#### Default Passing Mechanism
- **By Value**: Entire struct is copied—fields and all.
- **Impact**: Small structs are fine; big ones tank performance.

#### Internal Representation
```
Struct: struct { ID int; Name string }{42, "bro"}
Memory (Stack):
[ ID: 42 (8 bytes) | Name.ptr -> "bro" (8) | Name.len: 3 (8) ]
Heap: "bro" (string data)
```

#### Example
```go
package main

import "fmt"

type Monitor struct {
    CPU    float64
    Memory uint64
}

func processStruct(m Monitor) {
    m.CPU = 99.9 // Copy modified, original safe
}

func main() {
    data := Monitor{CPU: 75.5, Memory: 1234}
    processStruct(data)
    fmt.Println(data.CPU) // 75.5—unchanged
}
```

#### Struct Memory Alignment
- **Rule**: Fields align to their size (e.g., `int64` on 8-byte boundaries).
- **Example**:
```go
type Misaligned struct {
    a int8  // 1 byte
    b int64 // 8 bytes
    c int8  // 1 byte
}
// Size: 24 bytes (not 10!) due to padding:
// [ a: 1 | pad: 7 | b: 8 | c: 1 | pad: 7 ]
```
- **Optimized**:
```go
type Aligned struct {
    b int64 // 8 bytes
    a int8  // 1 byte
    c int8  // 1 byte
}
// Size: 16 bytes—less padding:
// [ b: 8 | a: 1 | c: 1 | pad: 6 ]
```

---

### 5. Channels
**Definition**: Thread-safe queue for goroutine communication.

#### Memory Allocation
- **Where**: **Heap**.
- **Why**: Dynamic, shared resource—needs runtime management.
- **Size**: Varies—includes buffer + metadata.

#### Default Passing Mechanism
- **By Value**: Copies the channel pointer (8 bytes).
- **Impact**: Cheap, acts like a reference to the same channel.

#### Internal Representation
```
Channel: make(chan int, 2)
Header (Stack):
[ ptr -> ]
Channel Struct (Heap):
[ buffer: [1, 2] | lock | waiters | ... ]
```

#### Example
```go
package main

import "fmt"

func processChan(ch chan int) {
    ch <- 99 // Sends to original channel
}

func main() {
    ch := make(chan int, 1)
    ch <- 42
    processChan(ch)
    fmt.Println(<-ch) // 42 (FIFO—99 queued)
}
```

#### Performance Impact
- **Pros**: Lightweight coordination.
- **Cons**: Buffered channels allocate heap memory—unbuffered are leaner.

---

### Function Calls & Memory Allocation
Here’s a step-by-step breakdown of how Go handles memory in function calls:

#### Example: Passing a Slice
```go
func process(data []int) []int {
    data[0] = 99       // Modifies backing array
    data = append(data, 100) // Might reallocate
    return data
}

func main() {
    s := []int{1, 2, 3}
    s = process(s)
}
```
1. **Stack Setup**: `main`’s stack frame gets `s` (24-byte header).
2. **Call `process`**: New stack frame, `data` copies `s`’s header (ptr, len, cap).
3. **Mutation**: `data[0] = 99` hits the heap array—shared via pointer.
4. **Append**: If `cap` exceeded, new heap array allocated, old data copied, `data` updated.
5. **Return**: `data` copied back to `s`—stack-to-stack, but heap array persists.
6. **GC**: Old array (if resized) becomes garbage—collected later.

#### Edge Case: Escape Analysis
```go
func escape() []int {
    x := []int{1, 2, 3}
    return x // Escapes to heap!
}
```
- **Before**: `x` could be stack-allocated.
- **After**: Compiler sees return—moves backing array to heap.

---

### Performance & Memory: Value vs. Reference
- **By Value (Arrays, Structs)**:
    - **Memory**: Full copy—stack or heap grows.
    - **Perf**: Slow for big data (O(n) copy time).
- **By Reference-ish (Slices, Maps, Channels)**:
    - **Memory**: Copies pointer—tiny footprint.
    - **Perf**: Fast, but shared data means mutation risks.

#### Optimization: Use Pointers
```go
func processPtr(m *Monitor) {
    m.CPU = 99.9 // Modifies original
}

data := Monitor{CPU: 75.5}
processPtr(&data) // 8-byte pointer, no copy
```

---

### Advanced Topics
1. **Slice Capacity**:
    - Preallocate with `make([]T, len, cap)` to avoid resizing.
    - **Tip**: Overestimate `cap` for append-heavy code.

2. **Struct Alignment**:
    - Order fields largest-to-smallest to minimize padding.
    - **Tool**: `unsafe.Sizeof()` to check.

3. **Garbage Collection**:
    - Heap allocations (slices, maps) trigger GC—minimize by reusing memory.
    - **Tip**: Pool objects (e.g., `sync.Pool`) for frequent allocations.

---

### Best Practices
- **Arrays**: Use sparingly—slices are more practical.
- **Slices**: Set `cap` upfront; avoid over-appending.
- **Maps**: Preallocate with `make(map[K]V, size)` for bulk inserts.
- **Structs**: Pass pointers for big ones; align fields.
- **Channels**: Use unbuffered unless buffering boosts perf.

---


## **Pass-by-Value vs. Pass-by-Reference in Go**
| Type        | Passed By   | Explanation |
|------------|------------|-------------|
| **Arrays** | Value      | A copy of the entire array is made. Changes inside the function **do not affect** the original. |
| **Slices** | Reference  | A slice is a **pointer** to an underlying array, so modifications **affect the original array**. |
| **Maps**   | Reference  | Maps are internally implemented as pointers, so changes **affect the original map**. |
| **Structs** | Value      | A copy of the struct is passed. Changes inside a function **do not affect the original** unless using pointers. |
| **Channels** | Reference | Channels are references, so passing them to a function means **both have access to the same channel**. |

---

## **Deep Dive: Behavior of Each Type**

### **1️⃣ Arrays – Passed by Value (Copy)**
```go
package main
import "fmt"

func modifyArray(arr [3]int) {
    arr[0] = 100
}

func main() {
    arr := [3]int{1, 2, 3}
    modifyArray(arr)  // Passes a copy of the array
    fmt.Println(arr)  // Output: [1 2 3] (Original remains unchanged)
}
```
💡 **Key Takeaway:** Arrays are copied when passed to a function.

---

### **2️⃣ Slices – Passed by Reference**
Slices behave like **references** because they contain:
1. A **pointer** to an underlying array.
2. A **length**.
3. A **capacity**.

```go
package main
import "fmt"

func modifySlice(s []int) {
    s[0] = 100  // Modifies original slice
}

func main() {
    s := []int{1, 2, 3}
    modifySlice(s)  // The function modifies the original slice
    fmt.Println(s)  // Output: [100 2 3] (Original is changed)
}
```
💡 **Key Takeaway:** Modifying a slice inside a function affects the **original slice**.

---

### **3️⃣ Maps – Passed by Reference**
Maps in Go **always use references**, so when you pass a map to a function, it **modifies the original map**.

```go
package main
import "fmt"

func modifyMap(m map[string]int) {
    m["Alice"] = 100  // Modifies original map
}

func main() {
    m := map[string]int{"Alice": 25, "Bob": 30}
    modifyMap(m)  // Changes the original map
    fmt.Println(m)  // Output: map[Alice:100 Bob:30]
}
```
💡 **Key Takeaway:** Maps are **always passed by reference**.

---

### **4️⃣ Structs – Passed by Value**
By default, **structs are copied** when passed to a function.

```go
package main
import "fmt"

type Person struct {
    Name string
    Age  int
}

func modifyStruct(p Person) {
    p.Age = 100  // Changes only the copy
}

func main() {
    p := Person{Name: "John", Age: 25}
    modifyStruct(p)  // Passes a copy of the struct
    fmt.Println(p)   // Output: {John 25} (Original is unchanged)
}
```
💡 **Key Takeaway:** Structs are **copied** when passed to functions.

#### **Using Pointers to Modify Structs**
To **modify the original struct**, pass a **pointer** instead.
```go
func modifyStruct(p *Person) {
    p.Age = 100  // Modifies the original struct
}
```
💡 **Best Practice:** Use `*Person` when you need to **modify** struct fields inside functions.

---

### **5️⃣ Channels – Passed by Reference**
Channels are **always passed by reference**.
```go
package main
import "fmt"

func sendMessage(ch chan string) {
    ch <- "Hello"  // Sends message through the channel
}

func main() {
    ch := make(chan string, 1)
    sendMessage(ch)
    fmt.Println(<-ch)  // Output: Hello
}
```
💡 **Key Takeaway:** Since channels are references, both **sender and receiver** share the same channel.

---

## **Final Summary**
| Type      | Default Behavior | How to Modify Original? |
|-----------|----------------|-------------------------|
| **Array**  | Passed by Value | Use a pointer `*[N]int` |
| **Slice**  | Passed by Reference | No need for pointers |
| **Map**    | Passed by Reference | No need for pointers |
| **Struct** | Passed by Value | Use a pointer `*StructName` |
| **Channel** | Passed by Reference | No need for pointers |

---

