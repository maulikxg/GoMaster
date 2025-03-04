package main

import (
	"fmt"
	"unsafe"
)

type yo struct {
	A int32
	B float64
}

func main() {

	var s yo

	s.A = 1
	s.B = 2

	ptr := unsafe.Pointer(&s)
	fmt.Println(uintptr(ptr))

	aptr := (*int)(ptr) // offset 0
	*aptr = 100

	bptr := (*float64)(unsafe.Pointer(uintptr(ptr) + unsafe.Offsetof(s.B))) // for offset of B
	fmt.Println(bptr)
	*bptr = 200

	fmt.Println(s.A, s.B)

}
