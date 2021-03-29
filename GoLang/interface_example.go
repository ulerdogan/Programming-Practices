package main

import "fmt"

type eurohesabi struct {
	iban int
	total int
}

type dolarhesabi struct {
	iban int
	total int
}

func (eur *eurohesabi) parayatir() int {
	eur.total++
	return eur.total
}

func (usd *dolarhesabi) parayatir() int {
	usd.total++
	return usd.total
}

type transaction interface {
	parayatir() int
}

func send(tr transaction) {
	fmt.Println("total:", tr.parayatir())
}

func main() {
	EH := eurohesabi{123456, 200}
	DH := dolarhesabi{98765, 300}

	send(&EH)
	send(&DH)

	fmt.Println(EH, DH)
}