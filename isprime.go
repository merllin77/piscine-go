package student

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	if nb%2 == 1 {
		return true
	} else {
		return false
	}
}
