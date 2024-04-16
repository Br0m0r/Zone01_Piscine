package piscine

func AtoiBase(s string, base string) int {
	if !isValidBase(base) {
		return 0
	}

	baseLength := len(base)
	baseIndexMap := make(map[rune]int)
	for i, ch := range base {
		baseIndexMap[ch] = i
	}

	var number int
	for _, ch := range s {
		if index, exists := baseIndexMap[ch]; exists {
			number = number*baseLength + index
		} else {
			return 0 // If character is not in base, return 0
		}
	}

	return number
}

// isValidBase checks if the base string is valid according to the given rules.
func isValidBase(base string) bool {
	if len(base) < 2 || Contains(base, '+') || Contains(base, '-') {
		return false
	}
	charMap := make(map[rune]bool)
	for _, ch := range base {
		if charMap[ch] {
			return false // Duplicate character found
		}
		charMap[ch] = true
	}
	return true
}

// Contains checks if the character exists in the string.
func Contains(str string, char rune) bool {
	for _, ch := range str {
		if ch == char {
			return true
		}
	}
	return false
}
