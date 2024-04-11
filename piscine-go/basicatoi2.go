package piscine

// Insert your comment here
func BasicAtoi2(s string) int {
	var num int
	for _, char := range s {
		if char >= '0' && char <= '9' {
			num = num*10 + int(char-'0')
		} else {
			return 0
		}
	}
	return num
}
