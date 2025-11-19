package student

import "fmt"

func IsPrime(nbr int) bool {
	if nbr <= 1 { // 1 is not a Prime
		return false
	}

	if nbr == 2 { // 2 is the first Prime number
		return true
	}

	if nbr%2 == 0 { // If modulo == 0 is NOT a Prime number
		return false
	}

	for i := 3; i*i <= nbr; i = i + 2 { // check after number 2 which is the first Prime number
		if nbr%i == 0 { // if nbr is even return TRUE
			return false
		}
	}
	return true
}

func Map(f func(int) bool, a []int) []bool {
	rsltbool := []bool{}
	for _, r := range a {
		if f(r) == true {
			rsltbool = append(rsltbool, true)
		} else {
			rsltbool = append(rsltbool, false)
		}
	}
	return rsltbool
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 13}
	result := Map(IsPrime, a)
	fmt.Println(result)
}
