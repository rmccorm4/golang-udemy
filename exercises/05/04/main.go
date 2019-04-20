package main

import (
	"fmt"
)

func main() {
	// Anonymous struct
	p1 := struct {
		first string
		last  string
	}{
		first: "Ryan",
		last:  "McCormick",
	}

	fmt.Println(p1)
}
