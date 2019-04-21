package main

import (
	"fmt"
)

func factRecursive(n int) int {
	switch {
	case n < 0:
		return -1 // Invalid
	case n == 0:
		return 1
	default:
		return n * factRecursive(n-1)
	}
}

func main() {
	fmt.Println(factRecursive(-7))
	fmt.Println(factRecursive(0))
	fmt.Println(factRecursive(1))
	fmt.Println(factRecursive(5))
}
