package main

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
)

func main() {
	var x int64 = int64(math.Pow(2, 10))

	var wg sync.WaitGroup
	wg.Add(2)

	lock := sync.Mutex{}

	go func() {
		for i := 0; i < 1000; i++ {
			atomic.AddInt64(&x, +2)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			lock.Lock()
			//atomic.AddInt64(&x, -2)
			x -= 2;
			lock.Unlock()
		}
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(x);
}
