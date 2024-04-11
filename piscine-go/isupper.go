package piscine

// function that returns true if the string passed as parameter contains only uppercase characters, otherwise, returns false.
func IsUpper(s string) bool {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			count += 1
		}
	}
	return count == len(s)
}
