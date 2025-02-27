package main

import (
	"fmt"
	"sync"
)

const (
	totaljobs   = 4
	totalworker = 2
)

func main() {

	jobs := make(chan int, totaljobs)

	results := make(chan int, totaljobs)

	for w := 1; w <= totalworker; w++ {
		go workerr(w, jobs, results)
	}

	// sending jobs
	for j := 1; j <= totaljobs; j++ {
		jobs <- j
	}
	close(jobs)

	// receiving jobs
	for a := 1; a <= totaljobs; a++ {
		<-results
	}
	close(results)
}

func workerr(id int, jobs <-chan int, results chan<- int) {

	var wg sync.WaitGroup

	for j := range jobs {

		wg.Add(1)

		go func(job int) {

			defer wg.Done()

			fmt.Printf("Worker %d started job %d\n", id, job)

			result := job * 2
			results <- result

			fmt.Printf("Worker %d finished job %d\n", id, job)
		}(j)

	}

	wg.Wait()

}
