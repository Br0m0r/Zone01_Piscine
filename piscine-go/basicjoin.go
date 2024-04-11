package piscine

// function that returns a concatenated string from the 'strings' passed as arguments.
func BasicJoin(elems []string) string {
	var elems2 string
	for _, r := range elems {
		elems2 += r
	}
	return elems2
}
