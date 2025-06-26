package helpers

func AtoAn(s string) string {
    a := []rune(s)
    for i := 1; i < len(a); i++ {
        if (a[i] == 'a' || a[i] == 'A') &&
           (a[i-1] == ' ' || a[i-1] == '\'') &&
           i+2 < len(a) &&
           a[i+1] == ' ' &&
           IsVowelOrH(a[i+2]) {

            t := a[i]
            a = append(a[:i], a[i+1:]...)
            a = append(a[:i], append([]rune{t, 'n'}, a[i:]...)...)
        }
    }
    return string(a)
}

func IsVowelOrH(r rune) bool {
	return r == 'a' || r == 'e' || r == 'o' || r == 'i' || r == 'u' || r == 'h' || r == 'A' || r == 'E' || r == 'O' || r == 'I' || r == 'U' || r == 'H'
}
