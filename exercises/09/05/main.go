package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {
	var ctr int64 = 0
	const gs = 100
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&ctr, 1)
			runtime.Gosched()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(ctr)
}
