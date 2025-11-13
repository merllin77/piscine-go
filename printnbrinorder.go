package student

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	var (
		counterar [10]int // create a var type array ([]) with lentgh 10 containing integers
		digit     int
	)

	if n == 0 {
		counterar[0] = 1
	}

	for n > 0 { // for positive only integers
		digit = n % 10     // take the last digit of int
		counterar[digit]++ // adds (++) how many times the last digit [digit] appears to the array of int "counterar"
		n = n / 10         // removes the last digit of int
	}

	for i := 0; i <= 9; i++ { // goes through all possible digits from 0-9 (as digits)
		for counterar[i] > 0 { // print the digit if the i EXISTS in counterar and is bigger than 0 (if there is a digit with zero counter skips the loop)
			z01.PrintRune(rune('0' + i))
			counterar[i]-- // brings the value back to 0 to exit loop (for counterar[i])
		}
	}
}
