package main

import "fmt"

type MyInterface interface {
	DoSomething()
	DoNothing()
}

type MyStruct struct{}

func (m MyStruct) DoSomething() {
	fmt.Println("Doing something!")
}

func (m MyStruct) DoNothing() {
	fmt.Println("Do nothing")
}

func main() {

	var yo MyStruct

	var i MyInterface

	i = &yo

	i.DoSomething()
	i.DoNothing()

}
