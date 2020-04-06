package main

import (
	"fmt"
	"sync"
)

var rwMutex sync.RWMutex
var wg sync.WaitGroup
var x int = 0

func read(num int) {
	rwMutex.RLock()
	fmt.Println("I am read goroutine ", num, "x is ", x)
	rwMutex.RUnlock()
	wg.Done()
}

func write(num int) {
	rwMutex.Lock()
	x++
	fmt.Println("I am write goroutine ", num, "x is ", x)
	rwMutex.Unlock()
	wg.Done()
}

func main() {
	for i := 0; i < 7; i++ {
		wg.Add(1)
		go read(i)
	}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go write(i)
	}

	wg.Wait()
	fmt.Println("x is", x)

}
