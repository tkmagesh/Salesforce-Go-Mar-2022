/*
Higher order functions
	* passing functions as arguments to other functions
*/
package main

import "fmt"

func main() {

	add(100, 200)
	subtract(100, 200)

	log(100, 200, add)
	log(100, 200, subtract)

	/* create 2 functions which when invoked will perform add / subtract with logging */

}

func add(x, y int) {
	fmt.Println("Add result =", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract result =", x-y)
}

func log(x, y int, operation func(int, int)) {
	fmt.Println("invocation started")
	operation(x, y)
	fmt.Println("invocation completed")
}
