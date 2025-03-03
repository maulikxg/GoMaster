
## **1. What is a Closure in Go?**
A **closure** is a function that captures variables from its surrounding scope even after the outer function has finished execution.

### **Example: Basic Closure**
```go
package main

import "fmt"

func outer() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	counter := outer()
	fmt.Println(counter()) // Output: 1
	fmt.Println(counter()) // Output: 2
	fmt.Println(counter()) // Output: 3
}
```
- Here, the anonymous function inside `outer` remembers the `count` variable, even after `outer()` has executed.
- Each time `counter()` is called, it increments `count`.

---

## **2. How Closures Work in Memory**
Closures store the references to the variables in their lexical scope, rather than copying them.

### **Example: Multiple Closures Using the Same Variable**
```go
package main

import "fmt"

func main() {
	funcs := []func(){}
	for i := 0; i < 3; i++ {
		funcs = append(funcs, func() {
			fmt.Println(i) // Captures `i` from loop scope
		})
	}

	for _, f := range funcs {
		f()
	}
}
```
### **Output (Unexpected)**
```
3
3
3
```
### **Explanation**
- The closure captures the reference to `i`, not its value at that point.
- When the closure executes, `i` has already become `3`.

### **Fix: Pass `i` as a Function Argument**
```go
package main

import "fmt"

func main() {
	funcs := []func(){}

	for i := 0; i < 3; i++ {
		funcs = append(funcs, func(n int) func() {
			return func() {
				fmt.Println(n) // Captures the value of `i`
			}
		}(i))
	}

	for _, f := range funcs {
		f()
	}
}
```
### **Correct Output**
```
0
1
2
```

---

## **3. Types of Closures**
Closures in Go can take different forms:

### **3.1 Anonymous Function Closure**
An anonymous function that directly captures variables.
```go
package main

import "fmt"

func main() {
	message := "Hello"
	changeMessage := func() {
		message = "Hi"
	}
	changeMessage()
	fmt.Println(message) // Output: Hi
}
```
- The function `changeMessage` modifies `message` because it holds a reference to it.

---

### **3.2 Closure as a Function Return**
Functions returning closures.
```go
package main

import "fmt"

func multiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func main() {
	double := multiplier(2)
	triple := multiplier(3)

	fmt.Println(double(4))  // Output: 8
	fmt.Println(triple(4))  // Output: 12
}
```
- `double` captures `factor = 2`, while `triple` captures `factor = 3`.

---

### **3.3 Closures with Goroutines**
Closures inside goroutines can lead to unexpected behavior if they capture loop variables.

#### **Incorrect Closure in Goroutine**
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 3; i++ {
		go func() {
			fmt.Println(i) // Captures reference to `i`
		}()
	}
	time.Sleep(time.Second)
}
```
#### **Output (Unpredictable)**
```
3
3
3
```
- All goroutines print `3` because they capture the same reference.

#### **Fix: Pass `i` as an Argument**
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 3; i++ {
		go func(n int) {
			fmt.Println(n) // Captures the value at loop iteration
		}(i)
	}
	time.Sleep(time.Second)
}
```
#### **Output**
```
0
1
2
```

---

## **4. Closures with Struct Methods**
Closures can be used within methods to capture struct fields.

```go
package main

import "fmt"

type Counter struct {
	count int
}

func (c *Counter) Incrementer() func() int {
	return func() int {
		c.count++
		return c.count
	}
}

func main() {
	c := Counter{}
	incr := c.Incrementer()

	fmt.Println(incr()) // Output: 1
	fmt.Println(incr()) // Output: 2
}
```
- The closure remembers `c.count`, allowing it to persist across calls.

---

## **5. Closures with Closures (Nested Closures)**
Functions returning other closures.

```go
package main

import "fmt"

func outer() func() func() int {
	count := 0
	return func() func() int {
		return func() int {
			count++
			return count
		}
	}
}

func main() {
	inner := outer()() // Calling both levels of functions
	fmt.Println(inner()) // Output: 1
	fmt.Println(inner()) // Output: 2
}
```

---

## **6. Practical Use Cases**
### **6.1 Memoization using Closures**
Closures help store state for caching expensive function calls.

```go
package main

import "fmt"

func memoizedAdder() func(int) int {
	cache := make(map[int]int)
	return func(n int) int {
		if _, exists := cache[n]; !exists {
			cache[n] = n + 10
		}
		return cache[n]
	}
}

func main() {
	add := memoizedAdder()

	fmt.Println(add(5))  // Computed: 15
	fmt.Println(add(5))  // Cached: 15
	fmt.Println(add(7))  // Computed: 17
}
```

---

### **6.2 Event Handlers in Web Applications**
In web servers, closures are useful for request handling.
```go
package main

import (
	"fmt"
	"net/http"
)

func handler(name string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %s!", name)
	}
}

func main() {
	http.HandleFunc("/", handler("Maulik"))
	http.ListenAndServe(":8080", nil)
}
```

---

## **7. Key Takeaways**
1. **Closures capture variables** from their lexical scope, not values.
2. Be careful with **loop variables**, as they capture references.
3. Use **function arguments** to capture correct values inside loops.
4. Closures are **useful for memoization, goroutines, and web handlers**.
5. **Nested closures** allow for layered function calls.
6. Closures are heavily used in **functional programming and concurrency**.

---

## **8. Edge Cases & Best Practices**
### **Avoid Modifying Global State Unintentionally**
```go
package main

import "fmt"

func globalCounter() func() int {
	var count int
	return func() int {
		count++
		return count
	}
}

func main() {
	counter1 := globalCounter()
	counter2 := globalCounter()

	fmt.Println(counter1()) // 1
	fmt.Println(counter1()) // 2
	fmt.Println(counter2()) // 1
}
```
- Each closure gets its own `count`, preventing conflicts.

