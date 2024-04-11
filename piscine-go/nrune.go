package piscine

// function that returns the nth rune of a string. If not possible, it returns 0.
func NRune(s string, n int) rune {
	sen := []rune(s)
	if n <= 0 || n > len(sen) {
		return 0
	}
	return sen[n-1]
}
