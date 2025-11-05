package main

import "github.com/01-edu/z01"

func main() {
	n := "0123456789"
	maxlen := len(n)
	for i := 0; i <= maxlen-1; i++ {
		z01.PrintRune(rune(n[i]))
	}
	z01.PrintRune('\n')
}
