package main

import "fmt"

func main() {
    n := 42
    switch n % 2 {
    case 0:
        fmt.Println("Even")
    case 1:
        fmt.Println("Odd")
    default:
        fmt.Println("Default case, should never happen here since integers must be even or odd")
    }
}
