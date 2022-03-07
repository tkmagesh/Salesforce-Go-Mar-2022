/*
Higher order functions
	* Functions can be returned as return values
*/
package main

import "fmt"

func main() {
	fn := getFn()
	fn()
	add := getAdder()
	fmt.Println(add(100, 200))

	adderFor100 := getAdderFor(100)
	fmt.Println(adderFor100(200))

	addOp := getOperationFor(100, add)
	fmt.Println(addOp(100))

	subOp := getOperationFor(100, subtract)
	fmt.Println(subOp(200))
}

func getFn() func() {
	return func() {
		fmt.Println("fn invoked")
	}
}

func getAdder() func(int, int) int {
	return func(x, y int) int {
		return x + y
	}
}

func getAdderFor(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func getOperationFor(x int, oper func(int, int) int) func(int) int {
	return func(y int) int {
		return oper(x, y)
	}
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}
