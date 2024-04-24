package piscine

func Compact(slice *[]string) int {
	temp := (*slice)[:0]

	count := 0

	for _, value := range *slice {
		if value != "" {
			temp = append(temp, value)
			count++
		}
	}
	*slice = temp

	return count
}
