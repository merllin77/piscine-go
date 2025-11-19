package student

func Any(f func(string) bool, a []string) bool {
	for _, r := range a {
		if f(r) == true {
			return true
		}
	}
	return false
}

/*func main() {
	a1 := []string{"Hello", "how", "are", "you"}
	a2 := []string{"This", "is", "4", "you"}

	result1 := Any(student.IsNumeric, a1)
	result2 := Any(student.IsNumeric, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}*/
