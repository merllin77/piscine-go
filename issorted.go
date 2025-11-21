package student

func Compare2nbr(n1, n2 int) int {
	if n1 > n2 {
		return 1
	}
	if n1 == n2 {
		return 0
	}
	return -1
}

func IsSorted(f func(a, b int) int, a []int) bool {
	isAsc := true
	isDesc := true

	for i := 0; i < len(a)-1; i++ {
		if f(a[i], a[i+1]) > 0 { // a[i] > a[i+1]
			isAsc = false
		}
		if f(a[i], a[i+1]) < 0 { // a[i] < a[i+1]
			isDesc = false
		}
	}
	return isAsc || isDesc
}
