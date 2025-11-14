package student

func AppendRange(min, max int) []int {
	if min >= max {
		return nil
	}
	arlen := max - min
	ar := make([]int, arlen)
	for i := 0; i < arlen; i++ {
		ar[i] = min + i
	}
	return ar
}
