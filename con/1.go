package main

import (
	"fmt"
	"time"
)

func somefunc(num string) {
	fmt.Println(num)
}

func main() {
	go somefunc("1")
	go somefunc("2")
	go somefunc("3")

	time.Sleep(time.Second * 2)

	fmt.Println("Hi form the main")
}
