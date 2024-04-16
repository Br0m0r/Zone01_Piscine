package piscine

func Map(f func(int) bool, a []int) []bool {
	boolslice := make([]bool, len(a))
	for i, r := range a {
		boolslice[i] = f(r)
	}
	return boolslice
}
