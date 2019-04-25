package main

import (
	"fmt"
	"sort"
)

func main() {
	unsorted := []int{4, 1, 64, 16, 9, 25, 49, 36}
	sorted := make([]int, len(unsorted))
	copy(sorted, unsorted)
	sort.Ints(sorted)
	fmt.Println(unsorted, sorted)

	unsortedStr := []string{"Hello", "World", "Abc"}
	sortedStr := make([]string, len(unsortedStr))
	copy(sortedStr, unsortedStr)
	sort.Strings(sortedStr)
	fmt.Println(unsortedStr, sortedStr)
}
