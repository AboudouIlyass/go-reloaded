package helpers

import (
	"strconv"
	"strings"
)

func HandlingFlags(input []string) []string {
	for i := 0; i < len(input); i++ {
		fl := isflag[input[i]]
		if fl.IsFlag {
			switch fl.f {
			case "(low, n)":
				c := fl.data
				numStr := strings.TrimLeft(c, "0")
				if numStr > "9223372036854775807" {
					numStr = "9223372036854775807"
				}
				num, _ := strconv.Atoi(numStr)
				if i-1 >= 0 {
					for j := i - 1; j >= 0 && num > 0; j = j - 1 {
						if input[j][0] != ' ' {
							input[j] = strings.ToLower(input[j])
							num--
						}
					}
				}
				if i < len(input)-1 {
					input = append(input[:i], input[i+2:]...)
					i--
				}
			case "(cap, n)":
				c := fl.data
				numStr := strings.TrimLeft(c, "0")
				if numStr > "9223372036854775807" {
					numStr = "9223372036854775807"
				}
				num, _ := strconv.Atoi(numStr)
				if i-1 >= 0 {
					for j := i - 1; j >= 0 && num > 0; j = j - 1 {
						if input[j][0] != ' ' {
							input[j] = Capitalize(input[j])
							num--
						}
					}
				}
				if i < len(input)-1 {
					input = append(input[:i], input[i+2:]...)
					i--
				}
			case "(up, n)":
				c := fl.data
				numStr := strings.TrimLeft(c, "0")
				if numStr > "9223372036854775807" {
					numStr = "9223372036854775807"
				}
				num, _ := strconv.Atoi(numStr)
				if i-1 >= 0 {
					for j := i - 1; j >= 0 && num > 0; j = j - 1 {
						if input[j][0] != ' ' {
							input[j] = strings.ToUpper(input[j])
							num--
						}
					}
				}
				if i < len(input)-1 {
					input = append(input[:i], input[i+2:]...)
					i--
				}

			case "(hex)":
				if i >= 0 {
					for j := i - 1; j >= 0; j-- {
						if input[j][0] != ' ' {
							num, err := strconv.ParseInt(input[j], 16, 64)
							if err == nil {
								input[j] = strconv.Itoa(int(num))
							}
							break
						}
					}
				}
				if i < len(input)-1 {
					input = append(input[:i], input[i+2:]...)
					i--
				}
			case "(bin)":
				if i >= 0 {
					for j := i - 1; j >= 0; j-- {
						if input[j][0] != ' ' {
							num, err := strconv.ParseInt(input[j], 2, 64)
							if err == nil {
								input[j] = strconv.Itoa(int(num))
							}
							break
						}
					}
				}
				if i < len(input)-1 {
					input = append(input[:i], input[i+2:]...)
					i--
				}
			case "(up)":
				if i >= 0 {
					for j := i - 1; j >= 0; j-- {
						if input[j][0] != ' ' {
							input[j] = strings.ToUpper(input[j])
							break
						}
					}
				}
				if i < len(input)-1 {
					input = append(input[:i], input[i+2:]...)
					i--
				}
			case "(cap)":
				if i >= 0 {
					for j := i - 1; j >= 0; j-- {
						if input[j][0] != ' ' {
							input[j] = Capitalize(input[j])
							break
						}
					}
				}
				if i < len(input)-1 {
					input = append(input[:i], input[i+2:]...)
					i--
				}
			case "(low)":
				if i >= 0 {
					for j := i - 1; j >= 0; j-- {
						if input[j][0] != ' ' {
							input[j] = strings.ToLower(input[j])
							break
						}
					}
				}
				if i < len(input)-1 {
					input = append(input[:i], input[i+2:]...)
					i--
				}
			}
		}
	}
	return input
}
