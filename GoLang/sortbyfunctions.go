package main

import (
	"fmt"
	"sort"
)

type byChar []string

func (a byChar) Len() int           { return len(a) }
func (a byChar) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byChar) Less(i, j int) bool { return string(a[i][1]) < string(a[j][1]) }

func main() {

	fmt.Println("Hola el mundo!")

	fruits := []string{"kiwi", "apple", "banana", "orange"}
	sort.Sort(byChar(fruits))
	fmt.Println(fruits)


}
