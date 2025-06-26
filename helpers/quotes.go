package helpers

import (
	"unicode"
)

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
			// Check if there's a matching non-contraction quote ahead
			paired := false
			for j := i + 1; j < len(input); j++ {
				if input[j] == '\'' {
					// skip if that one is part of a contraction
					if !(j > 0 && j+1 < len(input) && unicode.IsLetter(input[j-1]) && unicode.IsLetter(input[j+1])) {
						paired = true
						break
					}
				}
			}

			if !paired {
				// no closing quote: append as is
				output = append(output, ch)
				continue
			}

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
