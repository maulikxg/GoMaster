package main

import "fmt"

func main() {

	original := []int{1, 2, 3}
	cp := original

	cp[0] = 10
	fmt.Println(cp)
	fmt.Println(original)

	or := []int{10, 20, 30, 40}
	cz := make([]int, len(or))
	copy(cz, or)

	cz[0] = 1
	fmt.Println(cz)
	fmt.Println(or)
}
