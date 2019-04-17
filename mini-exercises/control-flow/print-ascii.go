package main

import "fmt"

func main() {
    fmt.Println("==== ASCII Table ====")
    for i := 0; i < 128; i++ {
        fmt.Printf("%d: %#U\n", i, i)
    }
}
