package piscine

func Max(a []int) int {
	temp := 0
	for _, r := range a {
		if r > temp {
			temp = r
		}
	}
	return temp
}
