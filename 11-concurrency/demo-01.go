package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main started")
	go f1()
	f2()
	time.Sleep(1 * time.Millisecond) // DO NOT DO THIS
	/* DO NOT DO THE FOLLOWING */
	/*
		var input string
		fmt.Scanln(&input)
	*/
	fmt.Println("main completed")
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(5 * time.Millisecond)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
