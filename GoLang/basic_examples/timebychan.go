package main

import (
	"fmt"
	"time"
)

func readTime(ch chan<- string) {
	for {
		currentTime := time.Now().String()
		ch <- currentTime
		time.Sleep(1 * time.Second)
	}
}

func printTime(ch <-chan string) {
	for {
		currentTime := <-ch
		fmt.Println(currentTime)
	}
}

func main() {
	ch := make(chan string)
	go readTime(ch)
	go printTime(ch)

	var key string
	fmt.Scanln(&key)
	fmt.Println(key)
}
