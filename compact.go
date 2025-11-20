package main

import "fmt"

// Compact removes zero-value strings from the slice and returns the number of non-zero elements
func Compact(slice *[]string) int {
	if slice == nil || len(*slice) == 0 {
		return 0
	}

	nonZero := 0
	for _, str := range *slice {
		if str != "" {
			(*slice)[nonZero] = str
			nonZero++
		}
	}
	// Resize the slice to only keep non-zero elements
	*slice = (*slice)[:nonZero]

	return nonZero
}

func main() {
	s := []string{"a", "", "b", "", "c"}

	// Print original slice
	for _, v := range s {
		if v != "" {
			fmt.Println(v)
		}
	}

	count := Compact(&s)

	fmt.Println("Size after compacting:", count)
	for _, v := range s {
		fmt.Println(v)
	}
}
