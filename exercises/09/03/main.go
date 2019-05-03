package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ctr := 0
	const gs = 100
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			fmt.Println("GoRoutines:", runtime.NumGoroutine())
			myctr := ctr
			runtime.Gosched()
			myctr++
			ctr = myctr
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(ctr)
}
