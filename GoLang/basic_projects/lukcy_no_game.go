package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const maxTurns = 5

func main() {
	fmt.Println("🎉 Lucky No - Randomization Game")
	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("A NUM")
		return
	}

	guess, err := strconv.Atoi(args[0])

	if err != nil {
		fmt.Println("HATA")
		return
	}

	if guess<0 {
		fmt.Println("NEGATİF")
		return
	}

	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(guess + 1)

		if n == guess {
			fmt.Println("WIN")
			if turn == 0 {
				fmt.Println("İlk denemede")
			}
			return
		}
	}

	fmt.Println("LOST :(")
}
