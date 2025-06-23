package helpers

func Punctuations1(s string) string {
	input := []rune(s)

	for i := 0; i < len(input); i++ {
		if input[i] == ' ' && i+1 < len(input) && IsPunc(input[i+1]) {
			input = append(input[:i], input[i+1:]...)
		}
	}

	return string(input)
}

func Punctuations2(s string) string {
	input := []rune(s)

	for i := 0; i < len(input); i++ {
		if IsPunc(input[i]) && i+1 < len(input) && !IsPunc(input[i+1]) && input[i+1] != ' ' {
			input = append(input[:i+1], append([]rune{' '}, input[i+1:]...)...)
			i++
		}
	}
	return string(input)
}
