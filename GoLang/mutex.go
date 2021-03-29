package main

import (
	"fmt"
	"sync"
)

func main() {
	counter := 0
	var wg sync.WaitGroup
	wg.Add(100)

	var m sync.Mutex

	for i := 0; i < 100; i++ {
		go func() {
			m.Lock()
			vr := counter		
			vr++
			counter = vr
			fmt.Println(counter)
			m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
}
