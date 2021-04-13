package main

import (
	"fmt"
	//"unicode/utf8"
)

func main() {
	fmt.Println()

	// words := []string{
	// 	"gopher",
	// 	"programmer",
	// 	"go language",
	// 	"go standard library",
	// }

	// var bwords [][]byte

	// for _, i := range words {
	// 	x := []byte(i)
	// 	fmt.Println(x)
	// 	bwords = append(bwords, x)
	// }

	// for _, j := range bwords {
	// 	fmt.Println(string(j))
	// }
	// fmt.Println("________")

	const word = "console"

	for _, i := range word {
		fmt.Printf("%c\ndecimal: %[1]d\nhex: 0x%[1]x\nbinary: 0b%08[1]b\n", i)
	}
	// print the word manually using runes
	fmt.Printf("with runes       : %s\n",
		string([]byte{'c', 'o', 'n', 's', 'o', 'l', 'e'}))

	// print the word manually using decimals
	fmt.Printf("with decimals    : %s\n",
		string([]byte{99, 111, 110, 115, 111, 108, 101}))

	// print the word manually using hexadecimals
	fmt.Printf("with hexadecimals: %s\n",
		string([]byte{0x63, 0x6f, 0x6e, 0x73, 0x6f, 0x6c, 0x65}))
}
