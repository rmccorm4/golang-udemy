package main

import (
    "fmt"
)

func foo(xi ...int) int {
    sum := 0
    for _, v := range xi {
        sum += v
    }
    return sum
}

func bar(xi []int) int {
    sum := 0
    for _, v := range xi {
        sum += v
    }
    return sum
}

func main() {
    xi := []int{1,2,3,4,5}
    fmt.Println(foo(xi...), bar(xi))
}
