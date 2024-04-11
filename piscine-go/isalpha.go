package piscine

//  function that returns true if the string passed as the parameter only contains alphanumerical characters or is empty, otherwise, and returns false.
func IsAlpha(s string) bool {
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			count += 1
		}
		if s[i] >= 'A' && s[i] <= 'Z' {
			count += 1
		}
		if s[i] >= '0' && s[i] <= '9' {
			count += 1
		}

	}
	return count == len(s)
}
