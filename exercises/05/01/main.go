package main

import (
	"fmt"
)

type person struct {
	first     string
	last      string
	icFlavors []string
}

func main() {
	p1 := person{
		first: "Ryan",
		last:  "McCormick",
		icFlavors: []string{
			"Mint Chocolate Chip",
			"Cookies and Cream"},
	}

	p2 := person{
		first:     "Jun",
		last:      "Seba",
		icFlavors: []string{},
	}

	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.icFlavors {
		fmt.Println("\t", i, v)
	}
	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, v := range p2.icFlavors {
		fmt.Println("\t", i, v)
	}
}
