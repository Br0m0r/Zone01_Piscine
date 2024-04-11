package piscine

// function that capitalizes the first letter of each word and lowercases the rest.A word is a sequence of alphanumeric characters.
func Capitalize(s string) string {
	slice := []rune(s)
	// if the next letter should be uppercased.
	shouldCapitalizeNext := true

	for i, r := range slice {
		// Check if  rune is alphabetic
		isAlpha := (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')

		if shouldCapitalizeNext && isAlpha {
			// If the rune is alphabetic and should be capitalized, uppercase it
			if r >= 'a' && r <= 'z' {
				slice[i] -= 32
			}
			shouldCapitalizeNext = false // Reset the flag as we've capitalized a letter
		} else if !isAlpha {
			// If the rune is non-alphanumeric, the next alphabetic character should be uppercased
			shouldCapitalizeNext = true
		} else if r >= 'A' && r <= 'Z' {
			// Lowercase any alphabetic character that shouldn't be capitalized
			slice[i] += 32
		}
	}

	return string(slice)
}
