package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/*
		Desired Output:
			Hello
			World
			Hello
			World
			Hello
			World
			Hello
			World
			Hello
			World

		Important:
			The loop should be in the "print" function

	*/
	wg := &sync.WaitGroup{}
	ch1 := make(chan string, 1)
	ch2 := make(chan string)
	wg.Add(2)
	go print("Hello", ch1, ch2, wg)
	go print("World", ch2, ch1, wg)
	ch1 <- "start"
	wg.Wait()
}

func print(msg string, in, out chan string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		<-in
		fmt.Println(msg)
		time.Sleep(500 * time.Millisecond)
		out <- "done"
	}
	wg.Done()
}
