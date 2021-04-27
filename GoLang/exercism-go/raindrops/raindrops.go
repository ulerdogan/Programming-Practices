package raindrops

import "strconv"

func Convert(num int) string {
	var output string = ""

	if num%3 == 0 {
		output = output + "Pling"
	}
	if num%5 == 0 {
		output = output + "Plang"
	}
	if num%7 == 0 {
		output = output + "Plong"
	}

	if output == "" {
		return strconv.Itoa(num)
	}
	return output
}
