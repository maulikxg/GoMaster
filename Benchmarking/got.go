package main

import (
	"fmt"
	"hash/fnv"
	"strconv"
	"strings"
	"time"
)

const (
	nKeys       = 100000 // number of keys to insert into each map
	lookupLoops = 100    // number of times to loop through the keys for benchmarking
)

// generateStringKey returns a key in the format "i###i###i###i###i###i###i###i"
func generateStringKey(i int) string {
	parts := make([]string, 8)
	s := strconv.Itoa(i)
	for j := 0; j < 8; j++ {
		parts[j] = s
	}
	return strings.Join(parts, "###")
}

// hashString returns a uint64 hash for a given string using FNV-1a.
func hashString(s string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return h.Sum64()
}

// benchmarkStringMap benchmarks lookups on map[string]int.
func benchmarkStringMap() {
	// Build the string map.
	stringMap := make(map[string]int, nKeys)
	stringKeys := make([]string, 0, nKeys)
	for i := 0; i < nKeys; i++ {
		key := generateStringKey(i)
		stringMap[key] = i
		stringKeys = append(stringKeys, key)
	}

	// Warm-up lookups.
	for _, key := range stringKeys {
		_ = stringMap[key]
	}

	start := time.Now()
	// Do many lookups over the keys.
	var sum int
	for i := 0; i < lookupLoops; i++ {
		for _, key := range stringKeys {
			sum += stringMap[key]
		}
	}
	duration := time.Since(start)

	fmt.Printf("map[string]int: Total lookup time: %v (average per lookup: %v), checksum: %d\n",
		duration, duration/time.Duration(lookupLoops*nKeys), sum)
}

// benchmarkUint64Map benchmarks lookups on map[uint64]int.
func benchmarkUint64Map() {
	// Build the uint64 map.
	uint64Map := make(map[uint64]int, nKeys)
	uint64Keys := make([]uint64, 0, nKeys)
	for i := 0; i < nKeys; i++ {
		strKey := generateStringKey(i)
		hashKey := hashString(strKey)
		uint64Map[hashKey] = i
		uint64Keys = append(uint64Keys, hashKey)
	}

	// Warm-up lookups.
	for _, key := range uint64Keys {
		_ = uint64Map[key]
	}

	start := time.Now()
	// Do many lookups over the keys.
	var sum int
	for i := 0; i < lookupLoops; i++ {
		for _, key := range uint64Keys {
			sum += uint64Map[key]
		}
	}
	duration := time.Since(start)

	fmt.Printf("map[uint64]int: Total lookup time: %v (average per lookup: %v), checksum: %d\n",
		duration, duration/time.Duration(lookupLoops*nKeys), sum)
}

func main() {
	fmt.Println("Benchmarking map[string]int")
	benchmarkStringMap()
	fmt.Println("Benchmarking map[uint64]int")
	benchmarkUint64Map()
}
