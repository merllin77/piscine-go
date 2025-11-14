package student

func AppendRange(min, max int) []int {
	arlen := max - min
	if arlen < 0 {
		return nil
	} else {
		ar := make([]int, arlen)
		for i := 0; i < arlen; i++ {
			ar[i] = min + i
		}
		return ar
	}
}
