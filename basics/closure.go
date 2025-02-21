package main

import "fmt"

func main() {

	yo := 0

	counter := func() int {
		yo++
		return yo
	}

	fmt.Println(counter())
	fmt.Println(counter())
}
