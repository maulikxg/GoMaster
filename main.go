package main

import (
	"fmt"
	"pacx/mazamamaths"
)

func Summ(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

type Car struct {
	Make  string
	Model string
	Year  int
}

func (c *Car) YearChange(newY int) {
	c.Year = newY
}

func modify(arr *[3]int) {
	arr[0] = 100
}

func maza(a, b int, op func(int, int) int) int {
	return op(a, b)
}

func main() {
	fmt.Println("Hello World")

	// my custom package
	// result := mm.Add(2, 3)
	// fmt.Println(result)
	//
	check := mazamamaths.Add2(4, 2)
	fmt.Println(check)

	//mm.PublicFunction()

	//ans := Summ(1, 2, 3, 4, 5)
	//fmt.Println(ans)

	//var p *int
	//x := 10
	//p = &x
	//fmt.Printf("before the change :%d \n", *p)
	//*p = 100
	//fmt.Printf("after the change :%d \n", *p)

	//type Car struct {
	//	Brand string
	//}
	//
	//c := Car{"Tesla"}
	//fmt.Println(c.Brand)
	//pc := &c
	//pc.Brand = "BMW"     // Modifies original struct
	//fmt.Println(c.Brand) // BMW

	//var arr [3]int
	//arr[0] = 1
	//arr[1] = 2
	//arr[2] = 3
	//fmt.Println(arr)

	//mycar := Car{
	//	Make:  "Toyota",
	//	Model: "Paris",
	//	Year:  2021,
	//}

	//fmt.Println(mycar)
	//mycar.YearChange(2025)
	//fmt.Println(mycar)
	//
	//arr := [3]int{89, 90, 91}
	//fmt.Println(arr)
	//modify(&arr)
	//fmt.Println(arr)

	//s := []int{1, 2, 3, 4, 5}
	//fmt.Printf("%v %d %d\n", s, len(s), cap(s))
	//s = append(s, 6, 7, 8, 9, 10, 11)
	//fmt.Printf("%v %d %d\n", s, len(s), cap(s))
	//s = append(s, 12, 13, 14)
	//fmt.Printf("%v %d %d\n", s, len(s), cap(s))

	//var s []int
	//fmt.Println(s == nil)

	//nums := []int{10, 20, 30, 40, 50}
	//
	//for index, value := range nums {
	//	fmt.Println("index:", index, "value:", value)
	//}
	//for _, value := range nums { // Ignoring index
	//	fmt.Println(value)
	//}
	//for index := range nums { // Ignoring value
	//	fmt.Println(index)
	//}

	//var mymap map[string]int
	//fmt.Printf("%#v\n", mymap)
	//mymap = make(map[string]int)
	//fmt.Printf("%#v\n", mymap)

	//grads := map[string]int{
	//	"maths":     90,
	//	"chemistry": 80,
	//	"Physics":   50,
	//}
	//
	//grads["maths"] = 85
	//grads["History"] = 70
	//
	//fmt.Println(grads)
	//
	//var m map[string]int
	//m["test"] = 100

	add := func(a, b int) int {
		return a + b
	}
	//fmt.Println(add(2, 3))

	result := maza(1, 2, add)
	fmt.Println(result)

}
