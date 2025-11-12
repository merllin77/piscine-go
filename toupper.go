package student

func ToUpper(s string) string {
	newstr := []rune{}
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			r = r - 32
		}
		newstr = append(newstr, r)
	}
	return string(newstr)
}
