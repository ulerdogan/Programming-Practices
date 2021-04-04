package main

import (
	"fmt"
	"os"
	"strings"
)

const corpus = "" +
	"lazy cat jumps again and again and again"

func main() {

	words := strings.Fields(corpus)

	query := os.Args[1:]

queries: // create a label and  ----- has a different scope
	for _, q := range query {
	search:
		for i, w := range words {

			// passing unwanteds
			switch q {
			case "and", "or", "the":
				break search
			}

			if strings.ToUpper(q) == strings.ToUpper(w) {
				fmt.Printf("#%-2d: %q\n", i+1, w)
				//break
				break queries // break the main loop in inner -- finding only ONE word
				//continue queries // labeled cointinue alternative
			}
		}
	}

}
