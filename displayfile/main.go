package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	line := ""

	if len(args) == 0 {
		fmt.Println("File name missing")
	}
	if len(args) > 1 {
		fmt.Println("Too many arguments")
	}

	content, err := os.Open(args[0]) // open the file (contenct has the file's data in bytes)
	if err != nil {                  // err == nil means there is no error.
		fmt.Println("Error opening file")
		return
	}

	scanner := bufio.NewScanner(content) // create a "scanner" to read lines
	for scanner.Scan() {                 // moves to the next line and returns true while there are lines left.
		line = scanner.Text() // get current line as string
	}
	fmt.Println(line)
	defer content.Close()
}
