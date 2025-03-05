package main

import (
	"fmt"
)

func main() {

	//defer func() {
	//
	//	if r := recover(); r != nil {
	//		fmt.Println("from the recover")
	//		fmt.Println(r)
	//	}
	//}()
	//
	//fmt.Println("this is 1st")
	//panic("this is panic ")
	//
	//fmt.Println("this is 2nd")

	//defer fmt.Println("First deferred function")
	//defer fmt.Println("Second deferred function")
	//
	//fmt.Println("Before panic")
	//panic("Crash!")

	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Println("Recovered:", r)
	//		panic("New panic after recover!") // This new panic is NOT caught
	//	}
	//}()
	//
	//panic("First nanan panic")

	//go func() {
	//	defer func() {
	//		if r := recover(); r != nil {
	//			fmt.Println("Recovered in goroutine:", r) // ✅ Now it works
	//		}
	//	}()
	//	panic("Goroutine panic!")
	//}()
	//
	//time.Sleep(1 * time.Second)

	//go func() {
	//	panic("Unexpected error in goroutine!")
	//}()
	//
	//time.Sleep(1 * time.Second)
	//fmt.Println("Main function execution...")

	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Println("First recover:", r)
	//	}
	//}()
	//
	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Println("Second recover:", r)
	//	}
	//}()
	//
	//panic("Panic triggered!")

	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Println("Recovered and continuing execution:", r)
	//	}
	//}()
	//
	//panic("Panic occurred!")
	//fmt.Println("This line will never execute")

	//defer func() {
	//	fmt.Println("First defer")
	//	panic("Defer panic!")
	//}()
	//
	//defer func() {
	//	if r := recover(); r != nil {
	//		fmt.Println("Recovered from:", r)
	//	}
	//}()
	//
	//panic("Main panic!")

	//go func() {
	//	defer func() {
	//		if r := recover(); r != nil {
	//			fmt.Println("Recovered from panic in goroutine:", r)
	//		}
	//	}()
	//
	//	panic("Panic inside goroutine!")
	//}()
	//
	//time.Sleep(1 * time.Second)

	//for i := 1; i <= 5; i++ {
	//	go safeGoroutine(i)
	//}
	//
	//time.Sleep(1 * time.Second)

}

func safeGoroutine(id int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in goroutine %d: %v\n", id, r)
		}
	}()

	if id%2 == 0 {
		panic(fmt.Sprintf("Panic in goroutine %d!", id))
	}
}
