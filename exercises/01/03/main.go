package main

import "fmt"

var x = 42
var y = "James Bond"
var z = true

func main() {
	s := fmt.Sprintf("%d\n%s\n%t\n", x, y, z)
	fmt.Print(s)
}