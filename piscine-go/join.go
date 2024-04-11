package piscine

//  function that returns the concatenation of all the strings of a slice of strings separated by the separator passed as the argument sep.
func Join(strs []string, sep string) string {
	var strs2 string

	for i, r := range strs {
		if i != len(strs)-1 {
			strs2 += r + sep
		} else {
			strs2 += r
		}
	}
	return strs2
}
