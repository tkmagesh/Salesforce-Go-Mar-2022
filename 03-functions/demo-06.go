/* Handling errors */
package main

import (
	"errors"
	"fmt"
)

var DivideByZeroError = errors.New("divisor cannot be zero")

func main() {
	q, r, err := divide(100, 0)
	if err == DivideByZeroError {
		fmt.Println("Invalid input. Cannot divide by zero")
		return
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(q, r)
}

/*
func divide(x, y int) (int, int, error) {
	if y == 0 {
		//creating error
		err := errors.New("divisor cannot be zero")
		return 0, 0, err
	}
	quotient := x / y
	remainder := x % y
	return quotient, remainder, nil
}
*/

func divide(x, y int) (quotient, remainder int, err error) {
	if y == 0 {
		//creating error
		err = DivideByZeroError
		return
	}
	quotient = x / y
	remainder = x % y
	return
}
