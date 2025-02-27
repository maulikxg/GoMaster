package main

import (
	"time"
)

func sendData(ch chan<- int) {
	ch <- 41
}

func main() {

	ch := make(chan int)

	sendCh := (chan<- int)(ch)
	go sendData(sendCh) // sending the data

	recvc := (<-chan int)(ch) // receiving the data
	x := <-recvc
	println(x)

	time.Sleep(time.Second)

}
