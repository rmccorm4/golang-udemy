package main

import "fmt"

type ID int
var x ID

func main() {
	fmt.Println(x)
	// Print the type
	fmt.Printf("%T\n", x)
	x := 42
	fmt.Println(x)
}