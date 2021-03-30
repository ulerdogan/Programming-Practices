package main

import "fmt"

func main() {
	channel := make(chan int)

	go func(top int) {
		for i := 0; i < top; i++ {
			channel <- i
		}
		close(channel)
	}(20)

	for j := range channel {
		fmt.Println(j)
	}

	fmt.Println("end")

}
