package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	for _, r := range arg[1:] {
		for _, j := range r {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}
