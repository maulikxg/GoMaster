package main

func main() {
	A()
}

func A() {
	B()
}

func B() {
	C()
}

func C() {
	panic("panic is here.")
}
