package student

func AlphaCount(s string) int {
	counter := 0
	for _, r := range s {
		if r >= 'a' && r <= 'z' || (r >= 'A' && r <= 'Z') {
			counter++
		}
	}
	if counter == 0 {
		return 0
	} else {
		return counter
	}
}
