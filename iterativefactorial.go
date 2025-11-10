package student

func IterativeFactorial(nb int) int {
	rslt := 1
	if nb == 0 {
		return 1
	} else if nb > 0 {
		for i := 1; i <= nb; i++ {
			if rslt > 0 && rslt > 1<<63/i {
				return 0
			}
			rslt = rslt * i
		}
		return rslt
	} else {
		return 0
	}
}
