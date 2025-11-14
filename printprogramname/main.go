package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	progname := os.Args[0]
	slash := 0
	for i := 0; i < len(progname); i++ {
		if progname[i] == '/' {
			slash = i
		}
	}
	for i := slash + 1; i < len(progname); i++ {
		z01.PrintRune(rune(progname[i]))
	}

	z01.PrintRune('\n')
}
