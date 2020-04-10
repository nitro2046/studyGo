package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mtex sync.Mutex
var con *sync.Cond
var shop [5]int

var conMaker *sync.Cond
var mtexMaker sync.Mutex
var mianbaoCount int = 0

var eatNum int = 0

func main() {
	con = sync.NewCond(&mtex)
	conMaker = sync.NewCond(&mtexMaker)
	wg.Add(2)
	go makePan()
	go eatPan(1)
	go eatPan(2)
	go eatPan(3)

	wg.Wait()
}

func makePan() {
	mianbao := 1001
	num := 0
	for {
		mtex.Lock()
		time.Sleep(time.Millisecond * 100)
		shop[num] = mianbao
		fmt.Printf("生产了%d号饼\n", mianbao)
		mianbao++
		mianbaoCount++
		num = (num + 1) % 5
		con.Signal()
		mtex.Unlock()

		if mianbaoCount == 5 {
			mtexMaker.Lock()
			conMaker.Wait()
			mtexMaker.Unlock()
		}
	}

	wg.Done()
}

func eatPan(p int) {
	for {
		mtex.Lock()
		con.Wait()
		if mianbaoCount > 0 {
			time.Sleep(time.Millisecond * 10)
			fmt.Printf("%d号吃货，吃了%d号饼\n", p, shop[eatNum])
			eatNum = (eatNum + 1) % 5
			mianbaoCount--
		}
		mtex.Unlock()
		conMaker.Signal()

	}
	wg.Done()
}
