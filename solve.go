package goreloaded

import (
	"strings"

	"goreloaded/helpers"
)

func Solve(input string) string {
	output := []string{}
	lines := strings.Split(input, "\n")
	for _, line := range lines {

		// clean line :
		l := helpers.CleanLine(line)
		// handling flags :
		l = helpers.HandlingFlags(l)
		output = append(output, strings.Join(l, ""))

	}

	return strings.Join(output, "\n")
}
