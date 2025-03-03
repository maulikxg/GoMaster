

# 1. Basic Unit Testing
## **Writing and Running a Simple Test**
📌 **Create a file `mathutils.go`:**
```go
package mathutils

// Add returns the sum of two numbers
func Add(a, b int) int {
    return a + b
}
```

📌 **Create a test file `mathutils_test.go`:**
```go
package mathutils

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5

    if result != expected {
        t.Errorf("Add(2, 3) = %d; want %d", result, expected)
    }
}
```

📌 **Run the test:**
```sh
go test
```

### **✅ Expected Output:**
```
ok      mathutils        0.002s
```
🔍 **Explanation:**
- `ok` → The test **passed successfully**.
- `mathutils` → Name of the package being tested.
- `0.002s` → Time taken to run the test.

📌 **For a failing test:**  
Modify the test to expect `6` instead of `5`:
```go
expected := 6
```
Run `go test` again:
```
--- FAIL: TestAdd (0.00s)
    mathutils_test.go:10: Add(2, 3) = 5; want 6
FAIL
exit status 1
FAIL    mathutils    0.001s
```
🔍 **Explanation:**
- `--- FAIL: TestAdd` → The test failed.
- `mathutils_test.go:10` → Error happened at line 10.
- `Add(2, 3) = 5; want 6` → The function returned `5`, but the test expected `6`.

---

# 2. Table-Driven Testing
### **Refactoring the Test for Multiple Cases**
```go
func TestAddTableDriven(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"Positive numbers", 2, 3, 5},
        {"Negative numbers", -2, -3, -5},
        {"Zero values", 0, 0, 0},
    }

    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            result := Add(tc.a, tc.b)
            if result != tc.expected {
                t.Errorf("Add(%d, %d) = %d; want %d", tc.a, tc.b, result, tc.expected)
            }
        })
    }
}
```

📌 **Run the test with verbose output:**
```sh
go test -v
```

### **✅ Expected Output:**
```
=== RUN   TestAddTableDriven
=== RUN   TestAddTableDriven/Positive_numbers
=== RUN   TestAddTableDriven/Negative_numbers
=== RUN   TestAddTableDriven/Zero_values
--- PASS: TestAddTableDriven (0.00s)
    --- PASS: TestAddTableDriven/Positive_numbers (0.00s)
    --- PASS: TestAddTableDriven/Negative_numbers (0.00s)
    --- PASS: TestAddTableDriven/Zero_values (0.00s)
PASS
ok      mathutils    0.002s
```
🔍 **Explanation:**
- `=== RUN   TestAddTableDriven` → The **main test function** is running.
- `=== RUN   TestAddTableDriven/Positive_numbers` → Running **sub-test** for positive numbers.
- `PASS` → All tests passed.

📌 **If a test fails:**  
Modify `expected` value for one case (e.g., `"Zero values", 0, 0, 1`).  
Run `go test -v` again:
```
=== RUN   TestAddTableDriven
=== RUN   TestAddTableDriven/Positive_numbers
=== RUN   TestAddTableDriven/Negative_numbers
=== RUN   TestAddTableDriven/Zero_values
    mathutils_test.go:17: Add(0, 0) = 0; want 1
--- FAIL: TestAddTableDriven (0.00s)
    --- PASS: TestAddTableDriven/Positive_numbers (0.00s)
    --- PASS: TestAddTableDriven/Negative_numbers (0.00s)
    --- FAIL: TestAddTableDriven/Zero_values (0.00s)
FAIL
exit status 1
FAIL    mathutils    0.002s
```
🔍 **Explanation:**
- `FAIL` appears only for `"Zero values"`.
- The function returned `0`, but the test expected `1`.

---

# 3. Benchmarking
## **Measuring Performance of `Add()`**
```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(100, 200)
    }
}
```

📌 **Run the benchmark:**
```sh
go test -bench=.
```

### **✅ Expected Output:**
```
BenchmarkAdd-8      1000000000      0.312 ns/op
```
🔍 **Explanation:**
- `BenchmarkAdd-8` → Running on **8 CPU threads**.
- `1000000000` → Function ran **1 billion times**.
- `0.312 ns/op` → Each operation took **0.312 nanoseconds**.

---

# 4. Test Coverage & Profiling
## **Measuring Code Coverage**
```sh
go test -cover
```

### **✅ Expected Output:**
```
coverage: 100.0% of statements
```
🔍 **Explanation:**
- **100%** → All lines of code were tested.

## **Generating Coverage Report**
```sh
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```
🔍 **Explanation:**
- Opens an HTML report showing **which lines are covered**.

---

# 5️. Fuzz Testing
## **Finding Unexpected Edge Cases**
```go
func FuzzAdd(f *testing.F) {
    testcases := []struct {
        a, b int
    }{
        {1, 2}, {0, 0}, {-1, -2},
    }

    for _, tc := range testcases {
        f.Add(tc.a, tc.b)
    }

    f.Fuzz(func(t *testing.T, a, b int) {
        result := Add(a, b)
        if result < a || result < b {
            t.Errorf("Unexpected result: Add(%d, %d) = %d", a, b, result)
        }
    })
}
```

📌 **Run the fuzz test:**
```sh
go test -fuzz=FuzzAdd
```

### **✅ Expected Output (Finding Edge Cases):**
```
fuzz: elapsed: 3s, execs: 50000 (16667/sec), new interesting: 3
FAIL
```
🔍 **Explanation:**
- `execs: 50000` → **50,000 test cases were generated automatically**.
- `new interesting: 3` → **3 unexpected behaviors** were found.

---

# **📌 Final Summary**
| Concept | Explanation | Expected Output |
|---------|------------|----------------|
| **Unit Testing** | Checks function correctness | `ok mathutils 0.002s` |
| **Table-Driven Tests** | Tests multiple cases efficiently | Sub-tests are displayed in verbose mode |
| **Benchmarking** | Measures performance | Shows operations per nanosecond |
| **Test Coverage** | Ensures all code is tested | `coverage: 100%` |
| **Fuzz Testing** | Finds edge cases | Shows how many inputs broke the function |

---
