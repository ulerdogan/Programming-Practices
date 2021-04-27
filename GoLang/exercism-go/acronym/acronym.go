package acronym

import (
	"unicode"
	"strings"
)

func Abbreviate(s string) string {

	abb := make([]rune, 0)
	s = strings.ToUpper(s)

	for i, _ := range []rune(s) {
		if i == 0 && unicode.IsLetter(rune(s[0])) {
			abb = append(abb, rune(s[0]))
		}

		if rune(s[i]) == ' ' || rune(s[i]) == '-' || rune(s[i]) == '_' {
			if unicode.IsLetter(rune(s[i+1])) {
				abb = append(abb, rune(s[i+1]))
			}
		}
	}

	return string(abb)
}
