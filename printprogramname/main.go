package main

import (
	"os"
	"path/filepath"

	"github.com/01-edu/z01"
)

func main() {
	progname := filepath.Base(os.Args[0])
	for i := 0; i < len(progname); i++ {
		z01.PrintRune(rune(progname[i]))
	}
	z01.PrintRune('\n')
}
