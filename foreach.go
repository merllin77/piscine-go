package main

import "fmt"

func PrintNbr(nbr int) {
	if nbr >= 10 {
		PrintNbr(nbr / 10)
	}
	fmt.Print(rune(nbr%10 + '0'))
}

func ForEach(f func(int), a []int) {
	for _, r := range a {
		fmt.Print(r)
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	ForEach(PrintNbr, a)
}
