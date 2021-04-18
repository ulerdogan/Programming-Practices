package main

import (
	"fmt"
	"os"
)

//"Here's my spammy page: http://youth-elixir.com"

func main() {

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("gimme somethin' to mask!")
		return
	}

	fmt.Println()

	link := "http://"
	nlink := len(link)
	const mask = '*'

	text := args[0]
	size := len(text)
	buf := make([]byte, 0, size)

	var in bool

	for i := 0; i < size; i++ {
		
		if len(text[i:]) >= nlink && text[i:i+nlink] == link {
			in = true

			buf = append(buf, link...)
			i += nlink
		}

		c := text[i]
		switch c {
		case ' ', '\t', '\n':
			in = false
		}

		if in {
			c = mask
		}

		buf = append(buf, c)
	}

	fmt.Println(string(buf))

}
