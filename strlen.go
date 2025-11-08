package student

var cnt int

func StrLen(s string) int {
	for i := 1; i <= len(s); i++ {
		cnt++
	}
	return cnt
}
