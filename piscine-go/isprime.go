package piscine

// Write a function that returns true if the int passed as parameter is a prime number. Otherwise it returns false.Checks for timeouts.
func IsPrime(nb int) bool {
	if nb < 2 || nb != 2 && nb%2 == 0 {
		return false
	}

	for i := 3; i*i <= nb; i += 2 {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
