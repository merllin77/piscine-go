package student

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)
	word := ""

	for _, r := range str {
		if r != ' ' {
			word += string(r)
		} else {
			summary[word]++ // always increment, even if word is ""
			word = ""
		}
	}

	// add last word
	summary[word]++

	return summary
}
