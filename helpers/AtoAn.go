package helpers

// convert "a" to "an" if it come before vowel or "h"
func AtoAn(s string) string {
	a := []rune(s)
	for i := 0; i < len(a); i++ {
		if (a[i] == 'a' || a[i] == 'A') && (i == 0 || (i-1 >= 0 && a[i-1] == ' ')) {
			if i+2 < len(a) && a[i+1] == ' ' {
				for j := i + 1; j < len(a); j++ {
					if IsVowelOrH(a[j]) {
						t := a[i]
						a = append(a[:i], a[i+1:]...)
						a = append(a[:i], append([]rune{t, 'n'}, a[i:]...)...)
						break
					}
				}
			}
		}
	}
	return string(a)
}
