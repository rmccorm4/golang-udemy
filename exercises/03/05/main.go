package main

import "fmt"

func main() {
    for i := 10; i < 101; i++ {
        fmt.Println(i, i % 4)
    }
}
