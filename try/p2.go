package main

import "fmt"

func main() {

	defer func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("This is R:", r)
			}
		}()
	}()

	panic("this is me hii")
}
