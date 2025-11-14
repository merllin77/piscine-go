package student

func SplitWhiteSpaces(s string) []string {
	newstr := []string{}
	tmpword := []rune{}
	for _, r := range s {
		if r == ' ' || r == '\t' || r == '\n' {
			if len(tmpword) > 0 {
				newstr = append(newstr, string(tmpword))
				tmpword = []rune{}
			}
		} else {
			tmpword = append(tmpword, r)
		}
	}
	newstr = append(newstr, string(tmpword))
	return newstr
}
