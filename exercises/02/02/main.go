package main

import "fmt"

func main() {
	a := 7 == 42
	b := 7 != 42
	c := 7 <= 42
	d := 7 >= 42
	e := 7 < 42
	f := 7 > 42
	fmt.Println(a, b, c, d, e, f)
}