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
	//var pen Product
	//var pen Product = Product{100, "Pen", 10, 100, "Stationary"}
	//var pen = Product{100, "Pen", 10, 100, "Stationary"}
	//pen := Product{100, "Pen", 10, 100, "Stationary"}
	//pen := Product{Id: 100, Name: "Pen", Cost: 10, Units: 100}
	/*
		pen := Product{
			Id:    100,
			Name:  "Pen",
			Cost:  10,
			Units: 100,
		}
	*/
	/*
		pen := &Product{
			Id:    100,
			Name:  "Pen",
			Cost:  10,
			Units: 100,
		}

		var p = new(Product)
	*/

	pencil := &Product{
		Id:       100,
		Name:     "Pencil",
		Cost:     5,
		Units:    100,
		Category: "Stationary",
	}
	//fmt.Println((*pencil).Name)
	pencil.Name = "Industrial Pencil"
	fmt.Println(pencil.Name)

	pen := Product{
		Id:       100,
		Name:     "Pen",
		Cost:     10,
		Units:    100,
		Category: "Stationary",
	}
	//fmt.Printf("%#v\n", pen)
	fmt.Println(Format(pen))
	fmt.Println("After applying 10% discount")
	ApplyDiscount(&pen, 10)
	fmt.Println(Format(pen))
}

func Format(p Product) string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %f, Units = %d, Category = %s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

func ApplyDiscount(product *Product, discountPercentage float32) {
	//(*product).Cost = (*product).Cost * ((100 - discountPercentage) / 100)
	product.Cost = product.Cost * ((100 - discountPercentage) / 100)
}
