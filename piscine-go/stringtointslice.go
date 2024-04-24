package piscine

func StringToIntSlice(str string) []int {
	var new []int
	for _, r := range str {
		new = append(new, int(r))
	}
	return new
}
