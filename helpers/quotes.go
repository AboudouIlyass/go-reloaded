package helpers

import (
	"unicode"
)

// add space to close the quots if needed
func Quots1(s string) string {
	input := []rune(s)
	open := false
	pos := -1
	for i := 0; i < len(input); i++ {
		if input[i] == '\'' {
			if i > 0 && i+1 < len(input) && unicode.IsLetter(input[i-1]) && unicode.IsLetter(input[i+1]) {
				continue
			}
			if !open {
				pos = i
				open = true
			} else if open {
				if pos-1 >= 0 && input[pos-1] != ' ' {
					input = MyAdd(input, pos, ' ')
					i++
				}

				if i+1 < len(input)-1 && input[i+1] != ' ' {
					input = MyAdd(input, i+1, ' ')
					i++
				}
				open = false
			}
		}
	}
	return string(input)
}

// remove extra spaces to close the quots if needed
func Quots2(s string) string {
	input := []rune(s)
	pos := -1

	finished := false
	for !finished {
		open := false
		finished = true
		for i := 0; i < len(input); i++ {
			if input[i] == '\'' {
				if i > 0 && i+1 < len(input) && unicode.IsLetter(input[i-1]) && unicode.IsLetter(input[i+1]) {
					continue
				}
				if !open {
					pos = i
					open = true
				} else if open {
					if pos+1 < len(input)-1 && input[pos+1] == ' ' {
						input = MyDelete(input, pos+1)
						finished = false
					}
					if i-1 >= 0 && input[i-1] == ' ' {
						input = MyDelete(input, i-1)
						finished = false
					}
					open = false
				}
			}
		}
		if finished {
			break
		}
	}
	return string(input)
}

