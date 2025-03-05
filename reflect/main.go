package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Foo struct {
	A int
	B string
}

func main() {

	sl := []int{1, 2, 3}
	greeting := "hello"
	greetingPtr := &greeting
	f := Foo{A: 10, B: "Salutations"}
	fp := &f

	slType := reflect.TypeOf(sl)
	gType := reflect.TypeOf(greeting)
	grpType := reflect.TypeOf(greetingPtr)
	fType := reflect.TypeOf(f)
	fpType := reflect.TypeOf(fp)

	// fmt.Println(slType, gType, grpType, fType, fpType)

	checker(slType, 0)
	checker(gType, 0)
	checker(grpType, 0)
	checker(fType, 0)
	checker(fpType, 0)

}

func checker(t reflect.Type, depth int) {

	fmt.Println(strings.Repeat("\t", depth), "Type is", t.Name(), "<->", "Kind is", t.Kind())

	switch t.Kind() {

	case reflect.Array, reflect.Slice, reflect.Map, reflect.Chan, reflect.Ptr:
		fmt.Println(strings.Repeat("\t", depth+1), "Contained type ->>:")
		checker(t.Elem(), depth+1)

	case reflect.Struct:
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			fmt.Println(strings.Repeat("\t", depth+1), "Field", i+1, "name is", f.Name, ". Type is", f.Type.Name(), "and kind is", f.Type.Kind())

		}
	}

}
