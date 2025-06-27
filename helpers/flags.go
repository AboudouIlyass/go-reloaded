package helpers

import (
	"strconv"
	"strings"
)

func Flags(s string) string {
	input := strings.Split(s, " ")
	for i := 0; i < len(input); i++ {
		if len(input[i]) > 1 && input[i][0] == '(' && input[i][len(input[i])-1] == ',' {
			flag := input[i][1 : len(input[i])-1]
			if (flag == "up" || flag == "low" || flag == "cap") && i+1 < len(input) {
				if len(input[i+1]) > 1 {
					if input[i+1][len(input[i+1])-1] == ')' {
						numStr := input[i+1][:len(input[i+1])-1]
						if n, err := strconv.Atoi(numStr); err == nil {
							for j := i - 1; j >= 0 && n > 0; j-- {
								if input[j] != "" {
									switch flag {
									case "up":
										input[j] = strings.ToUpper(input[j])
									case "low":
										input[j] = strings.ToLower(input[j])
									case "cap":
										input[j] = Capitalize(input[j])
									}
									n--
								}
							}
							input = DeleteFromSlice(input, i+1)
							input = DeleteFromSlice(input, i)
							i--
							continue
						}
					}
				}
			}
		}
		switch input[i] {
		case "(up)":
			for j := i - 1; j >= 0; j-- {
				if input[j] != "" {
					input[j] = strings.ToUpper(input[j])
					break
				}
			}
			input = DeleteFromSlice(input, i)
			i--
		case "(low)":
			for j := i - 1; j >= 0; j-- {
				if input[j] != "" {
					input[j] = strings.ToLower(input[j])
					break
				}
			}
			input = DeleteFromSlice(input, i)
			i--
		case "(cap)":
			for j := i - 1; j >= 0; j-- {
				if input[j] != "" {
					input[j] = Capitalize(input[j])
					break
				}
			}
			input = DeleteFromSlice(input, i)
			i--
		case "(bin)":
			for j := i - 1; j >= 0; j-- {
				if input[j] != "" {
					num, err := strconv.ParseInt(input[j], 2, 64)
					if err == nil {
						input[j] = strconv.Itoa(int(num))
					}
					break
				}
			}
			input = DeleteFromSlice(input, i)
			i--
		case "(hex)":
			for j := i - 1; j >= 0; j-- {
				if input[j] != "" {
					num, err := strconv.ParseInt(input[j], 16, 64)
					if err == nil {
						input[j] = strconv.Itoa(int(num))
					}
					break
				}
			}
			input = DeleteFromSlice(input, i)
			i--
		}
	}
	out := strings.Join(input, " ")
	return out
}
