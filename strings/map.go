package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "Golang"

	modified := strings.Map(func(r rune) rune {
		if r == 'G' {
			return 'X'
		}
		return r
	}, s)

	fmt.Println(modified)
}
