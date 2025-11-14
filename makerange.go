package student

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	ar := []int{}
	for i := 0; i < (max - min); i++ {
		ar = append(ar, min+i)
	}
	return ar
}
