package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("file.txt")
	if err != nil {
		fmt.Println("Error in creating file", err)
		return
	}
	fmt.Println("File Created Sucessfully")

	// checking the file existence
	_, err = os.Stat("file.txt")
	if os.IsExist(err) {
		fmt.Println("No file Found")
	} else {
		fmt.Println("file is here.")
	}

	defer file.Close()

}
