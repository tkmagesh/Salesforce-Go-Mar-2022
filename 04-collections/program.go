package main

import "fmt"

func main() {
	//Arrays
	fmt.Printf("\nArray\n")
	//var nos [5]int
	//var nos [5]int = [5]int{3, 1, 4, 2, 5}
	//var nos = [5]int{3, 1, 4, 2, 5}
	//nos := [5]int{3, 1, 4, 2, 5}
	nos := [...]int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	//iterating an array (using indexer)
	fmt.Println("Iterating using indexer")
	for i := 0; i < len(nos); i++ {
		fmt.Println(nos[i])
	}

	fmt.Println("Iterating using range")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	//copying an array
	newNos := nos
	newNos[0] = 100
	fmt.Printf("newNos = %v, nos = %v\n", newNos, nos)

	fmt.Printf("\nSlice\n")
	//var products []string
	var products []string = []string{"Pen", "Pencil", "Marker"}
	fmt.Println(products)

	fmt.Println("adding new values")
	//products = append(products, "Notebook")
	//products = append(products, "Notebook", "Stapler")
	newProducts := []string{"Notebook", "Stapler"}
	products = append(products, newProducts...)
	products[3] = "Stylus"
	fmt.Println(products)

	//iterating an slice (using indexer)
	fmt.Println("Iterating using indexer")
	for i := 0; i < len(products); i++ {
		fmt.Println(products[i])
	}

	fmt.Println("Iterating using range")
	for idx, product := range products {
		fmt.Printf("products[%d] = %s\n", idx, product)
	}

	printProducts(products...)
	//printProducts("Pen", "Pencil")

}

func printProducts(products ...string) {
	fmt.Println("List of products")
	for _, product := range products {
		fmt.Println(product)
	}
}
