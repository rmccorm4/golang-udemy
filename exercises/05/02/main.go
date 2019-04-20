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

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	mp1 := m["McCormick"]
	mp2 := m["Seba"]

	fmt.Println(mp1)
	fmt.Println(mp2)

	fmt.Println(mp1.first)
	fmt.Println(mp1.last)
	for i, v := range mp1.icFlavors {
		fmt.Println("\t", i, v)
	}
	fmt.Println(mp2.first)
	fmt.Println(mp2.last)
	for i, v := range mp2.icFlavors {
		fmt.Println("\t", i, v)
	}
}
