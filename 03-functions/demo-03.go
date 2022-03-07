/*
higher order functions
	* functions can be assigned as values to variables
*/
package main

import "fmt"

func main() {
	func() {
		fmt.Println("anonymous function invoked")
	}()

	/* var fn = func() {
		fmt.Println("fn invoked")
	} */
	var fn func()
	fn = func() {
		fmt.Println("fn invoked")
	}
	fn()

	var add func(int, int) int
	add = func(x int, y int) int {
		return x + y
	}
	fmt.Println(add(100, 200))
}
