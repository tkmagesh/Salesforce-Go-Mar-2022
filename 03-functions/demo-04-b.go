/*
Higher order functions
	* passing functions as arguments to other functions
*/
package main

import "fmt"

func main() {
	/*
		add(100, 200)
		subtract(100, 200)
	*/
	/*
		Expected result:
			invocation started
			Add result = 300
			invocation completed

			invocation started
			Subtract result = -100
			invocation completed
	*/
	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/
	/*
		log(100, 200, add)
		log(100, 200, subtract)
	*/
	//logFns(100, 200, add, subtract)

	genericLogFn(func() {
		fmt.Println("Operation started")
	}, 100, 200, add, subtract)
}

func add(x, y int) {
	fmt.Println("Add result =", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract result =", x-y)
}

/*
func logAdd(x, y int) {
	fmt.Println("invocation started")
	add(x, y)
	fmt.Println("invocation completed")
}

func logSubtract(x, y int) {
	fmt.Println("invocation started")
	subtract(x, y)
	fmt.Println("invocation completed")
}
*/

func log(x, y int, operation func(int, int)) {
	fmt.Println("invocation started")
	operation(x, y)
	fmt.Println("invocation completed")
}

func logFns(x, y int, operations ...func(int, int)) {
	/*
		for _, operation := range operations {
			fmt.Println("invocation started")
			operation(x, y)
			fmt.Println("invocation completed")
		}
	*/
	for _, operation := range operations {
		log(x, y, operation)
	}
}

func genericLogFn(logOperation func(), x, y int, operations ...func(int, int)) {
	for _, operation := range operations {
		logOperation()
		operation(x, y)
	}
}
