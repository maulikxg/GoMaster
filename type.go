package main

import "fmt"

type Mytype = string
type Mytype2 string

func main() {

	//var name Mytype = "maulik"
	//
	//var surname Mytype2 = "yoho"
	//
	//fmt.Printf("%T , %s\n", name, name)
	//
	//fmt.Printf("%T , %s\n", surname, surname)

	// switch day := "monday"; day {

	//case "monday":
	//	fmt.Println("This is monday")
	//	fallthrough
	//case "tuesday":
	//	fmt.Println("This is tuesday")
	//default:
	//	fmt.Println("This is default")
	//
	//}

	yo := add()
	yo(5)
	fmt.Println(yo(10))

	type person struct {
		name string
		age  int
	}

	p := person{"maulik ", 50}
	fmt.Println(p)

}

func add() func(int) int {
	sum := 0

	return func(v int) int {
		sum += v
		return sum
	}

}
