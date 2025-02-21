package main

import "fmt"

const (
	A = (iota + 1) * 10
	B
	C
	D
)

func main() {
	fmt.Println(A, B, C, D)
}
