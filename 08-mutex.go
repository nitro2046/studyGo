package main

import (
	"fmt"
	"sync"
)

var x int = 0
var wg sync.WaitGroup
var mutex sync.Mutex

func main() {
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			mutex.Lock()
			x += 1
			mutex.Unlock()
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println(x)
}
