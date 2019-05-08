package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	const gs = 10
	for i := 0; i < gs; i++ {
		go func(k int) {
			for j := 0; j < 10; j++ {
				c <- k*10 + j
			}
		}(i)
	}

	for i := 0; i < 100; i++ {
		fmt.Println(i, <-c)
	}
}
