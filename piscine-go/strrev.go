package piscine

// reverses a string,will return the reversed string.
func StrRev(s string) string {
	RuneArray := []rune(s)
	EmptyArray := make([]rune, len(s))
	for i := 0; i < len(s); i++ {
		EmptyArray[i] = RuneArray[len(s)-i-1]
	}
	final := string(EmptyArray)
	return final
}
