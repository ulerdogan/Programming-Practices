package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

//board lengths
const (
	yLen int = 10
	xLen int = 50
)

func main() {
	//velocities
	var (
		vX int = 1
		vY int = 1
	)

	var cell rune

	var animationScreen [yLen][xLen]bool

	animationScreen[0][0] = true

	for {
		screen.Clear()
		screen.MoveTopLeft()
		wall(2*xLen+2, true)
		for i, y := range animationScreen {
			wall(1, false)
			for j, x := range y {

				cell = ' '

				if x == true {
					cell = '.' //🏀
					if (i == 0 && vY < 0) || (i == yLen - 1 && vY > 0) {
						vY = -vY
					}
					if (j == 0 && vX < 0) || (j == xLen - 1 && vX > 0) {
						vX = -vX
					}
					animationScreen[i][j] = false
					animationScreen[i+vY][j+vX] = true
				}
				fmt.Print(string(cell), " ")
			}
			wall(1, false)
			fmt.Println()
		}
		wall(2*xLen+2, true)

		time.Sleep(time.Second / 10)

	}

}

func wall(x int, y bool) {
	var i int
	for i = 0; i < x; i++ {
		fmt.Print("█")
	}
	if y {
		fmt.Println()
	}
}
