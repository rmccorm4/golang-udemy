package main

import (
	"fmt"
)

// x only exists within the scope of this function
func iterX() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	a := iterX()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	b := iterX()
	fmt.Println(b())
	fmt.Println(b())
}
