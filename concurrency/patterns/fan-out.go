package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {

	for job := range jobs {
		fmt.Printf("Worker %d processing job %d\n", id, job)
		time.Sleep(time.Second)
		results <- job * 2
	}

}

func main() {

	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// creating the workers
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	// sendig the jobs
	for j := 0; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// receiving the results
	for k := 0; k <= 5; k++ {
		fmt.Println("Results :", <-results)
	}

}
