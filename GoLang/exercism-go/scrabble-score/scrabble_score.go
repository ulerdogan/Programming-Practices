package scrabble

import "strings"

func Score(word string) (score int) {
	word = strings.ToUpper(word)
	letters := map[string]int{}

	for _, l := range word {
		letters[string(l)]++
	}

	for k, v := range letters {
		switch k {
		case "A", "E", "I", "O", "U", "L", "N", "R", "S", "T":
			score += 1*v
		case "D", "G":
			score += 2*v
		case "B", "C", "M", "P":
			score += 3*v
		case "F", "H", "V", "W", "Y":
			score += 4*v
		case "K":
			score += 5*v
		case "J", "X":
			score += 8*v
		case "Q", "Z":
			score += 10*v
		}
	}

	return
}