package piscine

func ConcatParams(args []string) string {
	if len(args) == 0 {
		return "" // Return an empty string if the input slice is empty
	}
	result := args[0]              // Start with the first element to avoid a newline at the start
	for _, arg := range args[1:] { // Start from the second element
		result += "\n" + arg // Concatenate each subsequent string with a newline in front
	}
	return result
}
