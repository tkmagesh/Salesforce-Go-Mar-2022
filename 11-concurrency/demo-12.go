package main

import (
	"fmt"
	"time"
)

//consumer
func main() {
	fmt.Println("main started")
	ch := genEvenNos()
	for {
		if evenNo, ok := <-ch; !ok {
			break
		} else {
			fmt.Println(evenNo)
		}
	}
	fmt.Println("main completed")
}

//producer
func genEvenNos() <-chan int {
	ch := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			time.Sleep(500 * time.Millisecond)
			ch <- i * 2
		}
		close(ch)
	}()
	return ch
}
