package main

import (
	"fmt"
)

func main() {
	x := 42
	fmt.Printf("%T\n", &x)
	fmt.Printf("%v\n", &x)
	fmt.Printf("%T\n", x)
	fmt.Printf("%v\n", x)
}
