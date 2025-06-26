package goreloaded

import (
	"goreloaded/helpers"
	"strings"
)

func Solve(input string) string {
	output := []string{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		// handling flag
		h := helpers.Clear(line)
		// punctuations
		h = helpers.Punctuations2(h)
		h = helpers.Punctuations1(h)
		// quotes
		quots := func(s string) string {
			s = helpers.Quots1(s)
			s = helpers.Quots2(s)
			return s
		}
		// a to an
		h = helpers.AtoAn(h)
		h = quots(h)
		output = append(output, h)
	}
	return strings.Join(output, "\n")
}
