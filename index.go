package main

import (
	"fmt"
)

func Index(s string, toFind string) int {

	totallen := len(s) - len(toFind)

	for i := 0; i <= totallen; i++ {
		for j := 0; j < len(toFind); j++ {
			if s[i+j] != toFind[j] {
				break
			} else {
				return i
			}
		}
	}
	return 0
}

func main() {
	fmt.Println(Index("Hello!", "l"))
	fmt.Println(Index("Salut!", "alu"))
	fmt.Println(Index("Ola!", "hOl"))
}
