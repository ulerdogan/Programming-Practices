package luhn

import (
	"strconv"
	"strings"
	"unicode"
)

func Valid(id string) (luhn bool) {
	id = strings.ReplaceAll(id, " ", "")
	var sum int

	if len(id) <= 1 {
		return
	}

	for i, _ := range id {
		l := id[len(id)-1-i]
		if unicode.IsLetter(int32(l)) {
			return
		}

		if i%2 != 0 {
			digit, _ := strconv.Atoi(string(l))
			if digit*2 > 9 {
				sum += digit*2 - 9
			} else {
				sum += digit * 2
			}

		} else {
			digit, _ := strconv.Atoi(string(l))
			sum += digit
		}
	}

	if sum%10 == 0 {
		luhn = true
	}

	return
}
