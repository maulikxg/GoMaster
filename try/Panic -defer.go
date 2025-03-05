//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func main() {
//
//	go func() {
//
//		defer func() {
//
//			if r := recover(); r != nil {
//				fmt.Println("this is the issue", r)
//			}
//
//		}()
//		arr := [100]int{}
//		arr[1] = arr[500]
//
//	}()
//
//	time.Sleep(time.Second)
//	fmt.Println("yes")
//}

package main

import (
	"fmt"
	"time"
)

func worker() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in worker:", r)
		}
	}()
	panic("something went wrong")
}

func main() {
	go worker()
	time.Sleep(time.Second)
	fmt.Println("Main program continues")
}
