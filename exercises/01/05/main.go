package main

import "fmt"

type ID int

var x ID
var y int

func main() {
	fmt.Println(x)
	// Print the type
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

	// Conversion (not casting)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}