package student

func Map(f func(int) bool, a []int) []bool {
	rsltbool := []bool{}
	for _, r := range a {
		if f(r) == true {
			rsltbool = append(rsltbool, true)
		} else {
			rsltbool = append(rsltbool, false)
		}
	}
	return rsltbool
}

/*func main() {
	a := []int{1, 2, 3, 4, 5, 6, 13}
	result := Map(IsPrime, a)
	println(result)
} */
