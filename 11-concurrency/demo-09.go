package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main started")
	/*
		var resultCh chan int
		resultCh = make(chan int)
	*/
	resultCh := make(chan int)
	go add(100, 200, resultCh)
	result := <-resultCh //recieving the data from the channel
	fmt.Printf("result = %d\n", result)
	fmt.Println("main completed")
}

func add(x, y int, resultCh chan int) {
	time.Sleep(5 * time.Second)
	result := x + y
	resultCh <- result //sending data to the channle
}
