package main

import (
	"fmt"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan bool)

	go send(even, odd, quit)
	receive(even, odd, quit)

	fmt.Println("finito")

}

func send(even, odd chan<- int, quit chan<- bool) {
	for i := 0; i < 100; i++ {
		if i%2 == 0{
			even <- i
		} else {
			odd <- i
		}
	}
	quit <- true
}

func receive(even, odd <-chan int, quit <-chan bool) {
	for {
		select {
		case v := <-even:
			fmt.Println("e", v)
		case v := <-odd:
			fmt.Println("o", v)
		case <-quit:
			return
		}
	}
}