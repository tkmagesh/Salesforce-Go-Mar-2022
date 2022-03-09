package main

import (
	"fmt"
	"math"
)

/* Perimeter
Circle = 2 *Pi *R
Rectangle = 2 * (height + Width)
*/

//Day-1
type Circle struct {
	Radius float32
}

func (c Circle) Area() float32 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float32 {
	return 2 * math.Pi * c.Radius
}

//Day-2
type Rectangle struct {
	Height float32
	Width  float32
}

func (r Rectangle) Area() float32 {
	return r.Height * r.Width
}

func (r Rectangle) Perimeter() float32 {
	return 2 * (r.Height + r.Width)
}

//Day-3
type ShapeWithArea interface {
	Area() float32
}

func PrintArea(sa ShapeWithArea) {
	fmt.Printf("Area = %f\n", sa.Area())
}

//Day-4
type ShapeWithPerimeter interface {
	Perimeter() float32
}

func PrintPerimeter(sp ShapeWithPerimeter) {
	fmt.Printf("Perimeter = %f\n", sp.Perimeter())
}

//Day-5
/* Interface composition */
type Shape interface {
	ShapeWithArea
	ShapeWithPerimeter
}

func PrintShape(s Shape) {
	PrintArea(s)
	PrintPerimeter(s)
}

func main() {
	c := Circle{Radius: 12}
	/*
		PrintArea(c)
		PrintPerimeter(c)
	*/
	PrintShape(c)

	r := Rectangle{Height: 10, Width: 12}
	/*
		PrintArea(r)
		PrintPerimeter(r)
	*/
	PrintShape(r)

}
