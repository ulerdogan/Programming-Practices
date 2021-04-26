package main

import (
	"fmt"
)

func main() {
	myChan := make(chan string)

	go merhaba(myChan)

	fmt.Println(<-myChan)
}

func merhaba(chan1 chan string) {
	chan1 <- "merhaba"
}
