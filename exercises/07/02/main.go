package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

func changeMe(pp *person) {
	(*pp).first = "Bobby"
}

func main() {
	p := person{
		first: "Ryan",
		last:  "McCormick",
	}
	fmt.Println(p)
	changeMe(&p)
	fmt.Println(p)
}
