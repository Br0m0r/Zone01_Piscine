package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	current := l.Head

	for current != nil {
		// Perform the desired action on the current node

		if comp(current.Data, ref) {
			return &current.Data
		}
		// Move to the next node
		current = current.Next
	}
	return nil
}
