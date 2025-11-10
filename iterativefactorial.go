package student

func IterativeFactorial(nb int) int {
	rslt := nb
	if nb == 0 {
		return 1
	} else if nb > 0 {
		for i := nb - 1; i >= 1; i-- {
			rslt = rslt * i
		}
		return rslt
	} else {
		return 0
	}
}
