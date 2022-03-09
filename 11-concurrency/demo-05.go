package main

import (
	"fmt"
	"sync"
)

type OpCounter struct {
	counter int
	sync.Mutex
}

func (opCounter *OpCounter) Increment() {
	opCounter.Lock()
	{
		opCounter.counter++
	}
	opCounter.Unlock()
}

var opCounter OpCounter

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go fn(wg)
	}
	wg.Wait()
	fmt.Println(opCounter.counter)
}

func fn(wg *sync.WaitGroup) {
	//counter++
	opCounter.Increment()
	wg.Done()
}
