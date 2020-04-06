package main

import "sync"

var wg sync.WaitGroup
var mtex sync.Mutex

func main() {
	for {
		wg.Add(1)
		go makePan()
	}

	for {
		wg.Add(1)
		go eatPan()
	}

	wg.Wait()
}

func makePan() {
	wg.Done()
}

func eatPan() {
	wg.Done()
}
