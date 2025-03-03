# Go Testing: A Beginner's Guide

This guide will teach you about testing in Go, from basic concepts to advanced techniques. We'll start with simple examples and gradually move to more complex topics.

## 1. Basic Testing in Go

Go has a built-in testing framework in the `testing` package that makes it easy to write and run tests.

### Getting Started with Go Tests

Tests in Go are written as functions in files with names ending in `_test.go`. These files should be in the same package as the code being tested.

Here's a simple example:

```go
// main.go - This is our main code file
package main

// Sum adds two numbers together
func Sum(a, b int) int {
    return a + b
}
```

```go
// sum_test.go - This is our test file
package main

import "testing"

func TestSum(t *testing.T) {
    // Step 1: Set up our test values
    a, b := 2, 3
    expected := 5
    
    // Step 2: Call the function we want to test
    result := Sum(a, b)
    
    // Step 3: Check if the result matches what we expect
    if result != expected {
        t.Errorf("Sum(%d, %d) = %d; expected %d", a, b, result, expected)
    }
}
```

Notice that:
- Test functions must start with `Test` followed by a capitalized name
- Test functions take a parameter `t *testing.T`
- We use `t.Errorf()` to report test failures

### Running Tests

To run tests, use the `go test` command in your terminal:

```bash
go test                  # Run tests in current directory
go test ./...            # Run tests in current directory and all subdirectories
go test -v               # Verbose output (shows all test results, not just failures)
go test -run TestSum     # Run only tests matching "TestSum"
```

## 2. Table-Driven Tests

Table-driven tests are a common pattern in Go that allows you to test multiple scenarios with minimal code. This is especially useful when you want to test a function with different inputs.

```go
// main.go
package main

// Multiply returns the product of two integers
func Multiply(a, b int) int {
    return a * b
}
```

```go
// multiply_test.go
package main

import "testing"

func TestMultiply(t *testing.T) {
    // Define a list of test cases
    testCases := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive numbers", 2, 3, 6},
        {"zero multiplier", 5, 0, 0},
        {"negative multiplier", 3, -2, -6},
        {"both negative", -2, -3, 6},
    }
    
    // Loop through all test cases
    for _, tc := range testCases {
        // Use t.Run to create a subtest for each case
        t.Run(tc.name, func(t *testing.T) {
            result := Multiply(tc.a, tc.b)
            if result != tc.expected {
                t.Errorf("Multiply(%d, %d) = %d; expected %d", 
                         tc.a, tc.b, result, tc.expected)
            }
        })
    }
}
```

Benefits of table-driven tests:
- Easy to add new test cases
- Each test case has a clear description
- You can run specific test cases using the `-run` flag
- Code is more maintainable since you avoid duplicating test logic

## 3. Test Fixtures and Helper Functions

When testing more complex code, you'll often need to set up test environments (fixtures) and create helper functions.

```go
// user.go
package user

import "errors"

// User represents a person in our system
type User struct {
    ID   int
    Name string
    Age  int
}

// Repository stores and retrieves users
type Repository struct {
    users map[int]User
}

// NewRepository creates a new user repository
func NewRepository() *Repository {
    return &Repository{
        users: make(map[int]User),
    }
}

// Create adds a new user to the repository
func (r *Repository) Create(u User) error {
    if u.Name == "" {
        return errors.New("user name cannot be empty")
    }
    if u.Age < 0 {
        return errors.New("user age cannot be negative")
    }
    r.users[u.ID] = u
    return nil
}

// Get retrieves a user by ID
func (r *Repository) Get(id int) (User, error) {
    user, exists := r.users[id]
    if !exists {
        return User{}, errors.New("user not found")
    }
    return user, nil
}
```

```go
// user_test.go
package user

import (
    "testing"
)

// setupTest is a helper function to set up a test environment
func setupTest() *Repository {
    return NewRepository()
}

// teardownTest cleans up after a test
func teardownTest(r *Repository) {
    // In a real application, this might close database connections, etc.
}

func TestUserRepository(t *testing.T) {
    t.Run("create valid user", func(t *testing.T) {
        // Setup - create a fresh repository for testing
        repo := setupTest()
        defer teardownTest(repo)
        
        // Test creating a user
        user := User{ID: 1, Name: "John", Age: 30}
        err := repo.Create(user)
        
        // Check that no error was returned
        if err != nil {
            t.Errorf("Expected no error, got %v", err)
        }
        
        // Verify user exists in the repository
        savedUser, err := repo.Get(1)
        if err != nil {
            t.Errorf("Expected no error getting user, got %v", err)
        }
        if savedUser.Name != "John" {
            t.Errorf("Expected name John, got %s", savedUser.Name)
        }
    })
    
    t.Run("create invalid user", func(t *testing.T) {
        // Setup
        repo := setupTest()
        defer teardownTest(repo)
        
        // Test - try to create user with empty name
        user := User{ID: 2, Name: "", Age: 30}
        err := repo.Create(user)
        
        // Assert - should get an error
        if err == nil {
            t.Error("Expected error for empty name, got none")
        }
        
        // Test - try to create user with negative age
        user = User{ID: 3, Name: "Jane", Age: -1}
        err = repo.Create(user)
        
        // Assert - should get an error
        if err == nil {
            t.Error("Expected error for negative age, got none")
        }
    })
}
```

Key points:
- `setupTest()` creates a clean testing environment for each test
- `teardownTest()` cleans up resources after each test
- `defer teardownTest(repo)` ensures cleanup happens even if the test fails
- Each subtest focuses on a specific behavior

## 4. Test Coverage

Test coverage measures how much of your code is executed during tests. Go has built-in tools for measuring coverage.

```go
// string_utils.go
package utils

import "strings"

// StringUtils provides string manipulation functions
type StringUtils struct{}

// Reverse returns the reverse of a string
func (s *StringUtils) Reverse(input string) string {
    runes := []rune(input)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

// Capitalize capitalizes the first letter of each word
func (s *StringUtils) Capitalize(input string) string {
    words := strings.Fields(input)
    for i, word := range words {
        if len(word) > 0 {
            words[i] = strings.ToUpper(word[:1]) + word[1:]
        }
    }
    return strings.Join(words, " ")
}

// Count counts occurrences of a substring in a string
func (s *StringUtils) Count(input, substr string) int {
    return strings.Count(input, substr)
}
```

```go
// string_utils_test.go
package utils

import "testing"

func TestStringUtils(t *testing.T) {
    utils := &StringUtils{}
    
    t.Run("Reverse", func(t *testing.T) {
        testCases := []struct {
            input    string
            expected string
        }{
            {"hello", "olleh"},
            {"golang", "gnalog"},
            {"", ""},
            {"a", "a"},
        }
        
        for _, tc := range testCases {
            result := utils.Reverse(tc.input)
            if result != tc.expected {
                t.Errorf("Reverse(%q) = %q; expected %q", 
                         tc.input, result, tc.expected)
            }
        }
    })
    
    t.Run("Capitalize", func(t *testing.T) {
        testCases := []struct {
            input    string
            expected string
        }{
            {"hello world", "Hello World"},
            {"go testing", "Go Testing"},
            {"", ""},
        }
        
        for _, tc := range testCases {
            result := utils.Capitalize(tc.input)
            if result != tc.expected {
                t.Errorf("Capitalize(%q) = %q; expected %q", 
                         tc.input, result, tc.expected)
            }
        }
    })
    
    // Notice that we didn't write a test for the Count function!
    // The coverage report will show this missing coverage
}
```

### Running Coverage

```bash
# Generate coverage profile
go test -coverprofile=coverage.out

# View coverage summary in terminal
go tool cover -func=coverage.out

# Generate HTML report (opens in your browser)
go tool cover -html=coverage.out
```

The HTML report will highlight which lines of code were executed during tests (in green) and which weren't (in red).

## 5. Benchmarking in Go

Benchmarking helps you measure and compare the performance of your code. This is useful when you have multiple implementations of the same functionality.

```go
// sort.go
package sort

// BubbleSort implements bubble sort algorithm for integers
func BubbleSort(items []int) {
    n := len(items)
    for i := 0; i < n; i++ {
        for j := 0; j < n-i-1; j++ {
            if items[j] > items[j+1] {
                items[j], items[j+1] = items[j+1], items[j]
            }
        }
    }
}

// QuickSort implements quick sort algorithm for integers
func QuickSort(items []int) {
    if len(items) <= 1 {
        return
    }
    
    pivot := items[0]
    var left, right []int
    
    for _, item := range items[1:] {
        if item <= pivot {
            left = append(left, item)
        } else {
            right = append(right, item)
        }
    }
    
    QuickSort(left)
    QuickSort(right)
    
    copy(items, append(append(left, pivot), right...))
}
```

```go
// sort_test.go
package sort

import (
    "fmt"
    "math/rand"
    "sort"
    "testing"
)

// generateSlice creates a random slice of integers for testing
func generateSlice(size int) []int {
    slice := make([]int, size)
    for i := 0; i < size; i++ {
        slice[i] = rand.Intn(999) - 499 // Random numbers between -499 and 499
    }
    return slice
}

// isSorted checks if a slice is sorted
func isSorted(items []int) bool {
    for i := 1; i < len(items); i++ {
        if items[i] < items[i-1] {
            return false
        }
    }
    return true
}

// Tests to verify correctness
func TestSorting(t *testing.T) {
    t.Run("BubbleSort", func(t *testing.T) {
        items := []int{5, 1, 4, 2, 8}
        expected := []int{1, 2, 4, 5, 8}
        
        BubbleSort(items)
        
        for i, v := range items {
            if v != expected[i] {
                t.Errorf("BubbleSort: items[%d] = %d; expected %d", 
                         i, v, expected[i])
            }
        }
    })
    
    t.Run("QuickSort", func(t *testing.T) {
        items := []int{5, 1, 4, 2, 8}
        expected := []int{1, 2, 4, 5, 8}
        
        QuickSort(items)
        
        for i, v := range items {
            if v != expected[i] {
                t.Errorf("QuickSort: items[%d] = %d; expected %d", 
                         i, v, expected[i])
            }
        }
    })
}

// Benchmarks to compare performance
func BenchmarkBubbleSort(b *testing.B) {
    // Run the benchmark for different slice sizes
    for _, size := range []int{10, 100, 1000} {
        b.Run(fmt.Sprintf("size-%d", size), func(b *testing.B) {
            // Reset the timer to not include setup time
            b.ResetTimer()
            
            for i := 0; i < b.N; i++ {
                // Generate a new slice for each iteration
                b.StopTimer()
                items := generateSlice(size)
                b.StartTimer()
                
                // Run the sort
                BubbleSort(items)
            }
        })
    }
}

func BenchmarkQuickSort(b *testing.B) {
    for _, size := range []int{10, 100, 1000} {
        b.Run(fmt.Sprintf("size-%d", size), func(b *testing.B) {
            b.ResetTimer()
            
            for i := 0; i < b.N; i++ {
                b.StopTimer()
                items := generateSlice(size)
                b.StartTimer()
                
                QuickSort(items)
            }
        })
    }
}

func BenchmarkStandardSort(b *testing.B) {
    for _, size := range []int{10, 100, 1000} {
        b.Run(fmt.Sprintf("size-%d", size), func(b *testing.B) {
            b.ResetTimer()
            
            for i := 0; i < b.N; i++ {
                b.StopTimer()
                items := generateSlice(size)
                b.StartTimer()
                
                sort.Ints(items)
            }
        })
    }
}
```

### Running Benchmarks

```bash
# Run all benchmarks
go test -bench=.

# Run specific benchmark
go test -bench=BubbleSort

# Run benchmarks with memory allocation statistics
go test -bench=. -benchmem

# Run benchmarks for longer to get more stable results
go test -bench=. -benchtime=5s
```

Key points about benchmarks:
- Benchmark functions must start with `Benchmark` and take a parameter `b *testing.B`
- `b.N` is controlled by the testing framework and represents the number of iterations
- Use `b.StopTimer()` and `b.StartTimer()` to exclude setup time from measurements
- The `-bench` flag specifies which benchmarks to run (use `.` for all)

## 6. Test Mocking

Mocking is essential for isolating components during testing. This is particularly useful when testing code that interacts with databases, external APIs, or other services.

```go
// user_service.go
package user

import "context"

// User represents a user in the system
type User struct {
    ID    int
    Name  string
    Email string
}

// Repository defines the interface for user data operations
type Repository interface {
    Get(ctx context.Context, id int) (*User, error)
    Save(ctx context.Context, user *User) error
}

// NotificationService defines the interface for sending notifications
type NotificationService interface {
    SendWelcomeEmail(ctx context.Context, email, name string) error
}

// Service provides user-related operations
type Service struct {
    repo      Repository
    notifier  NotificationService
}

// NewService creates a new user service
func NewService(repo Repository, notifier NotificationService) *Service {
    return &Service{
        repo:     repo,
        notifier: notifier,
    }
}

// CreateUser creates a new user and sends a welcome email
func (s *Service) CreateUser(ctx context.Context, name, email string) (*User, error) {
    // Create user object
    user := &User{
        Name:  name,
        Email: email,
    }
    
    // Save user to repository
    if err := s.repo.Save(ctx, user); err != nil {
        return nil, err
    }
    
    // Send welcome email
    if err := s.notifier.SendWelcomeEmail(ctx, email, name); err != nil {
        return nil, err
    }
    
    return user, nil
}
```

```go
// user_service_test.go
package user

import (
    "context"
    "errors"
    "testing"
)

// MockRepository is a fake implementation of Repository for testing
type MockRepository struct {
    users map[int]*User
    SaveFunc func(ctx context.Context, user *User) error
}

func NewMockRepository() *MockRepository {
    return &MockRepository{
        users: make(map[int]*User),
    }
}

func (m *MockRepository) Get(ctx context.Context, id int) (*User, error) {
    user, exists := m.users[id]
    if !exists {
        return nil, errors.New("user not found")
    }
    return user, nil
}

func (m *MockRepository) Save(ctx context.Context, user *User) error {
    // If SaveFunc is defined, use that implementation
    if m.SaveFunc != nil {
        return m.SaveFunc(ctx, user)
    }
    
    // Default implementation
    if user.ID == 0 {
        user.ID = len(m.users) + 1
    }
    
    m.users[user.ID] = user
    return nil
}

// MockNotificationService is a fake implementation of NotificationService
type MockNotificationService struct {
    SendWelcomeEmailFunc func(ctx context.Context, email, name string) error
    EmailsSent           map[string]bool
}

func NewMockNotificationService() *MockNotificationService {
    return &MockNotificationService{
        EmailsSent: make(map[string]bool),
    }
}

func (m *MockNotificationService) SendWelcomeEmail(ctx context.Context, email, name string) error {
    if m.SendWelcomeEmailFunc != nil {
        return m.SendWelcomeEmailFunc(ctx, email, name)
    }
    
    m.EmailsSent[email] = true
    return nil
}

func TestCreateUser(t *testing.T) {
    t.Run("successful creation", func(t *testing.T) {
        // Setup mocks
        repo := NewMockRepository()
        notifier := NewMockNotificationService()
        
        // Create service with mocks
        service := NewService(repo, notifier)
        
        // Test user creation
        ctx := context.Background()
        user, err := service.CreateUser(ctx, "John Doe", "john@example.com")
        
        // Assertions
        if err != nil {
            t.Fatalf("Expected no error, got %v", err)
        }
        if user == nil {
            t.Fatal("Expected user, got nil")
        }
        if user.Name != "John Doe" {
            t.Errorf("Expected name 'John Doe', got '%s'", user.Name)
        }
        if user.Email != "john@example.com" {
            t.Errorf("Expected email 'john@example.com', got '%s'", user.Email)
        }
        if user.ID == 0 {
            t.Error("Expected non-zero ID")
        }
        
        // Verify email was sent
        if !notifier.EmailsSent["john@example.com"] {
            t.Error("Expected welcome email to be sent")
        }
    })
    
    t.Run("repository error", func(t *testing.T) {
        // Setup mocks
        repo := NewMockRepository()
        notifier := NewMockNotificationService()
        
        // Make the repository return an error
        repo.SaveFunc = func(ctx context.Context, user *User) error {
            return errors.New("database error")
        }
        
        // Create service with mocks
        service := NewService(repo, notifier)
        
        // Test user creation
        ctx := context.Background()
        user, err := service.CreateUser(ctx, "John Doe", "john@example.com")
        
        // Assertions
        if err == nil {
            t.Fatal("Expected error, got nil")
        }
        if user != nil {
            t.Errorf("Expected nil user, got %+v", user)
        }
        
        // Verify no email was sent
        if notifier.EmailsSent["john@example.com"] {
            t.Error("Expected no welcome email to be sent")
        }
    })
}
```

Key points about mocking:
- Create mock implementations of interfaces to replace real dependencies
- Use custom functions to control behavior (like simulating errors)
- Verify interactions with dependencies (like checking if an email was sent)
- This allows testing components in isolation

## 7. Advanced: Fuzz Testing

Fuzz testing automatically generates random inputs to find bugs and edge cases in your code.

```go
// parser.go
package parser

import (
    "encoding/json"
    "errors"
    "strings"
)

// Config represents application configuration
type Config struct {
    Name    string `json:"name"`
    Version string `json:"version"`
    Debug   bool   `json:"debug"`
    Timeout int    `json:"timeout"`
}

// ParseJSON parses a JSON string into a Config struct
func ParseJSON(input string) (Config, error) {
    var config Config
    
    // Trim whitespace
    input = strings.TrimSpace(input)
    
    // Check for empty input
    if input == "" {
        return config, errors.New("empty input")
    }
    
    // Parse JSON
    err := json.Unmarshal([]byte(input), &config)
    if err != nil {
        return config, err
    }
    
    // Validation
    if config.Name == "" {
        return config, errors.New("name is required")
    }
    
    if config.Timeout < 0 {
        return config, errors.New("timeout cannot be negative")
    }
    
    return config, nil
}
```

```go
// parser_test.go
package parser

import (
    "strings"
    "testing"
)

func TestParseJSON(t *testing.T) {
    testCases := []struct {
        name        string
        input       string
        expectError bool
        expected    Config
    }{
        {
            name:        "valid config",
            input:       `{"name":"app","version":"1.0","debug":true,"timeout":30}`,
            expectError: false,
            expected:    Config{Name: "app", Version: "1.0", Debug: true, Timeout: 30},
        },
        {
            name:        "empty input",
            input:       "",
            expectError: true,
        },
        {
            name:        "invalid json",
            input:       `{"name":"app", version:"1.0"}`,
            expectError: true,
        },
        {
            name:        "missing name",
            input:       `{"version":"1.0","debug":true,"timeout":30}`,
            expectError: true,
        },
        {
            name:        "negative timeout",
            input:       `{"name":"app","version":"1.0","debug":true,"timeout":-10}`,
            expectError: true,
        },
    }
    
    // parser_test.go (continued)
    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            config, err := ParseJSON(tc.input)
            
            // Check if we got an error when we expected one
            if tc.expectError && err == nil {
                t.Error("Expected error but got none")
            }
            
            // Check if we got an unexpected error
            if !tc.expectError && err != nil {
                t.Errorf("Expected no error but got: %v", err)
            }
            
            // If we didn't expect an error, verify the parsed config
            if !tc.expectError {
                if config.Name != tc.expected.Name {
                    t.Errorf("Expected name %q but got %q", tc.expected.Name, config.Name)
                }
                if config.Version != tc.expected.Version {
                    t.Errorf("Expected version %q but got %q", tc.expected.Version, config.Version)
                }
                if config.Debug != tc.expected.Debug {
                    t.Errorf("Expected debug %v but got %v", tc.expected.Debug, config.Debug)
                }
                if config.Timeout != tc.expected.Timeout {
                    t.Errorf("Expected timeout %d but got %d", tc.expected.Timeout, config.Timeout)
                }
            }
        })
    }
}

// FuzzParseJSON is a fuzz test for the ParseJSON function
func FuzzParseJSON(f *testing.F) {
    // Add seed corpus (inputs that are known to be interesting)
    seeds := []string{
        `{"name":"app","version":"1.0","debug":true,"timeout":30}`,
        `{"name":"app"}`,
        `{"name":"app","timeout":0}`,
    }
    
    for _, seed := range seeds {
        f.Add(seed)
    }
    
    // Fuzz test function
    f.Fuzz(func(t *testing.T, input string) {
        config, err := ParseJSON(input)
        
        // If parsing succeeds, verify invariants
        if err == nil {
            // Name should never be empty (as enforced by our validation)
            if config.Name == "" {
                t.Errorf("Successfully parsed config has empty name: %v", config)
            }
            
            // Timeout should never be negative
            if config.Timeout < 0 {
                t.Errorf("Successfully parsed config has negative timeout: %v", config)
            }
        }
    })
}

```
