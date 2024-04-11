package piscine

//  function that returns true if the string passed as a parameter contains only printable characters, otherwise, returns false.
func IsPrintable(s string) bool {
	for _, c := range s {
		if c < 32 || c > 126 {
			return false
		}
	}
	return true
}
