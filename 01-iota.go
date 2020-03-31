package main

import "fmt"

func main() {
	const (
		apple, banana = iota + 1, iota + 2
		a, b
		c, d
	)
	fmt.Println(apple, banana, a, b, c, d)
}
