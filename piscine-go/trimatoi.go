package piscine

// function that transforms numbers within a string, into an int.If the - sign is encountered before any number it should determine the sign of the returned int.This function should only return an int. In the case of an invalid input, the function should return 0.
func TrimAtoi(s string) int {
	if s == "" {
		return 0
	}
	result := 0
	sign := 1
	foundDigit := false // indicates when digit processing starts

	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			foundDigit = true // set true(we ve fouind a dig)
			result = result*10 + int(s[i]-'0')
		} else if !foundDigit { // Only check for sign before any digit is found
			if s[i] == '-' {
				sign = -1
			} else if s[i] == '+' {
				sign = 1
			}
		} else {
			// Once a digit has been found, ignore all characters that are not digits
			continue
		}
	}
	return result * sign
}
