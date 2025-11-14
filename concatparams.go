package student

func ConcatParams(args []string) string {
	var answer string
	for i, r := range args {
		answer = answer + r
		if i != len(args)-1 {
			answer = answer + "\n"
		}
	}
	return answer
}
