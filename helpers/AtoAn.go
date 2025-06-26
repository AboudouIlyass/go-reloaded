package helpers

func AtoAn(s string) string {
	a := []rune(s)
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[j] != ' ' {
				if (a[i] == 'a' || a[i] == 'A') &&
					(i == 0 || (i-1 >= 0 && a[i-1] == ' ')) &&
					IsVowelOrH(a[j]) {

					t := a[i]
					a = append(a[:i], a[i+1:]...)
					a = append(a[:i], append([]rune{t, 'n'}, a[i:]...)...)
				}
				break // stop after the first non-space char
			}
		}
	}
	return string(a)
}
func IsVowelOrH(r rune) bool {
	return r == 'a' || r == 'e' || r == 'o' || r == 'i' || r == 'u' || r == 'h' || r == 'A' || r == 'E' || r == 'O' || r == 'I' || r == 'U' || r == 'H'
}
