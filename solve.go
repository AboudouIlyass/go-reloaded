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
		fmt.Println(l)

		// handling flags :
		l = helpers.HandlingFlags(l)

		// punctuations
		h := strings.Join(l, "")
		h = helpers.Punctuations1(h)
		h = helpers.Punctuations2(h)
		
		// quotes
		h = helpers.Quotes(h)
		h = helpers.AtoAn(h)
		output = append(output, h)

	}
	ret := strings.Join(output, "\n")

	if len(ret) > 0{
		ret = ret[:len(ret)-1]
	}
	return ret
}
