package helpers

import "strings"

func AtoAn(s string) string {
	a := strings.Split(s, " ")
	for i := 1; i < len(a); i++ {
		if (a[i-1] == "a" || a[i-1] == "A") || (strings.HasPrefix(a[i-1],"A")) && IsVowelOrH(rune(a[i][0])) {
			a[i-1] += "n"
		}
	}
	return strings.Join(a, " ")
}

func IsVowelOrH(r rune) bool {
	return r == 'a' || r == 'e' || r == 'o' || r == 'i' || r == 'u' || r == 'h' || r == 'A' || r == 'E' || r == 'O' || r == 'I' || r == 'U' || r == 'H'
}
