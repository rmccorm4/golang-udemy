package main

import "fmt"

func main() {
	x := `Name
	a
	b
	c`
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}