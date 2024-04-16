package piscine

func Any(f func(string) bool, a []string) bool {
	counter := 0
	for _, r := range a {
		if f(r) {
			counter++
		}
	}
	return counter > 0
}
