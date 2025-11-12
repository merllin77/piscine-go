package student

func IsAlpha(s string) bool {
	for _, r := range s {
		if r < '0' {
			return false
		}
	}
	return true
}
