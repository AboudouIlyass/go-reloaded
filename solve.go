package goreloaded

import (
	"strings"

	"goreloaded/helpers"
)

// process the file line by line
func Solve(input string) string {
	output := []string{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		h := helpers.Flags(line)
		h = helpers.Punctuations(h)
		quots := func(s string) string {
			s = helpers.Quots1(s)
			s = helpers.Quots2(s)
			return s
		}
		h = quots(h)
		h = helpers.AtoAn(h)
		output = append(output, h)
	}
	return strings.Join(output, "\n")
}
