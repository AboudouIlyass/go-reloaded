package goreloaded

import (
	"strings"

	"goreloaded/helpers"
)

func Solve(input string) string {
	output := []string{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		// handling flag
		h := helpers.Flags(line)
		// punctuations
		h = helpers.Punctuations(h)
		// quotes
		quots := func(s string) string {
			s = helpers.Quots1(s)
			s = helpers.Quots2(s)
			return s
		}
		h = quots(h)
		// a to an
		h = helpers.AtoAn(h)
		output = append(output, h)
	}
	return strings.Join(output, "\n")
}
