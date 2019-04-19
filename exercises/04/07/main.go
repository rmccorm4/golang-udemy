package main

import (
    "fmt"
)

func main() {
    spreadsheet := [][]string{}
    jb := []string{"James", "Bond", "Shaken, not stirred."}
    mp := []string{"Miss", "Moneypenny", "Helloooooo, James."}
    spreadsheet = append(spreadsheet, jb, mp)

    for i, row := range spreadsheet {
        for j, col := range row {
            fmt.Println(i, j, col)
        }
    }
    fmt.Println(spreadsheet)
}
