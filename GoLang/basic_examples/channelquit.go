package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	arr := make([]int, 0)
	ch := make(chan int)
	q := make(chan bool)

	go randToChan(ch, q)
	getToChan(ch, q, &arr)
	
	fmt.Println(arr)

}

func giveRand(max int, seed int64) int {
	rand.Seed(seed)
	return rand.Intn(max)
}

func randToChan(ch chan<- int, q chan<- bool) {
	for {
		switch r := giveRand(100, time.Now().UnixNano()); r {
		case 33:
			q <- true
		default:
			ch <- r
		}
	}
}

func getToChan(ch <-chan int, q <-chan bool, arr *[]int) {
	for {
		select {
		case i := <-ch:
			*arr = append(*arr, i)
		case <-q:
			return
		}
	}
}