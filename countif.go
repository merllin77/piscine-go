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

// func main() {
// 	tab1 := []string{"Hello", "how", "are", "you"}
// 	tab2 := []string{"This", "1", "is", "4", "you"}
// 	answer1 := CountIf(student.IsNumeric, tab1)
// 	answer2 := CountIf(student.IsNumeric, tab2)
// 	fmt.Println(answer1)
// 	fmt.Println(answer2)
// }
