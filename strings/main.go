package main

import (
	"fmt"
	"strings"
)

func main() {

	q := []string{"yo", "ho", "maulik", "heloo"}
	r := strings.Join(q, "/")
	fmt.Println(r)

}
