package main

import "fmt"

func fibo() func() int {
	a, b := 0, 1
	return func() int {
		temp := a
		a, b = b, a+b
		return temp
	}
}
func main() {
	fib := fibo()

	fmt.Println(fib())
	fmt.Println(fib())
	fmt.Println(fib())
	fmt.Println(fib())
	fmt.Println(fib())
	fmt.Println(fib())

}
