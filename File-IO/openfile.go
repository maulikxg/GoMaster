package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.OpenFile("file.txt", os.O_APPEND|os.O_WRONLY, 644) // ONLY O_WRONLY is for overwrite

	if err != nil {
		fmt.Println("open file failed", err)
		return
	}

	defer file.Close()

	file.Write([]byte("THIS IS FINAL COUTDOWN."))

	fmt.Println("File over written Sucessfully")

}
