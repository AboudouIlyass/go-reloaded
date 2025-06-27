package helpers

func Punctuations(s string) string {
	// add spaces after punctuations id needed
	input := []rune(s)
	for i := 0; i < len(input); i++ {
		if IsPunc(input[i]) {
			if i+1 < len(input) && input[i+1] != ' ' && !IsPunc(input[i+1]) {
				input = MyAdd(input, i+1, ' ')
			}
		}
	}

	// remove spaces before puntuations if needed
	for i := 1; i < len(input); i++ {
		if IsPunc(input[i]) {
			for i > 0 && input[i-1] == ' ' {
				input = MyDelete(input, i-1)
				i--
			}
		}
	}
	return string(input)
}
