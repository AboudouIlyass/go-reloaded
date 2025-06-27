package helpers

import "unicode"

func Capitalize(s string) string {
	r := []rune(s)
	for i, v := range s {
		if unicode.IsLetter(v) {
			r[i] = unicode.ToUpper(r[i])
			break
		}
	}
	return string(r)
}

func AreDigits(s string) bool {
	if len(s) == 0 {
		return false
	}
	if s[0] == '+' || s[0] == '-' {
		s = s[1:]
		if len(s) == 0 {
			return false
		}
	}
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}

func IsPunc(s rune) bool {
	return (s == '.' || s == ',' || s == '!' || s == '?' || s == ':' || s == ';')
}

func IsVowelOrH(r rune) bool {
	return r == 'a' || r == 'e' || r == 'o' || r == 'i' || r == 'u' || r == 'h' || r == 'A' || r == 'E' || r == 'O' || r == 'I' || r == 'U' || r == 'H'
}

func MyDelete(s []rune, pos int) []rune {
	if pos >= len(s) || pos < 0 {
		return s
	}
	return append(s[:pos], s[pos+1:]...)
}

func DeleteFromSlice(s []string, pos int) []string {
	if pos >= len(s) || pos < 0 {
		return s
	}
	return append(s[:pos], s[pos+1:]...)
}

func AddToSlice(s []string, pos int, str string) []string {
	if pos >= len(s) || pos < 0 {
		return s
	}
	return append(s[:pos], append([]string{str}, s[pos:]...)...)
}

func MyAdd(s []rune, pos int, char rune) []rune {
	if pos >= len(s) || pos < 0 {
		return s
	}
	return append(s[:pos], append([]rune{char}, s[pos:]...)...)
}
