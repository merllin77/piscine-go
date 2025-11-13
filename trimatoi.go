package main

import "fmt"

func TrimAtoi(s string) int {
	number := 0
	chartoint := 0
	isnegative := false

	for i := 0; i < len(s); i++ { // begin checking each char in string "s"
		dgt := s[i] // hold the "byte" of the position [i] of the string (s) in dgt

		if dgt == '-' {
			isnegative = true
		}

		if dgt >= '0' && dgt <= '9' { // if we find a "number" (!as string!)
			chartoint = int(dgt - '0')     // convert char (eg '5' from string) to integer (eg 5 (number))
			number = number*10 + chartoint // adds the converted digits one next to other L -> R (from left to right)
		}

		if isnegative == true {
			number = number * -number
		}
	}
	return number
}

func main() {
	fmt.Println(TrimAtoi("12345"))
	fmt.Println(TrimAtoi("str123ing45"))
	fmt.Println(TrimAtoi("012 345"))
	fmt.Println(TrimAtoi("Hello World!"))
	fmt.Println(TrimAtoi("sd+x1fa2W3s4"))
	fmt.Println(TrimAtoi("sd-x1fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1-fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1+fa2W3s4"))
}
