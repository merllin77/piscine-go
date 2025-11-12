package student

func IsPrintable(s string) bool {
	for _, r := range s {
		if r < ' ' {
			return false
		}
	}
	return true
}
