package student

func Capitalize(s string) string {
	newstr := []rune{}
	iswordStart := true
	for _, r := range s {
		if (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') { // if r is alphanumeric
			if iswordStart { // if we are in the start of the word
				if r >= 'a' && r <= 'z' {
					r = r - 32
				}
				iswordStart = false
			} else { // if we are NOT in the start (inside word)
				if r >= 'A' && r <= 'Z' {
					r = r + 32
				}
			}
		} else {
			iswordStart = true
		}
		newstr = append(newstr, r)
	}
	return string(newstr)
}
