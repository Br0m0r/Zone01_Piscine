package piscine

func DescendAppendRange(max, min int) []int {
	if min > max || min == 0 && max == 0 {
		return []int{}
	}
	numbers := []int{}
	for i := max; i > min; i-- {
		numbers = append(numbers, i)
	}
	return numbers
}
