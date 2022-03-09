package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("main started")
	/*
		var resultCh chan int
		resultCh = make(chan int)
	*/
	resultCh := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go add(100, 200, resultCh, wg)
	result := <-resultCh //recieving the data from the channel
	wg.Wait()
	fmt.Printf("result = %d\n", result)
	fmt.Println("main completed")
}

func add(x, y int, resultCh chan int, wg *sync.WaitGroup) {
	time.Sleep(5 * time.Second)
	result := x + y
	resultCh <- result //sending data to the channle
	wg.Done()
}
