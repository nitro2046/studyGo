package main

import "fmt"

func main() {
	a, b := 10, 30
	i, j := addSub(a, b)
	fmt.Println(i, j)
	fmt.Println(numFunc(a, b, numAdd))
}

func addSub(c, d int) (int, int) {
	sum := c + d
	sub := c - d
	return  sum, sub
}

func numFunc(a, b int, math func(c,d int) int) int{
	return math(a, b)
}

func numAdd(a, b int) int  {
	return a + b
}