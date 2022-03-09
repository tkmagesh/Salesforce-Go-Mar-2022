package main

import (
	"fmt"
	"sort"
)

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

func (p Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %f, Units = %d, Category = %s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

type Products []Product

func (products Products) Print() {
	for _, product := range products {
		fmt.Println(product.Format())
	}
}

//sort.Interface implementation
func (products Products) Len() int {
	return len(products)
}

func (products Products) Less(i, j int) bool {
	return products[i].Id < products[j].Id
}

func (products Products) Swap(i, j int) {
	products[i], products[j] = products[j], products[i]
}

type ByName struct {
	Products
}

func (byName ByName) Less(i, j int) bool {
	return byName.Products[i].Name < byName.Products[j].Name
}

/*
func Sort(data sort.Interface) {
	for i := 0; i < data.Len()-1; i++ {
		for j := i + 1; j < data.Len(); j++ {
			if !data.Less(i, j) {
				data.Swap(i, j)
			}
		}
	}
}
*/

func main() {
	products := Products{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Utencil"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
	}

	fmt.Println(" => Initial List")
	products.Print()

	fmt.Println(" => Default Sort")
	sort.Sort(products)
	products.Print()

	fmt.Println(" => Sort by Name")
	sort.Sort(ByName{products})
	products.Print()

}
