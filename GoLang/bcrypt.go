package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)
// go get golang.org/x/crypto/bcrypt >> in bash

func main() {
	pwd := "password123"

	bs, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Println(bs)

	err = bcrypt.CompareHashAndPassword(bs, []byte(pwd))
	if err != nil {
		fmt.Println("YOU CAN'T LOGIN")
		return
	}

	fmt.Println("You're logged in")


}
