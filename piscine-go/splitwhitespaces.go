package piscine

func SplitWhiteSpaces(s string) []string {
	var words []string
	start := 0
	inWord := false

	for i, ch := range s {
		if isWhiteSpace(ch) {
			if inWord {
				words = append(words, s[start:i])
				inWord = false
			}
		} else if !inWord {
			start = i
			inWord = true
		}
	}
	if inWord { // to handle the last word if the string doesn't end with a whitespace
		words = append(words, s[start:])
	}

	return words
}

func isWhiteSpace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n'
}
