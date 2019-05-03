package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func foo() {
	fmt.Println("foo()")
	wg.Done()
}

func main() {
	const gs = 3
	wg.Add(gs)

	go foo()
	go func(){
		fmt.Println("bar in-line")
		wg.Done()
	}()
	go func(){
		fmt.Println("fubar in-line")
		wg.Done()
	}()

	fmt.Println("main")
	wg.Wait()
}