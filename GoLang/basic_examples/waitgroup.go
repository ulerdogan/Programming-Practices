package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("begin cpu", runtime.NumCPU())
	fmt.Println("begin gs", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("AAAAAA")
		wg.Done()
	}()

	go func() {
		fmt.Println("cpu", runtime.NumCPU())
		fmt.Println("gs", runtime.NumGoroutine())
		fmt.Println("BBBBBB")
		wg.Done()
	}()
	wg.Wait()

	fmt.Println("end cpu", runtime.NumCPU())
	fmt.Println("end gs", runtime.NumGoroutine())
}
