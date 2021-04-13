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
	spam := []byte("*")

	text := args[0]
	size := len(text)
	buf := make([]byte, 0, size)

	for i := 0; i < size; i++ {
		buf = append(buf, text[i])
		if len(text[i:]) >= nlink && text[i:i+nlink] == link {
			buf = append(buf, spam)
		}
	}

}
