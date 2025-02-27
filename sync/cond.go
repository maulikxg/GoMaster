package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	queue    []int // buffer for the store the item
	mu       sync.Mutex
	cond     = sync.NewCond(&mu) // cond for mutex
	capacity = 5                 // buffer size
)

func main() {

	go producer(1)
	go producer(2)

	go consumer(3)
	go consumer(4)

	time.Sleep(time.Second * 5)
}

func producer(id int) {

	for i := 1; i <= 3; i++ {

		mu.Lock()
		if len(queue) == capacity {
			fmt.Printf("The producer %d cant add the item %d\n", id, i)
			cond.Wait() // waiting for the signal now
		}

		queue = append(queue, i)
		fmt.Printf("The Produer %d addeed the the item %d\n", id, i)
		cond.Signal() // telling that i have done the change
		mu.Unlock()
		time.Sleep(time.Millisecond * 500)
	}
}

func consumer(id int) {

	for i := 1; i <= 3; i++ {

		mu.Lock()
		if len(queue) == 0 {
			fmt.Printf("The consumer %d is saying no item now\n", id)
			cond.Wait() // wait for the signal that resorces job are added
		}

		item := queue[0]
		queue = queue[1:]

		fmt.Printf("The consumer %d eated item %d\n", id, item)
		cond.Signal() // telling that i have changed
		mu.Unlock()
		time.Sleep(time.Millisecond * 700)
	}
}
