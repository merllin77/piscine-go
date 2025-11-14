package student

func AppendRange(min, max int) []int {
	arlen := max - min
	if arlen <= 0 {
		return []int{}
	}
	ar := make([]int, arlen)
	for i := 0; i < arlen; i++ {
		ar[i] = min + i
	}
	return ar
}
