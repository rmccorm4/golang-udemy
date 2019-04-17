package main

import "fmt"

func main() {
    year := 1997
    for year >= 1997 && year < 2020 {
        fmt.Println(year)
        year++
    }
}
