package piscine

// iterative function that returns the value of nb to the power of power.Negative powers will return 0. Overflows are not dealt with.
func IterativePower(nb int, power int) int {
	result := 1
	for i := 1; i <= power; i++ {
		result *= nb
	}
	if power < 0 {
		return 0
	}

	return result
}
