package student

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)
	word := ""

	for _, r := range str {
		if r != ' ' {
			word = word + string(r)
		} else if word != "" {
			summary[word]++
			word = ""
		}
	}
	if word != "" {
		summary[word]++
	}
	return summary
}

// func main() {
// 	summary := "Burger Water Carrot Coffee Water Water Chips Carrot Carrot Burger Carrot Water"
// 	for index, element := range ShoppingSummaryCounter(summary) {
// 		fmt.Println(index, "=>", element)
// 	}
//}
