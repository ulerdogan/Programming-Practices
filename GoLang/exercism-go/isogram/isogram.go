package isogram

import "strings"

func IsIsogram(word string) bool {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWYZ"
	word = strings.ToUpper(word)

	for _, l := range alphabet {
		if strings.Count(word, string(l)) > 1 {
			return false
		}
	}

	return true
}
