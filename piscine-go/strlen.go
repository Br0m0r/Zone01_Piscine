package piscine

// counts the runes of a string and that returns that count.
func StrLen(s string) int {
	counter := 0
	RuneArray := []rune(s)
	for i := 0; i < len(RuneArray); i++ {
		counter += 1
	}
	return counter
}
