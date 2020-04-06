package main

import (
	"fmt"
	"sync"
)

var one sync.Once
var x int = 0
var wg sync.WaitGroup

func main() {
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go one.Do(func() {
			x++
		})
		wg.Done()
	}
	wg.Wait()
	fmt.Println(x)
}
