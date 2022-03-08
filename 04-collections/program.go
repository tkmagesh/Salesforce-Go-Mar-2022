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
	fmt.Println("=> Iterating using indexer")
	for i := 0; i < len(nos); i++ {
		fmt.Println(nos[i])
	}

	fmt.Println("=> Iterating using range")
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

	fmt.Println("=> adding new values")
	//products = append(products, "Notebook")
	//products = append(products, "Notebook", "Stapler")
	newProducts := []string{"Notebook", "Stapler"}
	products = append(products, newProducts...)
	products[3] = "Stylus"
	fmt.Println(products)

	//iterating an slice (using indexer)
	fmt.Println("=> Iterating using indexer")
	for i := 0; i < len(products); i++ {
		fmt.Println(products[i])
	}

	fmt.Println("=> Iterating using range")
	for idx, product := range products {
		fmt.Printf("products[%d] = %s\n", idx, product)
	}

	//slicing
	/*
		[lo : hi] => from lo to hi-1
		[lo : ] => from lo to end
		[ : hi] => from 0 to hi-1
		[lo : lo] => empty slice
		[lo : lo+1] => single element slice with [lo]
		[:] => copy of the slice

	*/

	fmt.Println("products[2:4] => ", products[2:4])
	fmt.Println("products[:4] => ", products[:4])
	fmt.Println("products[2:] => ", products[2:])

	//printProducts(products...)
	//printProducts("Pen", "Pencil")

	fmt.Printf("\nMaps\n")
	//var productRank map[string]int = make(map[string]int)
	//productRank := map[string]int{"Pen": 4, "Pencil": 2, "Marker": 3}
	productRank := map[string]int{
		"Pen":    4,
		"Pencil": 2,
		"Marker": 3,
	}
	fmt.Println(productRank)
	fmt.Println("=> Adding a new item")
	productRank["Stylus"] = 1
	fmt.Println(productRank)

	fmt.Println("=> Iterating a map")
	for key, value := range productRank {
		fmt.Printf("productRank[%q] = %d\n", key, value)
	}

	fmt.Println("=>  # of items in map = ", len(productRank))

	fmt.Println("=> Rank of Pen is : ", productRank["Pen"])

	fmt.Println("=> Checking for the existence of a key")
	keyToCheck := "Mouse"
	if rank, exists := productRank[keyToCheck]; exists {
		fmt.Printf("Rank of %q is : %d\n", keyToCheck, rank)
	} else {
		fmt.Printf("key : %q , does not exist, rank = %d\n", keyToCheck, rank)
	}

	fmt.Println("=> Removing a key/value from a map")
	delete(productRank, "Stylus")
	fmt.Println("After deleting \"stylus\", productRank = ", productRank)

}

func printProducts(products ...string) {
	fmt.Println("List of products")
	for _, product := range products {
		fmt.Println(product)
	}
}
