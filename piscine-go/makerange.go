package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil // Return nil if min is greater than or equal to max
	}
	length := max - min
	rangeSlice := make([]int, length) // Create a slice with the exact needed size
	for i := 0; i < length; i++ {
		rangeSlice[i] = min + i // Populate the slice with the range of integers
	}
	return rangeSlice
}
