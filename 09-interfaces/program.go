package main

import (
	"fmt"
	"math"
)

//Day-1
type Circle struct {
	Radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

//Day-2
type Rectangle struct {
	Height float32
	Width  float32
}

func (r Rectangle) Area() float32 {
	return r.Height * r.Width
}

//Day-3
type ShapeWithArea interface {
	Area() float32
}

//
func PrintArea(sa ShapeWithArea) {
	fmt.Printf("Area = %f\n", sa.Area())
}

func main() {
	c := Circle{Radius: 12}
	PrintArea(c)

	r := Rectangle{Height: 10, Width: 12}
	PrintArea(r)
}
