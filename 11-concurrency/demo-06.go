package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int64

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go fn(wg)
	}
	wg.Wait()
	fmt.Println(counter)
}

func fn(wg *sync.WaitGroup) {
	atomic.AddInt64(&counter, 1)
	wg.Done()
}
