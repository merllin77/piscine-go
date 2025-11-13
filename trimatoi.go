package student

func TrimAtoi(s string) int {
	number := 0
	chartoint := 0
	foundneg := false
	founddigit := false

	for i := 0; i < len(s); i++ { // begin checking each char in string "s"
		dgt := s[i] // hold the "byte" of the position [i] of the string (s) in dgt
		if dgt == '-' && founddigit == false {
			foundneg = true
		} else if dgt >= '0' && dgt <= '9' { // if we find a "number" (!as a string!)
			chartoint = int(dgt - '0')     // convert char (eg '5' from string) to integer (eg 5 (number))
			number = number*10 + chartoint // adds the converted digits one next to other L -> R (from left to right)
			founddigit = true
		}
	}
	if foundneg == true {
		number = -number
	}
	return number
}
