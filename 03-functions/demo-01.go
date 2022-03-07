package main

import "fmt"

func main() {
	sayHi()
	greet("Hello there!")
	fmt.Printf("Is 22 even: ? %t\n", isEven(22))
	fmt.Println(add(100, 200))
	//fmt.Println(divide(100, 7))
	/*
		quotient, remainder := divide(100, 7)
		fmt.Println(quotient, remainder)
	*/
	/*
		quotient, _ := divide(100, 7)
		fmt.Println(quotient)
	*/
	_, remainder := divide(100, 7)
	fmt.Println(remainder)
}

func sayHi() {
	fmt.Println("Hi")
}

func greet(msg string) {
	fmt.Println(msg)
}

func isEven(n int) bool {
	return n%2 == 0
}

/*
func add(x int, y int) int {
	return x + y
}
*/

func add(x, y int) int {
	return x + y
}

/*
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
*/

func divide(x, y int) (quotient int, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}
