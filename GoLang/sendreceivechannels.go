package main

import (
	"fmt"
)

func main() {

	channel := make(chan int)
	go send(channel, 33)
	receive(channel)
}

// only can send to channel
func send(ch chan<- int, number int) {
	ch <- number
}

// only can receive from the channel
func receive(ch <-chan int) {
	fmt.Printf("The lucky number is %d\n", <-ch)
}
