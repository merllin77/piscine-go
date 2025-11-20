package student

func CountIf(f func(string) bool, tab []string) int {
	cnt := 0
	for _, r := range tab {
		if f(r) == true {
			cnt++
		}
	}
	return cnt
}
