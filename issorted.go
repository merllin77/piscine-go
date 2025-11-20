package student

import "fmt"

func InOrder(n1 []int) int {
	for i := 0; i < len(n1)-1; i++ {
		if n1[i] < n1[i+1] {
			fmt.Println("SORTED !!")
			return 1
		}
	}
	fmt.Println("Not sorted")
	return -1
}
