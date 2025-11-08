package student

var cnt int

func StrLen(s string) int {
	for i := 1; i <= len(s)-1; i++ {
		cnt++
	}
	return cnt
}
