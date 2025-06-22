package helpers

import (
	"strings"
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
	for i := 0; i < len(input); i++ {
		switch {
		case i+4 <= len(input) && input[i:i+4] == "(up)":
			out = append(out, "(up)")
			isflag["(up)"] = Flag{true, "(up)", ""}
			i += 3
		case i+5 <= len(input) && input[i:i+5] == "(low)":
			out = append(out, "(low)")
			isflag["(low)"] = Flag{true, "(low)", ""}
			i += 4
		case i+5 <= len(input) && input[i:i+5] == "(cap)":
			out = append(out, "(cap)")
			isflag["(cap)"] = Flag{true, "(cap)", ""}
			i += 4
		case i+5 <= len(input) && input[i:i+5] == "(bin)":
			out = append(out, "(bin)")
			isflag["(bin)"] = Flag{true, "(bin)", ""}
			i += 4
		case i+5 <= len(input) && input[i:i+5] == "(hex)":
			out = append(out, "(hex)")
			isflag["(hex)"] = Flag{true, "(hex)", ""}
			i += 4
		case IsPunc(input[i]):
			out = append(out, string(input[i]))
		default:
			if i+5 <= len(input) {
				up := input[i+5:]
				pos1 := strings.Index(up, ")")
				if pos1 != -1 {
					end := i + 5 + pos1
					if end+1 <= len(input) && end >= i+5 && input[i:i+5] == "(up, " && AreDigits(input[i+5:end]) {
						out = append(out, input[i:end+1])
						isflag[input[i:end+1]] = Flag{true, "(up, n)", input[i+5 : end]}
						i = end
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
						i = end
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
						i = end
						continue
					}
				}
			}
			out = append(out, string(input[i]))
		}
	}

	for i := 0; i < len(out)-1; i++ {
		if out[i] == " " && len(out[i+1]) > 0 && len(out[i+1]) == 1 && IsPunc(out[i+1][0]) {
			out[i], out[i+1] = out[i+1], out[i]
			i++
		}
	}

	// clean strings
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

func AreDigits(s string) bool {
	for _, d := range s {
		if d < '0' || d > '9' || !(s[0] == '+' || s[0] == '-' || s[0]>= '0' && s[0]<='9') {
			return false
		}
	}
	return true
}

func IsPunc(s byte) bool {
	return (s == '.' || s == ',' || s == '!' || s == '?' || s == ':' || s == ';')
}
