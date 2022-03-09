package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	ch <- 100
	no := <-ch
	fmt.Println(no)
}
