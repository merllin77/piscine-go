package main

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

//func IsSorted(f func(a, b int) int, a []int) bool {
//}

func main() {
	a1 := []int{11, 10, 12, 13, 14, 15}
	// a2 := []int{0, 2, 1, 3}

	// result1 := IsSorted(InOrder, a1)
	// result2 := IsSorted(InOrder, a2)

	// fmt.Println(result1)
	// fmt.Println(result2)

	fmt.Println(InOrder, a1)
	// for i := 0; i < len(a1)-1; i++ {
	// 	if a1[i] < a1[i+1] {
	// 		fmt.Println(a1[i])
	// 	} else {
	// 		fmt.Println("Is not sorted")
	// 		return
	// 	}
	// }
}
