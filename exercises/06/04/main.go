package main

import (
    "fmt"
)

type person struct {
    first string
    last string
    age int
}

func (p person) speak() {
    fmt.Println("Name:", p.first, p.last)
    fmt.Println("Age:", p.age)
}

func main() {
    p1 := person{
        first: "Ryan",
        last:  "McCormick",
        age:    22,
    }
    p1.speak()
}
