package main

import "github.com/01-edu/z01"

func main() {
	str := "abcdefghijklmnopqrstuvwxyz"
	for i := len(str) - 1; i >= 0; i-- {
		z01.PrintRune(rune(str[i]))
	}
}
