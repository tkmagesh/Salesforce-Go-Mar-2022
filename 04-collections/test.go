package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	nos := make([]int, n, n+10)
	fmt.Printf("len = %d, cap = %d, arr = %v, addr = %p\n", len(nos), cap(nos), nos, nos)
	nos2 := append(nos, 12)
	nos[2] = 9999
	fmt.Printf("[nos] len = %d, cap = %d, arr = %v, addr = %p\n", len(nos), cap(nos), nos, nos)
	fmt.Printf("[nos2] len = %d, cap = %d, arr = %v, addr = %p\n", len(nos2), cap(nos2), nos2, nos2)

	fmt.Println("&nos =", &nos, "&nos2 =", &nos2)
}
