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
		l := helpers.Clear(line)
		// handling flags :
		l = helpers.HandlingFlags(l)
		// punctuations
		h := strings.Join(l, "")
		h = helpers.Punctuations2(h)
		h = helpers.Punctuations1(h)
		// quotes
		quots := func(s string) string {
			s = helpers.Quots1(s)
			s = helpers.Quots2(s)
			return s
		}
		h = quots(h)
		// a to an
		h = helpers.AtoAn(h)

		if len(h) > 0 {
			h = h[:len(h)-1]
		}
		output = append(output, h)
	}
	return strings.Join(output, "\n")
}
