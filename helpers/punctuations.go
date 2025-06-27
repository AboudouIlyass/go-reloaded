package helpers

// add and remove spaces before and after punctuation if needed
func Punctuations(s string) string {
	input := []rune(s)
	for i := 0; i < len(input); i++ {
		// add
		if IsPunc(input[i]) {
			if i+1 < len(input) && input[i+1] != ' ' && !IsPunc(input[i+1]) {
				input = MyAdd(input, i+1, ' ')
			}
		}
	}

	for i := 1; i < len(input); i++ {
		// remove
		if IsPunc(input[i]) {
			for i > 0 && input[i-1] == ' ' {
				input = MyDelete(input, i-1)
				i--
			}
		}
	}
	return string(input)
}
