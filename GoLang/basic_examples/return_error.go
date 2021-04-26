package main

import (
	"fmt"
	"errors"
)

func main() {
	result, err := evenNum(12)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}

func evenNum(num int) (bool, error) {
	if num%2 != 0 {
		return false, errors.New("HATA")
	}

	return true, nil
}
