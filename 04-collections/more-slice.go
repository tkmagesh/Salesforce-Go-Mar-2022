package main

import "fmt"

func main() {
	/*
		nos := []int{10, 20, 30, 40}
		fmt.Printf("len = %d, cap = %d, nos = %v\n", len(nos), cap(nos), nos)
		fmt.Println("After appending")
		nos = append(nos, 50)
		fmt.Printf("len = %d, cap = %d, nos = %v\n", len(nos), cap(nos), nos)
	*/

	nos := make([]int, 3, 10)
	nos[0] = 10
	nos[1] = 20
	nos[2] = 30
	nos = append(nos, 40)
	fmt.Printf("len = %d, cap = %d, nos = %v\n", len(nos), cap(nos), nos)

	nos = append(nos, 50)
	nos = append(nos, 60)
	nos = append(nos, 70)
	nos = append(nos, 80)
	nos = append(nos, 90)
	nos = append(nos, 100)
	fmt.Printf("len = %d, cap = %d, nos = %v\n", len(nos), cap(nos), nos)

	nos = append(nos, 110)
	fmt.Printf("len = %d, cap = %d, nos = %v\n", len(nos), cap(nos), nos)
	nos = append(nos, 120)
	nos = append(nos, 130)
	nos = append(nos, 130)
	nos = append(nos, 150)
	nos = append(nos, 160)
	nos = append(nos, 170)
	nos = append(nos, 180)
	nos = append(nos, 190)
	nos = append(nos, 200)
	fmt.Printf("len = %d, cap = %d, nos = %v\n", len(nos), cap(nos), nos)

	nos = append(nos, 210)
	fmt.Printf("len = %d, cap = %d, nos = %v\n", len(nos), cap(nos), nos)
}
