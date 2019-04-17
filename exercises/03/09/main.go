package main

import "fmt"

func main() {
	favSport := "Soccer"
	switch favSport {
        case "Baseball":
            fmt.Println("Gross")
        case "Soccer":
            fmt.Println("Same")
        case "Basketball":
            fmt.Println("Nice")
        default:
            fmt.Println("Huh, never heard of it")
	}
}
