package piscine

// iterative function that returns the factorial of the int passed as parameter.
func IterativeFactorial(nb int) int {
	if nb == 0 {
		return 1
	}
	if nb > 20 || nb < -20 {
		return 0
	}

	res := 1
	for i := 1; i <= nb; i++ {
		res *= i
		if res > 9223372036854775807 { // max int64 value
			return 0
		}
	}

	return res
}
