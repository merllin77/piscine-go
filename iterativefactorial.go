package piscine

func IterativeFactorial(nb int) int {
	if nb == 0 {
		return 1
	} else if nb > 0 {
		for i := nb - 1; i >= 1; i-- {
			nb = nb * i
		}
		return nb
	} else {
		return 0
	}
}
