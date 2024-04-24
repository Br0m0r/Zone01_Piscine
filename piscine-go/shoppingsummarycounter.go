package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	// Initialize the summary map
	summary := make(map[string]int)
	// Check if the input string is empty
	if str == "" {
		summary[""] = 1
		return summary
	}
	// Variables to track the start index and end index of each item
	var start, end int
	// Iterate over the summary string
	for i := 0; i <= len(str); i++ {
		if i == len(str) || str[i] == ' ' {
			// Found a space or reached the end of the string, extract the item
			end = i
			item := str[start:end]
			// Add the item to the summary map and increment its count
			summary[item]++
			// Update the start index for the next item
			start = i + 1
		}
	}
	return summary
}

func WhiteSpace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}
