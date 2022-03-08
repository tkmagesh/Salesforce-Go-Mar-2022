package main

import (
	"fmt"
	calc "modules-demo/calculator"
	"modules-demo/calculator/utils"

	"github.com/fatih/color"
)

func main() {
	fmt.Println(calc.Add(100, 200))
	fmt.Println(calc.Subtract(100, 200))
	fmt.Println("Operation Count = ", calc.OpCount())
	fmt.Println(utils.IsEven(20))
	color.Red("This line will be printed in Red")

}
