package main

import "fmt"

type MyStr string

func (ms MyStr) Length() int {
	return len(ms)
}

func main() {
	str := MyStr("This is my string")
	fmt.Println(str.Length())
}
