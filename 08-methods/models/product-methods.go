package models

import "fmt"

func (p Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %f, Units = %d, Category = %s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

func (product *Product) ApplyDiscount(discountPercentage float32) {
	//(*product).Cost = (*product).Cost * ((100 - discountPercentage) / 100)
	product.Cost = product.Cost * ((100 - discountPercentage) / 100)
}
