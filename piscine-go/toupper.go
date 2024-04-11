package piscine

// function that capitalizes each letter in a string.
func ToUpper(s string) string {
	s2 := make([]rune, len(s))
	for i, l := range s {
		if l >= 'a' && l <= 'z' {
			s2[i] = l - 32
		} else {
			s2[i] = l
		}
	}
	return string(s2)
}
