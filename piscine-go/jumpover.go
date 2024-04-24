package piscine

func JumpOver(str string) string {
	var new []string
	var new2 string

	if str == "" || len(str) < 2 {
		return "\n"
	}
	for i := 2; i < len(str); i += 3 {
		new = append(new, string(str[i]))
	}
	for _, r := range new {
		new2 += r
	}
	return new2 + "\n"
}
