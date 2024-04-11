package piscine

// function that lower cases for each letter in a string.
func ToLower(s string) string {
	s2 := make([]rune, len(s))
	for i, l := range s {
		if l >= 'A' && l <= 'Z' {
			s2[i] = l + 32
		} else {
			s2[i] = l
		}
	}
	return string(s2)
}
