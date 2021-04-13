package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println()

	text := "Ulaş Erdoğan ğğğüüüüğüğüğüğxiI"

	// for i := 0; i < len(text); {
	// 	r, size := utf8.DecodeRuneInString(text[i:])
	// 	fmt.Printf("%c", r)

	// 	i += size
	// }
	
	//same with upr
	for _, r := range text {
		fmt.Printf("%c", r)
	}

	fmt.Println()

	x := []byte("ulaş erdoğan")

	for 0 < len(x) {
		//		gets first utf8 value and returns rune and width
		a, b := utf8.DecodeRune(x)

		fmt.Printf("%c %v \n", a, b)
		x = x[b:]

	}


}
