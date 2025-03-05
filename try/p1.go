package main

import "fmt"

func main() {

	defer func() {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("secret:", r)
			}
		}()

		if r := recover(); r != nil {
			fmt.Println("ind:", r)
		}

		panic("panic is inside.")

	}()
	panic("this is form main.")
}
