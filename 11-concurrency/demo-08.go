package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		fmt.Println("anonymous start")
		fmt.Println("[anon] attempting to write the data from the channel")
		ch <- 100
		fmt.Println("[anon] attempt to write the data from the channel is a success")
		fmt.Println("anonymous end")
	}()
	fmt.Println("attempting to read the data from the channel")
	no := <-ch
	fmt.Println(no)
}
