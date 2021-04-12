package main

import (
	"fmt"
)

func main() {
	fmt.Println()

	var x [2][2]int

	x[0][0] = 1
	x[0][1] = 2
	fmt.Println(x[0])
	for range x[0] {
		fmt.Println("aaa")
	}
}
