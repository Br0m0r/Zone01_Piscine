package piscine

// function that returns true if the string passed as the parameter contains only lowercase characters, otherwise, returns false.
func IsLower(s string) bool {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			count += 1
		}
	}
	return count == len(s)
}
