package piscine

//  function that returns the first prime number that is equal or superior to the int passed as parameter.
func FindNextPrime(n int) int {
	for !IsPrime(n) {
		n++
	}
	return n
}
