package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func main() {

	// normal stuff

	//m := Message{"Alice", "Hello", 1294706395881547000}
	//
	//b, err := json.Marshal(m)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(string(b))
	//
	//var d Message
	//
	//_ = json.Unmarshal(b, &d)
	//
	//fmt.Println(d)

	// with interface
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)

	var f interface{}

	err := json.Unmarshal(b, &f)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println(f)

}
