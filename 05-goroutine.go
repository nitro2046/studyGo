package main

import (
	"fmt"
	"time"
)

func main() {
	go tiyan(time.Millisecond * 100)
	fmt.Println("\n", fib(1000))
}

func fib(num int) int {
	x, y := 0, 1
	for i := 1; i < num; i++ {
		x, y = y, x+y
	}
	return y
}

func tiyan(shijian time.Duration) {
	for {
		for _, s := range `-/|\` {
			fmt.Printf("\r%c", s)
			time.Sleep(shijian)
		}
	}

}
