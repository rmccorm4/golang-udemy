package main

import (
	"fmt"
)

func evenApply(f func(...int) int, xi ...int) int {
	var ei []int
	for _, v := range xi {
		if v%2 == 0 {
			ei = append(ei, v)
		}
	}

	return f(ei...)
}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func product(xi ...int) int {
	prod := 1
	for _, v := range xi {
		prod *= v
	}
	return prod
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("All Sum:", sum(nums...))
	fmt.Println("Even Sum:", evenApply(sum, nums...))
	fmt.Println("All Product:", product(nums...))
	fmt.Println("Even Product:", evenApply(product, nums...))
}
