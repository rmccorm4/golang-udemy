package main

import "fmt"

func main() {
    answers := `
        true && true   = true
        true && false  = false
        true || true   = true
        true || false  = true
        !true          = false`

    fmt.Println(answers)
}
