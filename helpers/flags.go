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
							if unicode.IsLetter(k) {
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
			case "(up, n)":
			case "(hex)":
			case "(bin)":
			case "(up)":
			case "(cap)":
			case "(low)":
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
