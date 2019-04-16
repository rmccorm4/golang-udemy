package main

import "fmt"

func main() {
	var x int = 42
	fmt.Printf("%d %b %#x\n", x, x, x)
	x <<= 1
	fmt.Printf("%d %b %#x\n", x, x, x)
}