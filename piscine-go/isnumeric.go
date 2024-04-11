package piscine

// function that returns true if the string passed as a parameter contains only numerical characters, otherwise, returns false.
func IsNumeric(s string) bool {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			count += 1
		}
	}
	return count == len(s)
}
