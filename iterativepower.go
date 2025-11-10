package student

func IterativePower(nb int, power int) int {
	if nb > 20 || power < 0 || power > 20 {
		return 0
	}
	if nb == 0 {
		return 1
	}
	rslt := 1
	for i := 1; i <= power; i++ {
		rslt = rslt * nb
	}
	return rslt
}
