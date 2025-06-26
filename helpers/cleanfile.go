package helpers

import (
	"strings"
	"unicode/utf8"
)

type Flag struct {
	IsFlag bool
	f      string
	data   string
}

var isflag = map[string]Flag{}

func CleanLine(input string) []string {
	out := []string{}

	// detect flags
	for i := 0; i < len(input); {
		switch {
		case i+4 <= len(input) && input[i:i+4] == "(up)":
			out = append(out, "(up)")
			isflag["(up)"] = Flag{true, "(up)", ""}
			i += 4
		case i+5 <= len(input) && input[i:i+5] == "(low)":
			out = append(out, "(low)")
			isflag["(low)"] = Flag{true, "(low)", ""}
			i += 5
		case i+5 <= len(input) && input[i:i+5] == "(cap)":
			out = append(out, "(cap)")
			isflag["(cap)"] = Flag{true, "(cap)", ""}
			i += 5
		case i+5 <= len(input) && input[i:i+5] == "(bin)":
			out = append(out, "(bin)")
			isflag["(bin)"] = Flag{true, "(bin)", ""}
			i += 5
		case i+5 <= len(input) && input[i:i+5] == "(hex)":
			out = append(out, "(hex)")
			isflag["(hex)"] = Flag{true, "(hex)", ""}
			i += 5
		case IsPunc(rune(input[i])):
			out = append(out, string(input[i]))
			i++
		default:
			if i+5 <= len(input) {
				up := input[i+5:]
				pos1 := strings.Index(up, ")")
				if pos1 != -1 {
					end := i + 5 + pos1
					if end+1 <= len(input) && end >= i+5 && input[i:i+5] == "(up, " && AreDigits(input[i+5:end]) {
						out = append(out, input[i:end+1])
						isflag[input[i:end+1]] = Flag{true, "(up, n)", input[i+5 : end]}
						i = end + 1
						continue
					}
				}
			}
			if i+6 <= len(input) {
				cap := input[i+6:]
				pos1 := strings.Index(cap, ")")
				if pos1 != -1 {
					end := i + 6 + pos1
					if end+1 <= len(input) && end >= i+6 && input[i:i+6] == "(cap, " && AreDigits(input[i+6:end]) {
						out = append(out, input[i:end+1])
						isflag[input[i:end+1]] = Flag{true, "(cap, n)", input[i+6 : end]}
						i = end + 1
						continue
					}
				}
			}
			if i+6 <= len(input) {
				low := input[i+6:]
				pos1 := strings.Index(low, ")")
				if pos1 != -1 {
					end := i + 6 + pos1
					if end+1 <= len(input) && end >= i+6 && input[i:i+6] == "(low, " && AreDigits(input[i+6:end]) {
						out = append(out, input[i:end+1])
						isflag[input[i:end+1]] = Flag{true, "(low, n)", input[i+6 : end]}

						i = end + 1
						continue
					}
				}
			}
			r, size := utf8.DecodeRuneInString(input[i:])
			out = append(out, string(r))
			i += size
		}
	}

	// add new spaces with flags if needed
	for i := 0; i < len(out); i++ {
		isf := isflag[out[i]]
		if isf.IsFlag {
			if i > 0 && out[i-1][len(out[i-1])-1] != ' ' {
				out = append(out[:i], append([]string{" "}, out[i:]...)...)
				i++
			}
			if i+1 < len(out) && out[i+1][0] != ' ' {
				out = append(out[:i+1], append([]string{" "}, out[i+1:]...)...)
			}
		}

	}
	out = append(out, " ")

	// build a clean slice
	output := []string{}
	temp := ""

	for i := 0; i < len(out); i++ {
		isf := isflag[out[i]]
		if isf.IsFlag {
			if len(temp) != 0 {
				output = append(output, temp)
				temp = ""
			}
			output = append(output, out[i])
		} else if out[i] == " " {
			if len(temp) != 0 {
				output = append(output, temp)
				temp = ""
			}
			output = append(output, " ")
		} else {
			temp += out[i]
		}
	}
	if len(temp) != 0 {
		output = append(output, temp)
	}
	return output
}
