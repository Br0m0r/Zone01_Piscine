package piscine

func Unmatch(a []int) int {
	var comp []int
	for _, r := range a {
		comp = append(comp, r)
	}

	for _, r := range a {
		count := 0
		for i := 0; i < len(comp); i++ {
			if comp[i] == r {
				count++
			}
		}
		if count%2 != 0 {
			return r
		}

	}
	return -1
}
