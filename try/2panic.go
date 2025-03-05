package main

import "time"
import "fmt"

func workerr(id int) {

	defer fmt.Println("Worker", id, "exited")
	panic(fmt.Sprintf("panic in worker %d", id))
}

func main() {
	for i := 0; i < 3; i++ {
		go workerr(i)
	}
	time.Sleep(time.Second)
}
