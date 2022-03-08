package main

import (
	"errors"
	"fmt"
)

var IndexOutOfRangeError = errors.New("index out of range error")

func main() {
	defer func() {
		err := recover()
		if err == IndexOutOfRangeError {
			fmt.Println("Be careful iterating an array")
			return
		}
		if err != nil {
			fmt.Println("Panic occurred. Contact the adminstrator")
			fmt.Println(err)
			return
		}
		fmt.Println("Thank you!")
	}()
	//processNos()
	//fmt.Println(divide(100, 7))
	//fmt.Println(divide(100, 0))
	q, r, err := divide(100, 0)
	if err != nil {
		fmt.Println("error in divide function : ", err)
	} else {
		fmt.Printf("Quotient = %d, Remainder = %d\n", q, r)
	}
	processNos(false)
}

func divide(x, y int) (quotient, remainder int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	quotient = x / y
	remainder = x % y
	return
}

func processNos(shouldPanic bool) {
	if shouldPanic {
		panic(IndexOutOfRangeError)
	}
}
