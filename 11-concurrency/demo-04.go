package main

import (
	"fmt"
	"sync"
)

var counter int
var mutex sync.Mutex

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
	//counter++
	mutex.Lock()
	{
		counter = counter + 1
	}
	mutex.Unlock()
	wg.Done()
}
