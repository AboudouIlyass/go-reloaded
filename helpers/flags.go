package helpers

import (
	"strconv"
	"strings"
	"unicode"
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
						for _, k := range input[j] {
							if k != ' ' {
								input[j] = ToLower(&input[j])
								num--
								break
							}
						}
					}
				}
				input = append(input[:i], input[i+1:]...)
				i--

			case "(cap, n)":
				c := fl.data
				numStr := strings.TrimLeft(c, "0")
				if numStr > "9223372036854775807" {
					numStr = "9223372036854775807"
				}
				num, _ := strconv.Atoi(numStr)
				if i-1 >= 0 {
					for j := i - 1; j >= 0 && num > 0; j = j - 1 {
						for _, k := range input[j] {
							if k != ' ' {
								input[j] = Capitalize(&input[j])
								num--
								break
							}
						}
					}
				}
				input = append(input[:i], input[i+1:]...)
				i--
			case "(up, n)":
				c := fl.data
				numStr := strings.TrimLeft(c, "0")
				if numStr > "9223372036854775807" {
					numStr = "9223372036854775807"
				}
				num, _ := strconv.Atoi(numStr)
				if i-1 >= 0 {
					for j := i - 1; j >= 0 && num > 0; j = j - 1 {
						for _, k := range input[j] {
							if k != ' ' {
								input[j] = ToUpper(&input[j])
								num--
								break
							}
						}
					}
				}
				input = append(input[:i], input[i+1:]...)
				i--

			case "(hex)":
				if i > 0 {
					j := i - 1
					for ; j >= 0; j-- {
						if strings.TrimSpace(input[j]) != "" {
							num, err := strconv.ParseInt(input[j], 16, 64)
							if err == nil {
								input[j] = strconv.Itoa(int(num))
							}
							break
						}
					}
					input = append(input[:i], input[i+1:]...)
					i--
				}
			case "(bin)":
				if i > 0 {
					j := i - 1
					for ; j >= 0; j-- {
						if strings.TrimSpace(input[j]) != "" {
							num, err := strconv.ParseInt(input[j], 2, 64)
							if err == nil {
								input[j] = strconv.Itoa(int(num))
							}
							break
						}
					}
					input = append(input[:i], input[i+1:]...)
					i--
				}
			case "(up)":
				if i > 0 {
					for j := i-1 ; j>= 0; j--{
						if strings.TrimSpace(input[j]) != "" {
							ToUpper(&input[j])
							break
						}
					}
				}
				input = append(input[:i], input[i+1:]...)
				i--
			case "(cap)":
				if i > 0 {
					for j := i-1 ; j>= 0; j--{
						if strings.TrimSpace(input[j]) != "" {
							Capitalize(&input[j])
							break
						}
					}
				}
				input = append(input[:i], input[i+1:]...)
				i--
			case "(low)":
				if i > 0 {
					for j := i-1 ; j>= 0; j--{
						if strings.TrimSpace(input[j]) != "" {
							ToLower(&input[j])
							break
						}
					}
				}
				input = append(input[:i], input[i+1:]...)
				i--
			}
		}
	}
	return input
}

func ToLower(s *string) string {
	r := []rune(*s)
	for i, v := range r {
		if unicode.IsLetter(v) {
			r[i] = unicode.ToLower(v)
		}
	}
	*s = string(r)
	return *s
}

func ToUpper(s *string) string {
	r := []rune(*s)
	for i, v := range r {
		if unicode.IsLetter(v) {
			r[i] = unicode.ToUpper(v)
		}
	}
	*s = string(r)
	return *s
}
func Capitalize(s *string) string {
	r := []rune(*s)
	for i, v := range r {
		if unicode.IsLetter(v) {
			r[i] = unicode.ToUpper(r[i])
			break
		}
	}
	*s = string(r)
	return *s
}
