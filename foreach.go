package student

import "github.com/01-edu/z01"

func PrintNbr(nbr int) {
	if nbr >= 10 {
		PrintNbr(nbr / 10)
	}
	z01.PrintRune(rune(nbr%10 + '0'))
}

func ForEach(f func(int), a []int) {
	for _, r := range a {
		z01.PrintRune(r)
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	ForEach(PrintNbr, a)
}
