package helpers

import (
	"fmt"
	"unicode"
)

func Quotes(s string) string {
	// s = AddSpaceIfNeeded(s)
	input := []rune(s)

	finished := false
	for !finished {
		open := 0
		finished = true
		for i := 0; i < len(input); i++ {
			if i-1 >= 0 && unicode.IsLetter(input[i-1]) && i+1 < len(input) && input[i] == '\'' && unicode.IsLetter(input[i+1]) {
				continue
			}
			if input[i] == '\'' {
				switch open {
				case 0: // remove space after opening quots
					if i+1 < len(input)-1 && input[i+1] == ' ' {
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

func AddSpaceIfNeeded(s string) string {
	input := []rune(s)
	open := 0
	for i := 0; i < len(input); i++ {

		if i-1 >= 0 && unicode.IsLetter(input[i-1]) && i+1 < len(input) && input[i] == '\'' && unicode.IsLetter(input[i+1]) {
			continue
		}
		if input[i] == '\'' {
			switch open {
			case 0: // add space before opening quots
				if i-1 >= 0 && input[i-1] != ' ' {
					input = append(input[:i], append([]rune{' '}, input[i:]...)...)
				}
				open = 1
			case 1: // add space after closing quots
				if i+1 < len(input)-1 && input[i+1] != ' ' {
					input = append(input[:i+1], append([]rune{' '}, input[i+1:]...)...)
				}
				open = 0
			}
		}
	}
	return string(input)
}

func Quots(s string) string {
	a := Quotss1(s)
	// should do a to an here
	a = Quotss2(a)
	a = AtoAn(a)
	fmt.Println(a)
	return a
}

func Quotss1(s string) string {
	input := []rune(s)
	var output []rune

	for i, ch := range input {
		// Detect contraction quotes like don't, it's
		if ch == '\'' &&
			i > 0 && i+1 < len(input) &&
			unicode.IsLetter(input[i-1]) &&
			unicode.IsLetter(input[i+1]) {
			output = append(output, ch)
			continue
		}

		if ch == '\'' {
			// space before '
			if len(output) > 0 && output[len(output)-1] != ' ' {
				output = append(output, ' ')
			}
			// the quote itself
			output = append(output, '\'')
			// space after '
			if i+1 < len(input) && input[i+1] != ' ' {
				output = append(output, ' ')
			}
		} else {
			output = append(output, ch)
		}
	}
	return string(output)
}

func Quotss2(s string) string {
		input := []rune(s)
	quotePos := []int{}

	// collect positions of real (non-contraction) single quotes
	for i := 0; i < len(input); i++ {
		if input[i] == '\'' {
			if i > 0 && i+1 < len(input) && unicode.IsLetter(input[i-1]) && unicode.IsLetter(input[i+1]) {
				continue // skip contraction quotes
			}
			quotePos = append(quotePos, i)
		}
	}

	// only process complete quote pairs
	if len(quotePos)%2 != 0 {
		// Odd number of quotes — drop the last unmatched one
		quotePos = quotePos[:len(quotePos)-1]
	}

	// work backwards to avoid messing up indexes
	for i := len(quotePos) - 2; i >= 0; i -= 2 {
		open := quotePos[i]
		close := quotePos[i+1]

		// Remove space after opening quote
		if open+1 < len(input) && input[open+1] == ' ' {
			input = append(input[:open+1], input[open+2:]...)
			close-- // everything shifted left
		}

		// Remove space before closing quote
		if close-1 >= 0 && input[close-1] == ' ' {
			input = append(input[:close-1], input[close:]...)
		}
	}

	return string(input)
}
