package student

func Map(f func(int) bool, a []int) []bool {
	rsltbool := make([]bool, len(a))
	for i, r := range a {
		rsltbool[i] = f(r)
	}
	return rsltbool
}
