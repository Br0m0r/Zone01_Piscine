package piscine

func CompareNum(a, b int) int {
	if a > b {
		return 1
	} else if a < b {
		return -1
	}
	return 0
}

// IsSorted checks if the array is sorted in either ascending or descending order.
func IsSorted(f func(a, b int) int, arr []int) bool {
	isAscending := true
	isDescending := true

	for i := 0; i < len(arr)-1; i++ {
		comparisonResult := f(arr[i], arr[i+1])
		if comparisonResult > 0 {
			isAscending = false
		}
		if comparisonResult < 0 {
			isDescending = false
		}
		if !isAscending && !isDescending {
			return false
		}
	}
	// If at least one sorting condition (ascending or descending) is true, return true.
	return true
}
