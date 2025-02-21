package mazamamaths

import "fmt"

func Add2(x, y int) int {
	return x - y
}

// Exported function (Accessible outside the package)
func PublicFunction() {
	fmt.Println("I am an exported function")
}

// Unexported function (Not accessible outside this package)
func privateFunction() {
	fmt.Println("I am a private function")
}
