package piscine

// transforms a number represented as a string in a number represented as an int.
func Atoi(s string) int {
	if s == "" {
		return 0
	}
	result := 0
	sign := 1
	startIndex := 0

	if s[0] == '-' {
		sign = -1
		startIndex = 1
	} else if s[0] == '+' {
		startIndex = 1
	}

	for i := startIndex; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0
		}
		result = result*10 + int(s[i]-'0')
	}
	return result * sign
}
