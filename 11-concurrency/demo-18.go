package main

import (
	"fmt"
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
	print("Hello")
	print("World")
}

func print(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg)
		time.Sleep(500 * time.Millisecond)
	}
}
