package helpers

import "unicode"

// remove spaces before puntuations
func Punctuations1(s string) string {
	input := []rune(s)
	for i := 1; i < len(input); i++ {
		if i-1 >= 0 && i+1 < len(input)-1 && (unicode.IsLetter(input[i-1]) || unicode.IsDigit(input[i-1])) &&
			(unicode.IsLetter(input[i+1]) || unicode.IsDigit(input[i+1])) && IsPunc(input[i]) {
			continue
		}
		if IsPunc(input[i]) {
			for i > 0 && input[i-1] == ' ' {
				input = MyDelete(input, i-1)
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
		if i-1 >= 0 && i+1 < len(input)-1 && (unicode.IsLetter(input[i-1]) || unicode.IsDigit(input[i-1])) &&
			(unicode.IsLetter(input[i+1]) || unicode.IsDigit(input[i+1])) && IsPunc(input[i]) {
			continue
		}
		if IsPunc(input[i]) {
			if i+1 < len(input) && input[i+1] != ' ' && !IsPunc(input[i+1]) {
				input = MyAdd(input, i+1, ' ')
			}
		}
	}
	return string(input)
}
