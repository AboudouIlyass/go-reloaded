package helpers

import (
	"fmt"
)

func Quotes(s string) string {
	input := []rune(s)
	open := 0
	finished := false
	fmt.Println(s)
	for !finished {
		finished = true
		for i := 0; i < len(input); i++ {
			if input[i] == '\'' {
				switch open {
				case 0:
					if i+1 < len(input) && input[i+1] == ' ' {
						input = append(input[:i+1], input[i+2:]...)
						finished = false
					}
					open = 1
				case 1:
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
