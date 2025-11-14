package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	progname := os.Args[0]
	slash := 0
	for i, r := range progname {
		if r == '/' {
			slash = i
		}
	}
	for _, r := range progname[slash+1:] {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
