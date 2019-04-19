package main

import (
    "fmt"
)

func main() {
    x := []int{1,2,4,8,16,32,64,128,256,512,1024}
    fmt.Println(x[:5])
    fmt.Println(x[5:])
    fmt.Println(x[2:7])
}

