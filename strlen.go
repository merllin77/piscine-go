package student

var cnt int

func StrLen(s string) int {
	r := []rune(s)
	for i := 1; i <= len(r); i++ {
		cnt++
	}
	return cnt
}
