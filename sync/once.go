package main

import (
	"fmt"
	"sync"
)

func main() {

	var cnt int

	increment := func() {
		cnt++
	}

	var wg sync.WaitGroup
	var once sync.Once

	wg.Add(100)

	for i := 0; i < 100; i++ {

		go func() {
			defer wg.Done()

			once.Do(increment)
		}()
	}

	wg.Wait()
	fmt.Printf("Count is : %d\n", cnt)

}
