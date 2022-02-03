package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

var x int = 3

const (
	a = iota
	b
	c
)

func main() {
	defer fmt.Println(time.Now())
	xP := &x
	inc(xP)
	fmt.Println(x)

	for {
		fmt.Println(a, b, c)
		break
	}

	fmt.Println("3:", isPrime(3), "| 4:", isPrime(4), "| 5:", isPrime(5))

	rand.Seed(time.Now().UnixNano())
	getRand := rand.Intn(100)

	switch {
	case isPrime(getRand):
		fmt.Println(getRand, "prime")
		//fallthrough
	case !isPrime(getRand):
		fmt.Println(getRand, "not-prime")
	}

	var names [4]string
	names = [4]string{"John", "Paul", "George", "Ringo"}
	surnames := make([]string, 4)
	surnames = []string{"Lennon", "McCartney", "Harrison", "Starr"}
	ns := make(map[string]string)
	for i, name := range names {
		ns[name] = surnames[i]
	}
	fmt.Println(ns)

	js, _ := json.MarshalIndent(ns, "", "  ")
	fmt.Println(string(js))
	
	fmt.Println(time.Now())

}

func inc(_x *int) {
	*_x++
}

func isPrime(_x int) bool {
	if _x < 2 {
		return false
	}

	for i := 2; i < _x; i++ {
		if _x%i == 0 {
			return false
		}
	}
	return true
}
