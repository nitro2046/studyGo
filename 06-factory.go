package main

import "fmt"

type animal interface {
	eating()
	sleeping()
}

type dog struct {
	color string
}

func (d dog) eating() {
	fmt.Printf("%s's dog is eating\n", d.color)
}

func (d dog) sleeping() {
	fmt.Printf("%s's dog is sleeping\n", d.color)
}

type cat struct {
	color string
}

func (c cat) eating() {
	fmt.Printf("%s's cat is eating\n", c.color)
}

func (c cat) sleeping() {
	fmt.Printf("%s's cat is sleeping\n", c.color)
}

func factory(color, animal_type string) animal {
	switch animal_type {
	case "dog":
		return &dog{color}
	case "cat":
		return &cat{color}
	default:
		return nil
	}
}

func main() {
	var d1 = dog{"blue"}
	d1.eating()
	c1 := cat{"red"}
	c1.sleeping()
	fmt.Println("----------------------------")
	d2 := factory("black", "dog")
	d2.sleeping()
	c2 := factory("white", "cat")
	c2.eating()
	a := factory("green", "pig")
	fmt.Println(a)
}
