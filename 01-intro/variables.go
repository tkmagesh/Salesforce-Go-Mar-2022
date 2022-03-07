package main

import "fmt"

func main() {
	/*
		var msg string
		msg = "Hello World!"
	*/

	/*
		var msg string = "Hello World!"
	*/

	/*
		//Type inference
		var msg = "Hello World!"
	*/
	msg := "Hello World!" //=> applicable ONLY in a function (not at the package level)
	fmt.Println(msg)

	//Multiple variables
	/*
		var x int
		var y int
		var result int
		var str string

		x = 100
		y = 200
		str = "Result ="
		result = x + y
		fmt.Println(str, result)
	*/
	/*
		var x, y, result int
		var str string
		x = 100
		y = 200
		result = x + y
		str = "Result ="
		fmt.Println(str, result)
	*/

	/*
		var (
			x, y, result int
			str          string
		)
		x = 100
		y = 200
		result = x + y
		str = "Result ="
		fmt.Println(str, result)
	*/
	/*
		var x, y int = 100, 200
		var str string = "Result ="
		var result int = x + y
		fmt.Println(str, result)
	*/

	/*
		var (
			x, y   int    = 100, 200
			str    string = "Result ="
			result int    = x + y
		)
		fmt.Println(str, result)
	*/

	/*
		var (
			x, y   = 100, 200
			str    = "Result ="
			result = x + y
		)
		fmt.Println(str, result)
	*/

	/*
		var (
			x, y, str = 100, 200, "Result ="
			result    = x + y
		)
		fmt.Println(str, result)
	*/

	x, y, str := 100, 200, "Result ="
	result := x + y
	fmt.Println(str, result)
}
