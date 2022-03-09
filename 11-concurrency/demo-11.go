package main

import (
	"fmt"
	"time"
)

//consumer
func main() {
	fmt.Println("main started")
	ch := genEvenNos(10)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println("main completed")
}

//producer
func genEvenNos(count int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 1; i <= count; i++ {
			time.Sleep(500 * time.Millisecond)
			ch <- i * 2
		}
	}()
	return ch
}
