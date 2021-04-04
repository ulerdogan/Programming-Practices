package main

import (
	"fmt"
	"os"
	"time"
	"math/rand"
)

func main() {

	if args := os.Args[1:]; len(args) != 1 {
		fmt.Println("Invalid input")
		return
	}
	name := os.Args[1]

	rand.Seed(time.Now().UnixNano())

	moods := [...]string{"good", "happy", "awesome", "sad", "bad", "terrible"}

	randInt := rand.Intn(len(moods))

	fmt.Printf("%s feels %s\n", name, moods[randInt])

}
