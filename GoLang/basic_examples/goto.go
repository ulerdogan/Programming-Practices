package main

import (
	"fmt"
)

func main() {

	// cannot step over any variable declarations
	// goto loop XXXXXXXXXXXXXXXXXXXXX
	var i int

loop:
	if i < 3 {
		fmt.Println("looooop")
		i++
		goto loop
	}
}
