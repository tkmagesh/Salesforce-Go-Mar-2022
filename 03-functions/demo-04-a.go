/*
Higher order functions
	* passing functions as arguments to other functions
*/
package main

import "fmt"

func main() {
	exec(f1)
	exec(f2)
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}

func exec(fn func()) {
	fn()
}
