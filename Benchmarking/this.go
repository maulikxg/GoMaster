package main

import (
	"fmt"
	"hash/fnv"
	"strconv"
	"strings"
	"time"
)

// Hash function to convert string to uint
func hashString(s string) uint {
	h := fnv.New32a() // Using 32-bit FNV hash
	h.Write([]byte(s))
	return uint(h.Sum32())
}

func prepareKeys(nKeys int, repeatCount int) []string {
	var keys []string
	for i := 0; i < nKeys; i++ {
		s := strings.Repeat("#"+strconv.Itoa(i), repeatCount)
		keys = append(keys, s)
	}
	return keys
}

func main() {

	// Initialize maps and arrays
	stringMap := make(map[string]int)
	uintMap := make(map[uint]int)
	stringKeys := []string{}
	uintKeys := []uint{}

	// Insert sample keys

	keys := prepareKeys(100000, 10)
	//	keys := []string{"1#2#3#4#5#6#7#9#10#5#6#7#8#1#2#3#4#5#6#7#9#10#", "5#6#7#8#1#2#3#4#5#6#7#9#10#5#6#7#8#1#2#3#4#5#6#7#9#10#", "9#10#11#12#1#2#3#4#5#6#7#9#10#", "1#2#3#4#5#6#7#9#10#13#14#15#16#", "5#6#7#8#1#2#3#4#5#6#7#9#10#5#6#7#8#1#2#3#4#5#6#7#9#10#", "1#2#3#4#5#6#7#9#10#13#14#15#16#", "5#6#7#8#1#2#3#4#5#6#7#9#10#5#6#7#8#1#2#3#4#5#6#7#9#10#", "1#2#3#4#5#6#7#9#10#13#14#15#16#", "5#6#7#8#1#2#3#4#5#6#7#9#10#5#6#7#8#1#2#3#4#5#6#7#9#10#", "1#2#3#4#5#6#7#9#10#13#14#15#16#", "5#6#7#8#1#2#3#4#5#6#7#9#10#5#6#7#8#1#2#3#4#5#6#7#9#10#"}
	for _, key := range keys {
		length := len(key)
		stringMap[key] = length
		stringKeys = append(stringKeys, key)

		hash := hashString(key)
		uintMap[hash] = length
		uintKeys = append(uintKeys, hash)
	}
	// warmup for cpu cache
	for _, key := range stringKeys {
		_ = stringMap[key] // Access value
	}

	// Measure iteration time for stringKeys
	startString := time.Now()
	for _, key := range stringKeys {
		_ = stringMap[key] // Access value
	}
	elapsedString := time.Since(startString)

	for _, key := range uintKeys {
		_ = uintMap[key] // Access value
	}

	// Measure iteration time for uintKeys
	startUint := time.Now()
	for _, key := range uintKeys {
		_ = uintMap[key] // Access value
	}
	elapsedUint := time.Since(startUint)

	// Print results
	fmt.Println("Time to iterate string keys:", elapsedString)
	fmt.Println("Time to iterate uint keys:", elapsedUint)
	fmt.Println("Time Diff:", elapsedString-elapsedUint)

}
