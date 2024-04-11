package piscine

// Insert your comment here
func BasicAtoi(s string) int {
	res := 0
	for _, r := range s {
		if r >= '0' && r <= '9' {
			res = res*10 + int(r-'0')
		} else {
			return 0
		}
	}
	return res
}
