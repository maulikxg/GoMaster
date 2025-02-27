package main

import (
	"fmt"
	"math"
)

func main() {

	in := workGenerator([]int{0, 1, 2, 3, 4, 5, 6, 7, 8})

	out := filter(in)
	out = square(out)
	out = half(out)

	for value := range out {
		fmt.Println(value)
	}

}

func workGenerator(work []int) <-chan int {

	ch := make(chan int)

	go func() {
		defer close(ch)

		for _, w := range work {
			ch <- w
		}

	}()

	return ch

}

func filter(in <-chan int) <-chan int {

	out := make(chan int)

	go func() {
		defer close(out)

		for i := range in {
			if i%2 == 0 {
				out <- i
			}
		}

	}()

	return out
}

func square(in <-chan int) <-chan int {

	out := make(chan int)

	go func() {

		defer close(out)

		for i := range in {
			value := math.Pow(float64(i), 2)
			out <- int(value)
		}
	}()

	return out
}

func half(in <-chan int) <-chan int {

	out := make(chan int)

	go func() {

		defer close(out)

		for i := range in {
			value := i / 2
			out <- value
		}
	}()

	return out

}
