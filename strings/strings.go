package main

import (
	"fmt"
	"strings"
)

func main() {

	//q := []string{"yo", "ho", "maulik", "heloo"}
	//r := strings.Join(q, "/")
	//fmt.Println(r)

	//a := "go,java,hwa"
	//b := strings.Split(a, ",")
	//fmt.Println(b)

	//c := "    Hello !! Go   "
	//d := strings.Fields(c)
	//fmt.Println(d)

	//fmt.Println(strings.Contains("golang", "go"))
	//fmt.Println(strings.ContainsAny("team", "t"))
	//fmt.Println(strings.ContainsAny("   ", " "))

	//fmt.Println(strings.HasPrefix("golang", "Go"))
	//fmt.Println(strings.HasSuffix("golang", "lang"))

	//fmt.Println(strings.Index("golang", "lang"))
	//fmt.Println(strings.LastIndex("golang", "Go"))
	//fmt.Println(strings.Count("golanggo", "go"))

	//s := "Go is fast , Go is powerful"
	//fmt.Println(strings.Replace(s, "Go", "Golang", 1))
	//fmt.Println(strings.ReplaceAll(s, "Go", "Golang"))

	//fmt.Println(strings.ContainsRune("tems", 97))
	//fmt.Println(strings.ContainsRune("teams", 97))

	//fmt.Println(strings.Count("efive", "e"))
	//fmt.Println(strings.Count("five", ""))

	//Cut
	/*	show := func(s, sep string) {

			before, after, found := strings.Cut(s, sep)

			fmt.Printf("Cut(%q , %q) = %q , %q , %v)\n", s, sep, before, after, found)

		}

		show("Gopher", "Go")
		show("Gopher", "ph")
		show("Gopher", "er")
		show("Gopher", "Badger") */

	// fmt.Println(strings.EqualFold("ab", "AB"))
	//
	//s := []string{"tu", "hai", "to", "me"}
	//r := strings.Join(s, "   ")
	//fmt.Println(r)

	//fmt.Println(strings.LastIndex("go langgogl", "go"))
	//fmt.Println(strings.LastIndex("go golang", "go"))
	//fmt.Println(strings.LastIndex("go gopher", "ra"))

	//fmt.Println(strings.LastIndexAny("go gopher", "Gg"))
	//fmt.Println(strings.LastIndexAny("go gopher", "rodent"))
	//fmt.Println(strings.LastIndexAny("go gopher", "gail"))

	// fmt.Println("ba" + strings.Repeat("nana", 3))

	//fmt.Println(strings.Replace("maulik maulik maulik", "mau", "mo", -1))

	//fmt.Println(strings.ReplaceAll("maulik maulik maulik", "li", "lo"))

	//fmt.Printf("%q\n", strings.SplitAfter("a,b,c", ","))
	//fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", ",", 2))

	fmt.Println(strings.Trim("maulikmaulikma", "ma"))

}
