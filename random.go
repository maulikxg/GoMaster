package main

import "fmt"

func main() {
	var a = []int{0, 1, 2, 3, 4}

	fmt.Println("Array:", a, len(a), cap(a))
	a = a[0:2]
	fmt.Println("Array:", a, len(a), cap(a))
	a = a[:3]
	fmt.Println("Array:", a, len(a), cap(a))
	a = a[2:]
	fmt.Println("Array:", a, len(a), cap(a))
	a = a[0:]
	fmt.Println("Array:", a, len(a), cap(a))
	a = a[:2]
	fmt.Println("Array:", a, len(a), cap(a))
}
