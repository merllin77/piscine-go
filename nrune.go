package student

func NRune(s string, n int) rune {
	for index, r := range s {
		if index == n-1 {
			return r
		}
	}
	return 0
}
