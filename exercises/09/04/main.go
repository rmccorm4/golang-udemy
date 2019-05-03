package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mtx sync.Mutex

func main() {
	ctr := 0
	const gs = 100
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			mtx.Lock()
			fmt.Println("GoRoutines:", runtime.NumGoroutine())
			myctr := ctr
			myctr++
			ctr = myctr
			mtx.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(ctr)
}
