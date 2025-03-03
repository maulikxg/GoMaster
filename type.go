package main

import "fmt"

type Mytype = string
type Mytype2 string

func main() {

	var name Mytype = "maulik"

	var surname Mytype2 = "yoho"

	fmt.Printf("%T , %s\n", name, name)

	fmt.Printf("%T , %s\n", surname, surname)
}
