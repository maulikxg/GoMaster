package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"time"
)

func main() {

	debug.SetGCPercent(100) // default

	printmemstat("Start of the Program")

	var memHog [][]byte

	for i := 0; i < 10000; i++ {
		memHog = append(memHog, make([]byte, 10))

		if i%2000 == 0 {
			printmemstat(fmt.Sprintf("after the %d allocation", i))
		}
	}

	printmemstat("After the alloation done")
	memHog = nil

	runtime.GC()

	printmemstat("After forcing gc")

	_ = make([]byte, 50*1024*1024)

	printmemstat("After Large Allocation")

	time.Sleep(2 * time.Second)

	printmemstat("After Sleep (GC May Have Run)")

	fmt.Println("End of program")

}

func printmemstat(msg string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Println("------------------------------------------------")
	fmt.Println(msg)
	fmt.Printf("HeapAlloc: %v KB\n", m.HeapAlloc/1024)
	fmt.Printf("HeapSys: %v KB\n", m.HeapSys/1024)
	fmt.Printf("NumGC: %v\n", m.NumGC)
	fmt.Println("------------------------------------------------\n")
}
