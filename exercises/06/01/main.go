package main

import (
    "fmt"
)

func foo() int {
    return 1
}

func bar() int {
    return 2
}

func main() {
    fmt.Println(foo(), bar())
}
