package main

import (
	"fmt"
)

//func modifySlice(s *[]int) {
//	fmt.Println("\nInside modifySlice function:")
//	fmt.Println("Before modification:", s)
//	(*s)[0] = 500        // Modify first element
//	*s = append(*s, 999) // Append a new element
//
//	fmt.Println("After modification:", s)
//}

func modifySlice(s []int) []int {
	fmt.Println("\nInside modifySlice function:")
	fmt.Println("Before modification:", s)
	s[0] = 500         // Modify first element
	s = append(s, 999) // Append a new element

	fmt.Println("After modification:", s)
	return s
}

func main() {
	// Step 1: Create an array and slice it
	arr := [5]int{10, 20, 30, 40, 50}
	slice := arr[1:4] // slice -> {20, 30, 40}

	fmt.Println("Initial Array:", arr)
	fmt.Println("Initial Slice:", slice)

	// Step 2: Pass slice to function
	//modifySlice(&slice)

	slice = modifySlice(slice)

	fmt.Println("\nAfter function call:")
	fmt.Println("Array:", arr)
	fmt.Println("Slice:", slice)
}

//package main
//
//import "fmt"
//
//func main() {
//	var printerFn []func()
//	var x = [4]int{1, 2, 3, 4}
//	for _, v := range x {
//		printerFn = append(printerFn, func() {
//			fmt.Printf("\n%v\n", v)
//		})
//	}
//	for _, fn := range printerFn {
//		fn()
//	}
//}
