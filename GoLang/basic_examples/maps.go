package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println()

	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("[eng] to [tr]")
	}
	query := args[0]

	//var dict map[string]string
	dict := map[string]string{
		"good":  "kötü",
		"great": "harika",
	}

	dict["good"] = "iyi"
	dict["up"] = "yukarı"

	//delete(dict, "good")

	if value, ok := dict[query]; ok {
		fmt.Printf("%q means %#v\n", query, value)
		return
	}

	fmt.Printf("%q not found\n", query)

	

}
