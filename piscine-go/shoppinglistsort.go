package piscine

func ShoppingListSort(slice []string) []string {
	for i := 0; i < len(slice); i++ {
		swapped := false
		for j := 0; j < len(slice)-i-1; j++ {
			if len(slice[j]) > len(slice[j+1]) {
				slice[j], slice[j+1] = slice[j+1], slice[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return slice
}
