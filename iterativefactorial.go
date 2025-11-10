package student

func IterativeFactorial(nb int) int {
	rslt := 1
	if nb == 0 {
		return 1
	} else if nb > 0 {
		for i := 1; i <= nb; i++ {
			rslt = rslt * i
		}
		return rslt
	} else {
		return 0
	}
}
