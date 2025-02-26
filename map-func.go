package main

import (
	"fmt"
	"maps"
)

func main() {

	m := make(map[string]int)

	// some data in map
	m["key1"] = 10
	m["key2"] = 20

	// for the iretating over the map(random order) -> [maps.All]
	for key, val := range maps.All(m) {
		fmt.Println(key, val)
	}

	// creating the shallow copy of the map(shadow in go) -> [maps.	Clone]
	m1 := map[string]map[string]int{
		"M1": map[string]int{
			"key1": 10,
			"key2": 20,
		},
		"M2": map[string]int{
			"key1": 100,
			"key2": 200,
		},
	}

	m1copy := maps.Clone(m1)

	m1copy["M1"]["key1"] = 30

	m1copy["M2"] = map[string]int{
		"new": 500,
	}

	fmt.Println(m1)
	fmt.Println(m1copy)

}
