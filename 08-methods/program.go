package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func main() {
	pen := &Product{
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
}

func (p Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %f, Units = %d, Category = %s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

func (product *Product) ApplyDiscount(discountPercentage float32) {
	//(*product).Cost = (*product).Cost * ((100 - discountPercentage) / 100)
	product.Cost = product.Cost * ((100 - discountPercentage) / 100)
}
