package goreloaded

import (
	"strings"

	"goreloaded/helpers"
)

func Solve(input string) string {
	output := []string{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {

		// clean line
		l := helpers.CleanLine(line)

		// handling flags :
		l = helpers.HandlingFlags(l)

		// punctuations
		h := strings.Join(l, "")
		h = helpers.Punctuations1(h)
		h = helpers.Punctuations2(h)

		// quotes
		h = helpers.Quotes(h)

		// a to an
		h = helpers.AtoAn(h)
		
		if len(h) > 0 {
			h = h[:len(h)-1]
		}

		output = append(output, h)
	}
	ret := strings.Join(output, "\n")

	return ret
}
