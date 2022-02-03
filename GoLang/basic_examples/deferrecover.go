package main

import (
	"fmt"
)

func main() {
	ptr := new(int)

	*ptr = 2

	defer func() {
		h := recover()
		fmt.Println("rc", h)
	}()
	
	if *ptr !=3 {
		panic("pnc")
	}
	fmt.Println("ulas")	
}