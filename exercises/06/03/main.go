package main

import (
    "fmt"
)

func foo() {
    fmt.Println("foo")
}

func main() {
    defer foo()
    fmt.Println("Hi")
}
