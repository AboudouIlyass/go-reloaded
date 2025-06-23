package goreloaded

import (
	"fmt"
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
		fmt.Println(h)
		h = helpers.Punctuations2(h)

		// quotes
		h = helpers.Quotes(h)

		output = append(output, h)

	}

	return strings.Join(output, "\n")
}
