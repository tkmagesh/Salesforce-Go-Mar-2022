package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

type Dummy struct {
	Name string
}
type PerishableProduct struct {
	Product
	Expiry string
}

func NewPerishableProduct(id int, name string, cost float32, units int, category string, expiry string) *PerishableProduct {
	return &PerishableProduct{
		Product: Product{
			Id:       id,
			Name:     name,
			Cost:     cost,
			Units:    units,
			Category: category,
		},
		Expiry: expiry,
	}
}

func main() {
	//apple := PerishableProduct{Product{Id: 200, Name: "Apple", Cost: 20, Units: 10, Category: "Fruits"}, "2 Days"}
	apple := PerishableProduct{
		Product: Product{Id: 200, Name: "Apple", Cost: 20, Units: 10, Category: "Fruits"},
		Expiry:  "2 Days",
	}

	fmt.Println(apple)

	//accessing the attributes
	fmt.Println(apple.Product.Cost)
	fmt.Println(apple.Cost)

	fmt.Println(apple.Product.Name)
	fmt.Println(apple.Name)

	grapes := NewPerishableProduct(201, "Grapes", 40, 10, "Fruits", "4 Days")
	fmt.Println(grapes)

	//fmt.Println(Format(apple.Product))
	fmt.Println(FormatPerishableProduct(apple))
	ApplyDiscount(&apple.Product, 10)
	fmt.Println("After applying 10% discount")
	//fmt.Println(Format(apple.Product))
	fmt.Println(FormatPerishableProduct(apple))
}

func Format(p Product) string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %f, Units = %d, Category = %s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

func ApplyDiscount(product *Product, discountPercentage float32) {
	product.Cost = product.Cost * ((100 - discountPercentage) / 100)
}

func FormatPerishableProduct(pp PerishableProduct) string {
	return fmt.Sprintf("%s, Expiry =%s", Format(pp.Product), pp.Expiry)
}
