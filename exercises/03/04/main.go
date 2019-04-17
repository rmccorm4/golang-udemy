package main

import "fmt"

func main() {
    year := 1997
    for {
        if year > 2019 {
            break
        } else if year >= 1997 {
            fmt.Println(year)
        }
        year++
    }
}
