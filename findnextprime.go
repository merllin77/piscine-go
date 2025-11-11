package main

import "fmt"

func IsPrime(nb int) bool { // vriskei tous protous arithmous mexri to nb
	if nb <= 1 {
		return false
	}
	if nb == 2 {
		return true
	}
	if nb%2 == 0 {
		return false
	}
	for i := 3; i*i <= nb; i = i + 2 {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func FindNextPrime(nb int) int {
	if nb <= 1 {
		return 2
	}

	for {
		if IsPrime(nb) {
			return nb
		}
		nb++
	}
}

func main() {
	fmt.Println(FindNextPrime(1000))
	fmt.Println(FindNextPrime(0))
}
