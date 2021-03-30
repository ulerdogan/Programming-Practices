package main

import (
	"fmt"
	"time"
)

func main() {
	channelTen := make(chan bool)
	channelThree := make(chan bool)
	go tenSeconds(channelTen)
	go threeSeconds(channelThree)

	for i := 1; i <= 2; i++ {
		select {
		case ten := <-channelTen:
			fmt.Println(ten)
		case three := <-channelThree:
			fmt.Println(three)
		}
	}

}

func tenSeconds(ch chan bool) {
	time.Sleep(time.Second * 10)
	ch <- true
}

func threeSeconds(ch chan bool) {
	time.Sleep(time.Second * 3)
	ch <- true
}
