package piscine

// function that counts the letters of a string and returns the count.
func AlphaCount(s string) int {
	count := 0
	for _, i := range s {
		if i >= 'a' && i <= 'z' {
			count++
		}
		if i >= 'A' && i <= 'Z' {
			count++
		}
	}
	return count
}
