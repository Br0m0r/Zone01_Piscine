package piscine

func Abort(a, b, c, d, e int) int {
	temp := []int{a, b, c, d, e}
	for i := 0; i < len(temp); i++ {
		for j := 0; j < len(temp)-i-1; j++ {
			if temp[j] > temp[j+1] {
				temp[j], temp[j+1] = temp[j+1], temp[j]
			}
		}
	}
	return temp[2]
}
