package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"sync"
	"time"
)

func main() {

	f, err := os.Create("goroutine.prof")
	if err != nil {
		fmt.Println("issues in creating file")
		return
	}

	defer f.Close()

	var wg sync.WaitGroup
	wg.Add(3)

	go goroutineFunction(&wg)
	go goroutineFunction(&wg)
	go goroutineFunction(&wg)

	wg.Wait()
	pprof.Lookup("goroutine").WriteTo(f, 0)

}

func goroutineFunction(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second * 1)
}
