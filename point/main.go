package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func printXval(nx int) { // transpose int x from ptr.x to a rune
	if nx >= 10 { // if int x has more than 1 digit
		printXval(nx / 10) // recursion of printXval printing all digits except last (/10 moves left of nx)
	}
	z01.PrintRune(rune(nx%10 + '0')) // print last digit as a rune (with +'0')
}

func printYval(ny int) {
	if ny >= 10 {
		printYval(ny / 10)
	}

	z01.PrintRune(rune(ny%10 + '0'))
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func main() {
	points := &point{}

	setPoint(points)

	// fmt.Printf("x = %d, y = %d \n", points.x, points.y)

	printStr("x = ")
	printXval(points.x)
	z01.PrintRune(',')
	printStr(" y = ")
	printYval(points.y)
	z01.PrintRune('\n')
}
