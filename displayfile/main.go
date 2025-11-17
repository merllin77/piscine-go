package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("File name missing")
		return
	}
	if len(args) > 1 {
		fmt.Println("Too many arguments")
		return
	}

	file, err := os.Open(args[0]) // open the file (contenct has the file's data in bytes)
	if err != nil {               // err == nil means there is no error.
		fmt.Println("Error opening file")
		return
	}
	defer file.Close() // closing the file after function ends (Important!)

	content, err := io.ReadAll(file) // content has the file's content in bytes
	if err != nil {
		fmt.Println("Error reading files", err)
	}
	fmt.Println(string(content)) // transpose bytes to string for output
}
