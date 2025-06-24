package helpers

import (
	"unicode"
)

func Quotes(s string) string {
	input := []rune(s)

	finished := false
	for !finished {
		open := 0
		finished = true
		for i := 0; i < len(input); i++ {

			if i-1>= 0 && unicode.IsLetter(input[i-1]) && i+1 < len(input) && input[i] == '\'' && unicode.IsLetter(input[i+1]) {
				continue
			}

			if input[i] == '\'' {
				switch open {
				case 0: // remove space after opening quots 
					if i+1 < len(input) && input[i+1] == ' ' {
						input = append(input[:i+1], input[i+2:]...)
						finished = false
					}
					open = 1
				case 1: // remove space before closing quots
					if i-1 >= 0 && input[i-1] == ' ' {
						input = append(input[:i-1], input[i:]...)
						finished = false
						i--
					}
					open = 0
				}
			}
		}
		if finished {
			break
		}
	}
	return string(input)
}
