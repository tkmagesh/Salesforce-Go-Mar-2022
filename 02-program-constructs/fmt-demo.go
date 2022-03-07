package main

import "fmt"

func main() {
	//x, y := 100, 200
	var x, y int

	/*
		fmt.Println("Enter the value of x :")
		fmt.Scanln(&x)
		fmt.Println("Enter the value of y :")
		fmt.Scanln(&y)
	*/
	fmt.Println("Enter the values of x and y :")
	//fmt.Scanln(&x, &y)
	fmt.Scanf("%d:%d\n", &x, &y)
	result := x + y
	//fmt.Printf("x = %d, y = %d, x + y = %d\n", x, y, result)
	s := fmt.Sprintf("x = %d, y = %d, x + y = %d\n", x, y, result)
	fmt.Print(s)
}
