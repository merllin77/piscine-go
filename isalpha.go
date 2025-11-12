package main

import "fmt"

func IsAlpha(s string) bool {
	for _, r := range s {
		if (r < '0' && r > '9') || r < 'A' {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsAlpha("Hello! How are you?"))
	fmt.Println(IsAlpha("HelloHowareyou"))
	fmt.Println(IsAlpha("What's this 4?"))
	fmt.Println(IsAlpha("Whatsthis4"))
}
