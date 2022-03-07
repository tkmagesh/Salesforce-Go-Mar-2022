package main

import "fmt"

func main() {
	var userChoice, n1, n2, result int

LOOP:
	for {
		fmt.Println("1. Add")
		fmt.Println("2. Subtract")
		fmt.Println("3. Multiply")
		fmt.Println("4. Divide")
		fmt.Println("5. Exit")
		fmt.Println("Enter your choice :")
		fmt.Scanln(&userChoice)
		if userChoice >= 1 && userChoice <= 4 {
			fmt.Println("Enter the numbers :")
			fmt.Scanln(&n1, &n2)
		}
		switch userChoice {
		case 1:
			result = n1 + n2
		case 2:
			result = n1 - n2
		case 3:
			result = n1 * n2
		case 4:
			result = n1 / n2
		case 5:
			break LOOP
		default:
			fmt.Println("Invalid choice")
			continue LOOP
		}
		fmt.Printf("Result = %d\n", result)
	}
	fmt.Println("Thank you!")
}
