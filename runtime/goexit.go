package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	go worker()
	fmt.Println("Hello World")
	time.Sleep(1 * time.Second)
}

func worker() {

	defer fmt.Println("im from the defer")
	fmt.Println("Im the starting of the worker.")
	runtime.Goexit()
	fmt.Println("Im the ending the worker.")
	defer fmt.Println("Ding Ding ding")
}
