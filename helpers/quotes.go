package helpers

import (
	"unicode"
)

func Quotss1(s string) string {
    input := []rune(s)
    // Step 1: collect positions of real quotes (excluding contraction apostrophes)
    var quotePos []int
    for i, ch := range input {
        if ch == '\'' {
            if i > 0 && i+1 < len(input) && unicode.IsLetter(input[i-1]) && unicode.IsLetter(input[i+1]) {
                continue // skip apostrophes in contractions
            }
            quotePos = append(quotePos, i)
        }
    }
    // Pair up and drop unmatched
    if len(quotePos)%2 == 1 {
        quotePos = quotePos[:len(quotePos)-1]
    }
    paired := make(map[int]bool)
    for i := 0; i < len(quotePos); i += 2 {
        paired[quotePos[i]] = true   // opening
        paired[quotePos[i+1]] = true // closing
    }

    // Step 2: build output, adding spaces only around paired quotes
    var output []rune
    for i, ch := range input {
        if ch == '\'' && paired[i] {
            // Check if this is opening (next paired is after) or closing
            isOpening := false
            // count how many paired before this
            count := 0
            for _, pos := range quotePos {
                if pos < i {
                    count++
                }
            }
            if count%2 == 0 {
                isOpening = true
            }
            // opening: space before and after
            if isOpening {
                if len(output) > 0 && output[len(output)-1] != ' ' {
                    output = append(output, ' ')
                }
                output = append(output, '\'')
                if i+1 < len(input) && input[i+1] != ' ' {
                    output = append(output, ' ')
                }
            } else {
                // closing: space before and after
                if len(output) > 0 && output[len(output)-1] != ' ' {
                    output = append(output, ' ')
                }
                output = append(output, '\'')
                if i+1 < len(input) && input[i+1] != ' ' {
                    output = append(output, ' ')
                }
            }
        } else {
            // default: copy character
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
