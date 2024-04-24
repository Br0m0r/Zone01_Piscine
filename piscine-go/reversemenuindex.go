package piscine

func ReverseMenuIndex(menu []string) []string {
	new := []string{}
	for i := len(menu) - 1; i >= 0; i-- {
		new = append(new, menu[i])
	}
	return new
}
