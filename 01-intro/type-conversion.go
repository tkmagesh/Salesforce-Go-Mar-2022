package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int8 = 100
	var y int16
	var s string
	y = int16(x)
	//s = strconv.FormatInt(int64(x), 10)
	s = strconv.Itoa(int(x))
	fmt.Println(x, y, s)
}
