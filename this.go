package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan int // ???

	go func() {
		fmt.Println(<-ch) // ???
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("Main function continues")
}
