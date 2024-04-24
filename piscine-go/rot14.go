package piscine

func Rot14(s string) string {
	ab := "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"
	AB := "ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var STR []rune

	for _, r := range s {
		if ALB(r) {
			if r >= 'a' && r <= 'z' {
				for i, s := range ab {
					if r == rune(s) {
						STR = append(STR, rune(ab[i+14]))
						break
					}
				}
			} else if r >= 'A' && r <= 'Z' {
				for i, s := range AB {
					if r == rune(s) {
						STR = append(STR, rune(AB[i+14]))
						break
					}
				}
			}
		} else {
			STR = append(STR, r)
		}
	}
	return string(STR)
}

func ALB(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}
