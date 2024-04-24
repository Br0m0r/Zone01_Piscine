package piscine

func LoafOfBread(str string) string {
	// We'll use two variables to manage characters and skipping logic
	var result string
	count := 0
	skipNext := false
	filteredIndex := -1

	// Create a filtered string with no spaces
	filtered := ""
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			filtered += string(str[i])
		}
	}

	// Check if there are at least 5 non-space characters
	if len(filtered) < 5 {
		return "Invalid Output\n"
	}

	// Process each character in the filtered string
	for i := 0; i < len(filtered); i++ {
		if skipNext {
			skipNext = false
			continue
		}

		result += string(filtered[i])
		filteredIndex++
		count++

		// Check if we need to add a space and skip the next character
		if count == 5 {
			if filteredIndex < len(filtered)-1 { // Ensure there is a character to skip
				skipNext = true
			}
			count = 0
			if i != len(filtered)-1 { // Ensure this isn't the last group
				result += " "
			}
		}
	}

	result += "\n"
	return result
}
