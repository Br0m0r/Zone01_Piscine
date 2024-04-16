package piscine

func CountIf(f func(string) bool, tab []string) int {
	counter := 0
	for _, r := range tab {
		if f(r) {
			counter++
		}
	}
	return counter
}
