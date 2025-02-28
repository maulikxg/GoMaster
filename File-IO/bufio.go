package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("file.txt")

	if err != nil {
		fmt.Println("Error in creating file.", err)
		return
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	write, err := writer.Write([]byte("This is the some stuff that from the bufio writer,hhehe"))
	if err != nil {
		return
	}
	writer.Flush()

	fmt.Println("Wrote", write)

}
