package student

func Index(s string, toFind string) int {
	totallen := len(s) - len(toFind)
	for i := 0; i <= totallen; i++ {
		for j := 0; j < len(toFind); j++ {
			if s[i+j] != toFind[j] {
				break
			} else {
				return i
			}
		}
	}
	return 0
}
