package main

import (
	"fmt"
	"github.com/inancgumus/screen"
	"time"
)

func main() {

	zero := [5]string{
		"███",
		"█ █",
		"█ █",
		"█ █",
		"███",
	}
	one := [5]string{
		"██ ",
		" █ ",
		" █ ",
		" █ ",
		"███",
	}
	two := [5]string{
		"███",
		"  █",
		"███",
		"█  ",
		"███",
	}
	three := [5]string{
		"███",
		"  █",
		"███",
		"  █",
		"███",
	}
	four := [5]string{
		"█ █",
		"█ █",
		"███",
		"  █",
		"  █",
	}
	five := [5]string{
		"███",
		"█  ",
		"███",
		"  █",
		"███",
	}
	six := [5]string{
		"███",
		"█  ",
		"███",
		"█ █",
		"███",
	}
	seven := [5]string{
		"███",
		"  █",
		"  █",
		"  █",
		"  █",
	}
	eight := [5]string{
		"███",
		"█ █",
		"███",
		"█ █",
		"███",
	}
	nine := [5]string{
		"███",
		"█ █",
		"███",
		"  █",
		"███",
	}
	divider := [5]string{
		"   ",
		" █ ",
		"   ",
		" █ ",
		"   ",
	}
	blank := [5]string{
		"   ",
		"   ",
		"   ",
		"   ",
		"   ",
	}

	numbers := [10][5]string{zero, one, two, three, four, five, six, seven, eight, nine}

	screen.Clear()
	screen.MoveTopLeft()
	for i := 0; i < 5; i++ {
		for j := 0; j < 10; j++ {
			fmt.Print(numbers[j][i], " ")
		}
		fmt.Println()
	}

	for {
		screen.Clear()
		screen.MoveTopLeft()
		hour := time.Now().Hour()
		minutes := time.Now().Minute()
		seconds := time.Now().Second()

		for j := 0; j < 5; j++ {
			clock := [8]string{numbers[hour/10][j], numbers[hour%10][j], divider[j], numbers[minutes/10][j], numbers[minutes%10][j], divider[j], numbers[seconds/10][j], numbers[seconds%10][j]}

			if seconds%2 == 0 {
				clock[2], clock[5] = blank[j], blank[j]
			} else {
				clock[2], clock[5] = divider[j], divider[j]
			}

			for i := 0; i < 8; i++ {
				fmt.Print(clock[i], " ")
			}
			fmt.Println()
		}

		time.Sleep(time.Second)
	}

}
