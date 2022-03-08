package main

import (
	"fmt"
	"modules-demo/calculator"
)

func main() {
	fmt.Println(calculator.Add(100, 200))
	fmt.Println(calculator.Subtract(100, 200))
	fmt.Println("Operation Count = ", calculator.OpCount())
}
