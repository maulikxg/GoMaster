

# **1. The Empty Interface (`interface{}`)**
- In Go, an **empty interface** (`interface{}`) can hold a value of any type.
- This is useful when working with unknown types, generic functions, or handling dynamic data.

### **Example:**
```go
func PrintAnything(i interface{}) {
    fmt.Println(i)
}

func main() {
    PrintAnything(42)       // Output: 42
    PrintAnything("Hello")  // Output: Hello
    PrintAnything(3.14)     // Output: 3.14
}
```

### **Use Cases:**
- Storing different types in a single slice:
  ```go
  var values []interface{} = []interface{}{42, "text", 3.14}
  fmt.Println(values)
  ```
- Passing dynamic arguments to functions.
- JSON decoding (since JSON structures can contain different types).

---

# **2. Type Assertions**
- A **type assertion** is used to retrieve the dynamic type of an interface value.

### **Syntax:**
```go
value, ok := interfaceValue.(Type)
```

### **Example:**
```go
var i interface{} = "Hello"

s, ok := i.(string) // Asserting it’s a string
if ok {
    fmt.Println("String:", s) // Output: String: Hello
} else {
    fmt.Println("Not a string")
}
```

### **Panic Risk:**
- If a type assertion is incorrect and you **omit the `ok` check**, it results in a **runtime panic**.
  ```go
  var x interface{} = 42
  y := x.(string)  // Panic: interface conversion: int is not a string
  ```

---

# **3. Type Switches**
- Instead of asserting types one by one, **type switches** allow checking multiple possible types.

### **Example:**
```go
func CheckType(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Println("Integer:", v)
    case string:
        fmt.Println("String:", v)
    case float64:
        fmt.Println("Float:", v)
    default:
        fmt.Println("Unknown type")
    }
}

func main() {
    CheckType(42)        // Output: Integer: 42
    CheckType("Hello")   // Output: String: Hello
    CheckType(3.14)      // Output: Float: 3.14
}
```

### **Use Cases:**
- Useful for **handling different types dynamically**.
- Helps avoid repeated **type assertions**.

---

# **4. Stringers (`fmt.Stringer` Interface)**
- The `fmt.Stringer` interface lets you define **custom string representations** for types.

### **Example:**
```go
package main
import "fmt"

type Person struct {
    Name string
    Age  int
}

func (p Person) String() string {
    return fmt.Sprintf("Person(Name: %s, Age: %d)", p.Name, p.Age)
}

func main() {
    p := Person{"Alice", 30}
    fmt.Println(p) // Output: Person(Name: Alice, Age: 30)
}
```

### **Use Cases:**
- Custom formatting for logging or debugging.
- More readable `fmt.Println()` output.

---

# **5. Exercise: Stringers**
- Implement the `fmt.Stringer` interface for a custom struct.

**Example:**
```go
type City struct {
    Name  string
    State string
}

func (c City) String() string {
    return fmt.Sprintf("%s, %s", c.Name, c.State)
}

func main() {
    c := City{"Ahmedabad", "Gujarat"}
    fmt.Println(c) // Output: Ahmedabad, Gujarat
}
```

---

# **6. Errors (`error` Interface)**
- Errors in Go are handled using the `error` interface.

### **Example:**
```go
package main
import "errors"
import "fmt"

func Divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

func main() {
    result, err := Divide(10, 0)
    if err != nil {
        fmt.Println("Error:", err) // Output: Error: division by zero
    }
}
```

### **Custom Error Types:**
- You can define **custom error types** by implementing the `Error()` method.

```go
type MyError struct {
    Message string
}

func (e *MyError) Error() string {
    return e.Message
}

func main() {
    err := &MyError{"Something went wrong"}
    fmt.Println(err) // Output: Something went wrong
}
```

---

# **7. Exercise: Errors**
- Modify the `Divide` function to return a custom error type.

---

# **8. Readers (`io.Reader` Interface)**
- The `io.Reader` interface is used for **reading data streams**.

### **Example:**
```go
package main
import (
    "fmt"
    "strings"
)

func main() {
    r := strings.NewReader("Hello, Go!")

    buf := make([]byte, 5)
    n, _ := r.Read(buf)
    fmt.Println(string(buf[:n])) // Output: Hello
}
```

### **Use Cases:**
- Reading from files (`os.File` implements `io.Reader`).
- Network connections (`net.Conn` implements `io.Reader`).

---

# **9. Exercise: Readers**
- Implement a custom `io.Reader`.

---

# **10. Exercise: `rot13Reader`**
- A `rot13Reader` should apply ROT13 transformation while reading data.

### **Example Implementation:**
```go
package main
import (
    "fmt"
    "io"
    "os"
    "strings"
)

type rot13Reader struct {
    r io.Reader
}

func (r rot13Reader) Read(b []byte) (int, error) {
    n, err := r.r.Read(b)
    for i := 0; i < n; i++ {
        if (b[i] >= 'A' && b[i] <= 'M') || (b[i] >= 'a' && b[i] <= 'm') {
            b[i] += 13
        } else if (b[i] >= 'N' && b[i] <= 'Z') || (b[i] >= 'n' && b[i] <= 'z') {
            b[i] -= 13
        }
    }
    return n, err
}

func main() {
    s := strings.NewReader("Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, r)
}
```

---

# **11. Images (`image` Package)**
- The `image` package provides basic image manipulation functionality.

### **Example: Creating an Image**
```go
package main
import (
    "fmt"
    "image"
    "image/color"
)

func main() {
    img := image.NewRGBA(image.Rect(0, 0, 100, 100))
    img.Set(50, 50, color.RGBA{255, 0, 0, 255}) // Set a pixel to red
    fmt.Println(img.At(50, 50))
}
```
