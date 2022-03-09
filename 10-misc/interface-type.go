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
	var x interface{}
	//x = 100
	x = "abc"
	//x = true
	//x = 19.99
	//x = struct{}{}
	//y := x + 200

	if val, ok := x.(int); ok {
		fmt.Println(val + 200)
	} else {
		fmt.Println("x is not an int")
	}

	//x = 100
	//x = "This is a string"
	//x = true
	//x = []int{3, 1, 4, 2, 5}
	//x = 19.00
	x = Product{Name: "Phone"}
	switch val := x.(type) {
	case int:
		fmt.Println("x is an int, x + 100 = ", val+100)
	case string:
		fmt.Println("x is a string, len(x) = ", len(val))
	case bool:
		fmt.Println("x is a bool, x = ", val)
	case []int:
		fmt.Println("x is an int slice, len(x) = ", len(val))
	case Product:
		fmt.Println("x is a product with Name :", val.Name)
	default:
		fmt.Println("unknown type")
	}

}
