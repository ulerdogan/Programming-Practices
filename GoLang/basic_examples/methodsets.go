package main

import (
	"fmt"
)

type person struct {
	age int
}

func (_person *person) speak() {
	fmt.Println("Benim yasim", _person.age)
}

type human interface {
	speak()
}

func saySomething(_human human) {
	_human.speak()
}

func main() {
	
	kisi := person{16}

	saySomething(&kisi)
}