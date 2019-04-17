package main

import "fmt"

func main() {
    switch {
    case 42 % 2 == 0:
        fmt.Println("The meaning of life is even.")
    default:
        fmt.Println("Default")
    }
}
