package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main started")
	ch := make(chan int)
	go genEvenNos(10, ch)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println("main completed")
}

func genEvenNos(count int, ch chan int) {
	for i := 1; i <= count; i++ {
		time.Sleep(500 * time.Millisecond)
		ch <- i * 2
	}
}
