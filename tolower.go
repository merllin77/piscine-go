package student

func ToLower(s string) string {
	nstring := []rune{}
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			r = r + 32
		}
		nstring = append(nstring, r)
	}
	return string(nstring)
}
