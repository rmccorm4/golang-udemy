package main

import (
    "fmt"
)

func main() {
    // x := type{values} // composite literal
    x := []int{1, 2,4,8,16,32}
    fmt.Println(x)

    for i, v := range x {
        fmt.Printf("2^%d = %d\n", i, v)
    }

    fmt.Println(x[:])
    fmt.Println(x[1:])
    fmt.Println(x[1:3])
    fmt.Println(x[:2])
    fmt.Println(x[:len(x)-1])
    fmt.Println(x[len(x)-1])
}
