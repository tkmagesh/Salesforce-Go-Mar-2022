package main

import (
	"fmt"
	"methods-demo/models"
)

func main() {
	pen := &models.Product{
		Id:       100,
		Name:     "Pen",
		Cost:     10,
		Units:    100,
		Category: "Stationary",
	}
	//fmt.Printf("%#v\n", pen)
	//fmt.Println(Format(pen))
	fmt.Println(pen.Format())
	fmt.Println("After applying 10% discount")
	//ApplyDiscount(&pen, 10)
	//(&pen).ApplyDiscount(10)
	pen.ApplyDiscount(10)
	fmt.Println(pen.Format())

	//composition (methods)
	grapes := models.NewPerishableProduct(201, "Grapes", 40, 10, "Fruits", "4 Days")
	//fmt.Println(grapes.Product.Format())
	fmt.Println(grapes.Format())
}
