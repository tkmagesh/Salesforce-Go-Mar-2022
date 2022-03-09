package main

import (
	"fmt"
	"time"
)

//consumer
func main() {
	fmt.Println("main started")
	done := make(chan bool)
	ch := genEvenNos(done)
	go func() {
		var input string
		fmt.Scanln(&input)
		done <- true
	}()
	for evenNo := range ch {
		fmt.Println(evenNo)
	}
	fmt.Println("main completed")
}

//producer
func genEvenNos(done chan bool) <-chan int {
	ch := make(chan int)
	/*
		timeout := make(chan bool)
		go func() {
			time.Sleep(5 * time.Second)
			timeout <- true
		}() */
	timeout := time.After(5 * time.Second)
	go func() {
		var i = 1
	LOOP:
		for {
			select {
			case ch <- i * 2:
				i++
				time.Sleep(500 * time.Millisecond)
			case <-done:
				break LOOP
			case <-timeout:
				break LOOP
			}
		}
		close(ch)
	}()
	return ch
}
