package main

import "fmt"

func main() {
    for i := 0; i < 11; i++ {
        if i % 2 == 0 {
            fmt.Printf("%d is even\n", i)
        }
    }
}
