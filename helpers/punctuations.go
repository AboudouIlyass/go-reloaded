package helpers

// remove spaces before puntuations
func Punctuations1(s string) string {
	input := []rune(s)

	for i := 1; i < len(input); i++ {
		if IsPunc(input[i]) {
			for i > 0 && input[i-1] == ' ' {
				input = append(input[:i-1], input[i:]...)
				i--
			}
		}
	}
	return string(input)
}

// add spaces after punctuations id needed
func Punctuations2(s string) string {
	input := []rune(s)

	for i := 0; i < len(input); i++ {
		if IsPunc(input[i]) {
			if i+1 < len(input) && input[i+1] != ' ' {
				input = append(input[:i+1], append([]rune{' '}, input[i+1:]...)...)
			}
		}
	}
	return string(input)
}
