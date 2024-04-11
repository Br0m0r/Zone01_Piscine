package piscine

// function that returns the last rune of a string.
func LastRune(s string) rune {
	sen := []rune(s)
	return sen[len(s)-1]
}
