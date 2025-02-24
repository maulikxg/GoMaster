

# Go Data Types: The Ultimate Guide

Lets Go!!

## Numeric Types: Integers
Whole numbers—signed (positive or negative) or unsigned (positive only).

| Type     | Size    | Encoding             | Range                       | Max Storage       | Default Value | Use Case                       | Pitfalls                        |
|----------|---------|----------------------|-----------------------------|-------------------|---------------|--------------------------------|---------------------------------|
| `int8`   | 8 bits  | Signed (1 byte)      | -128 to 127                 | 256 values (2⁸)   | `0`           | Small counters, flags          | Overflow at 127!                |
| `uint8`  | 8 bits  | Unsigned (1 byte)    | 0 to 255                    | 256 values (2⁸)   | `0`           | Bytes, small counts (aka `byte`)| No negatives—careful!          |
| `int16`  | 16 bits | Signed (2 bytes)     | -32,768 to 32,767           | 65,536 (2¹⁶)      | `0`           | Port numbers (signed)          | Limited range for big stuff     |
| `uint16` | 16 bits | Unsigned (2 bytes)   | 0 to 65,535                 | 65,536 (2¹⁶)      | `0`           | Port numbers, packet sizes     | Perfect for networking ports    |
| `int32`  | 32 bits | Signed (4 bytes)     | -2,147,483,648 to 2,147,483,647 | 4.3B (2³²) | `0`           | General IDs (aka `rune`)       | Overflows on huge counts        |
| `uint32` | 32 bits | Unsigned (4 bytes)   | 0 to 4,294,967,295          | 4.3B (2³²)        | `0`           | Large positive counters        | No negatives, obviously         |
| `int64`  | 64 bits | Signed (8 bytes)     | -9.22Q to 9.22Q             | 18.4Q (2⁶⁴)       | `0`           | Timestamps, massive counts     | Overkill for small stuff        |
| `uint64` | 64 bits | Unsigned (8 bytes)   | 0 to 18.44Q                 | 18.4Q (2⁶⁴)       | `0`           | Huge throughput totals         | Memory hog if overused          |
| `int`    | Varies  | Platform-dependent   | 32-bit or 64-bit            | Varies            | `0`           | Default “I don’t care” integer | Range depends on your machine   |
| `uint`   | Varies  | Platform-dependent   | 32-bit or 64-bit            | Varies            | `0`           | Default unsigned integer       | Same deal—check your system     |

- **Notes**: 
  - “Q” = quintillion (~10¹⁸). Exact ranges: -2ⁿ⁻¹ to 2ⁿ⁻¹-1 (signed), 0 to 2ⁿ-1 (unsigned).
  - `uint8` is aliased as `byte`, `int32` as `rune` (for UTF-8 chars).
- **Teaching Tip**: Demo `int8(130)` overflowing to `-126` (two’s complement wrap-around)—blows minds!

## Numeric Types: Floating-Point
For decimals—think CPU percentages or latency.

| Type      | Size    | Encoding            | Range                | Max Storage       | Default Value | Precision       | Use Case                | Pitfalls                     |
|-----------|---------|---------------------|----------------------|-------------------|---------------|-----------------|-------------------------|------------------------------|
| `float32` | 32 bits | IEEE-754 (4 bytes)  | ~1.18e-38 to 3.4e38  | Varies (IEEE)     | `0.0`         | ~6-7 digits     | CPU %, lightweight stats| Precision loss on big nums   |
| `float64` | 64 bits | IEEE-754 (8 bytes)  | ~2.23e-308 to 1.8e308| Varies (IEEE)     | `0.0`         | ~15-17 digits   | Latency, precise metrics| Double memory vs `float32`  |

- **Fun Fact**: `float32(1.23456789)` rounds to ~`1.234568`—precision’s a thing!
- **Teaching Trick**: Compare `float32` vs `float64` in a loop adding `0.1`. `float32` drifts sooner.

## Other Numeric Types
Complex numbers—rare but cool.

| Type        | Size     | Encoding                   | Range                       | Default Value | Use Case            |
|-------------|----------|----------------------------|-----------------------------|---------------|---------------------|
| `complex64` | 8 bytes  | 32-bit real + 32-bit imag  | Per component (float32)     | `(0+0i)`      | Complex math        |
| `complex128`| 16 bytes | 64-bit real + 64-bit imag  | Per component (float64)     | `(0+0i)`      | High-precision math |

- **When?**: Signal processing or physics—not your network monitor!

## Boolean
True or false—simple as that.

| Type   | Size   | Range       | Default Value | Use Case         |
|--------|--------|-------------|---------------|------------------|
| `bool` | 1 byte | `true`, `false` | `false`     | Status flags     |

- **Pro Tip**: Default `false` = “off” for safety in uninitialized vars.

## String
Text data, UTF-8 encoded.

| Type     | Size            | Encoding | Max Length         | Default Value | Use Case            |
|----------|-----------------|----------|--------------------|---------------|---------------------|
| `string` | Variable (bytes)| UTF-8    | ~2³² or 2⁶⁴ bytes  | `""` (empty)  | IPs, names, logs    |

- **Gotcha**: `len("café")` = 5 (bytes), not 4 (chars), due to UTF-8.
- **Teaching Hook**: Show multi-byte chars (e.g., emojis) to explain UTF-8.

## Composite Types (Quick Hits)
- **`Arrays/Slices`**: Lists (e.g., `[]float64` for CPU readings over time).
- **`Structs`**: Custom types (e.g., `struct { CPU float64; Memory uint64 }`).
- **`Maps`**: Key-value pairs (e.g., `map[string]float64` for device-to-CPU).

## Network Monitoring System: Which Type for What?
Tracking CPU %, memory, latency, etc.? Here’s your playbook:

| Metric            | Best Type   | Why                                    | Example Value     |
|-------------------|-------------|----------------------------------------|-------------------|
| CPU Usage         | `float32` or `float64` | Decimals (use `float64` for precision) | `75.5` (%)      |
| Memory Usage      | `uint64`    | Huge, positive totals                  | `4294123` (bytes) |
| Throughput        | `uint64`    | Massive byte counts                    | `1234567` (B/s)   |
| Latency           | `float64`   | Precise decimals                       | `12.34` (ms)      |
| Packet Count      | `uint32` or `uint64` | Big enough (`uint64` for long-term) | `10000`        |
| Port Number       | `uint16`    | Fits 0–65,535 perfectly                | `8080`            |
| Server Status     | `bool`      | Simple up/down                         | `true`            |
| Device Name       | `string`    | Flexible text                          | `"server1"`       |

- **Example Struct**:
```go
type MonitorData struct {
    CPU       float64
    Memory    uint64
    Latency   float64
    Packets   uint32
    Port      uint16
    IsUp      bool
    Device    string
}
var data MonitorData // Defaults: 0.0, 0, 0.0, 0, 0, false, ""
```

## Type Casting: Making Types Work Together
Go’s statically typed—no auto-magic. Cast explicitly!

### 1. Numeric Casting
```go
x := 42         // int
y := float64(x) // 42.0
z := int32(y)   // 42
```
- **Trap**: `int(3.7)` -> `3` (truncates decimals).

### 2. String to Numeric (use `strconv`)
```go
s := "123.45"
f, _ := strconv.ParseFloat(s, 64) // 123.45 (float64)
i, _ := strconv.Atoi("42")       // 42 (int)
```
- **Use Case**: Parse CPU % from logs.

### 3. Numeric to String
```go
n := 42
s := strconv.Itoa(n)        // "42"
f := 3.14
fs := fmt.Sprintf("%.2f", f) // "3.14"
```
- **Use Case**: Display stats in your UI.

### 4. Signed vs Unsigned
```go
i := -10
u := uint(i)      // Huge positive num (two’s complement flip!)
u2 := uint32(255)
i2 := int32(u2)   // 255
```
- **Watch Out**: Negative-to-unsigned casts can surprise you.

### 5. Real-World Example (Monitoring)
```go
cpuStr := "75.5"
cpuFloat, _ := strconv.ParseFloat(cpuStr, 64) // 75.5
cpuInt := int(cpuFloat)                       // 75
cpuDisplay := fmt.Sprintf("CPU: %.1f%%", cpuFloat) // "CPU: 75.5%"
```

## Teaching Goodies
- **Default Values**: Uninitialized vars get “zero values” (0, 0.0, false, "")—bug-proofing built in!
- **Memory Matters**: `int64` = 8x `int8`. Don’t waste space on tiny data.
- **Overflow Demo**:
```go
var x int8 = 127
x++ // -128 (wraps!)
fmt.Println(x)
```
- **Float Precision**: `float32(0.1 + 0.2)` ≠ `0.3` exactly—show it!
- **UTF-8 Fun**: `len("🌟")` = 4 bytes, not 1 char.


---
