package main

import "fmt"

func main() {
	var no int = 100
	var noPtr *int //variable to hold the address of an int value

	noPtr = &no //assign the address of an int value
	fmt.Println(noPtr)

	/* dereferncing - accessing the underlying using the address */
	var no2 int = *noPtr
	fmt.Println(no2)

	fmt.Println("Before incrementing, no = ", no)
	increment(&no)
	fmt.Println("After incrementing, no = ", no)

	n1, n2 := 10, 20
	fmt.Printf("Before swaping, n1 = %d & n2 = %d\n", n1, n2)
	swap(&n1, &n2)
	fmt.Printf("Before swaping, n1 = %d & n2 = %d\n", n1, n2)
}

func increment(no *int) {
	*no++
}

func swap(x, y *int) /* No return values */ {
	var temp = *x
	*x = *y
	*y = temp
	//*x, *y = *y, *x
}
