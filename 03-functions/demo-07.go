/* deferred functions */

package main

import "fmt"

func main() {

	defer fmt.Println("[main] - deferred")
	fmt.Println("main started")
	f1()
	//divide(100, 0)
	fmt.Println("main completed")
}

func f1() {
	defer fmt.Println("[f1] - deferred - 1")
	defer func() {
		fmt.Println("[f1] - deferred - 2")
	}()
	defer func() {
		fmt.Println("[f1] - deferred - 3")
	}()
	fmt.Println("f1 started")
	f2()
	fmt.Println("f1 completed")
}

func f2() {
	defer func() {
		fmt.Println("[f2] - deferred - 1")
	}()
	defer func() {
		fmt.Println("[f2] - deferred - 2")
	}()
	defer func() {
		fmt.Println("[f2] - deferred - 3")
	}()
	fmt.Println("f2 started")
	fmt.Println("f2 completed")
}

func divide(x, y int) {
	defer func() {
		fmt.Println("[divide] - deferred")
	}()
	fmt.Println(x / y)
}
