package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("main started")
	wg.Add(2)
	go f1(1)
	f1(2)
	f2()
	wg.Wait() //execution is blocked until the wg counter becomes 0
	fmt.Println("main completed")
}

func f1(id int) {
	fmt.Printf("f1[%d] started\n", id)
	time.Sleep(5 * time.Second)
	fmt.Printf("f1[%d] completed\n", id)
	wg.Done()
}

func f2() {
	fmt.Println("f2 invoked")
}
