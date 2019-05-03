package main

import "fmt"

type human interface {
	speak()
}

type person struct {
	name string
	age  int
}

func (p *person) speak() {
	fmt.Println("Speak:", p.name, p.age)
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p := person{
		name: "Ryan",
		age:  22,
	}
	pp := &person{
		name: "Bob",
		age:  37,
	}

	// saySomething(p) // FAILS
	fmt.Println(p)
	saySomething(pp)
}
