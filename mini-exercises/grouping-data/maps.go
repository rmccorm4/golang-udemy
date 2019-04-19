package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"Ryan":    22,
		"Francis": 26,
	}

	fmt.Println(m["Ryan"])

	v, ok := m["Francis"]
	if ok {
		fmt.Printf("Francis exists: %d\n", v)
	}

	// Bring it all together
	if v, ok := m["Bobby"]; ok {
		fmt.Printf("Bobby exists: %d\n", v)
	} else {
		fmt.Println("Bobby does not exist.")
	}
}
