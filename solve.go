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
		h = helpers.AtoAn(h)

		c1 := strings.Fields(h)
		c2 := strings.Join(c1, " ")

		output = append(output, c2)

	}

	return strings.Join(output, "\n")
}
