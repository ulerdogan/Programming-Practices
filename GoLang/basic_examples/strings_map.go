package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	str := "ulaşerdoğan"

	rot13 := func(k rune) rune {
		return k + 13
	}

	strR := strings.Map(rot13, str)

	fmt.Println(strR)

	tr := unicode.TurkishCase
	strU := strings.ToUpperSpecial(tr, str)
	fmt.Println(strU)

}
