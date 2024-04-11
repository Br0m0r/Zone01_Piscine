package piscine

// Index returns the index of the first instance of substr in s, or -1 if substr is not present in s.
func Index(s string, toFind string) int {
	if len(s) == 0 {
		return 0
	}

	for i := 0; i <= len(s)-len(toFind); i++ {
		if toFind == s[i:i+len(toFind)] {
			return i
		}
	}
	return -1
}
