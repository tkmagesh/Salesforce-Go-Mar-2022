package main

import "fmt"

func main() {
	//if statement
	/*
		no := 21
		if no%2 == 0 {
			fmt.Println("no is even")
		} else {
			fmt.Println("no is odd")
		}
	*/

	if no := 21; no%2 == 0 {
		fmt.Println("no is even")
	} else {
		fmt.Println("no is odd")
	}

	//for construct
	/* ver 1.0 */
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("sum = ", sum)

	/* ver 2.0 (imitating the 'while' construct) */
	numSum := 1
	for numSum < 100 {
		numSum += numSum
	}
	fmt.Println("numSum =", numSum)

	/* ver 3.0 (infinite loop) */
	x := 1
	for {
		x += x
		if x > 100 {
			break
		} else {
			continue
		}
	}
	fmt.Println("x =", x)

	/* switch case */
	/*
		rank by score
		score 0 - 3 => "terrible"
		score 4 - 7 => "not bad"
		score 8 - 9 => "good"
		score 10 => "excellent"
		otherwise => "invalid score"
	*/

	score := 6
	/* switch score {
	case 0:
		fmt.Println("terrible")
	case 1:
		fmt.Println("terrible")
	case 2:
		fmt.Println("terrible")
	case 3:
		fmt.Println("terrible")
	case 4:
		fmt.Println("not bad")
	case 5:
		fmt.Println("not bad")
	case 6:
		fmt.Println("not bad")
	case 7:
		fmt.Println("not bad")
	case 8:
		fmt.Println("good")
	case 9:
		fmt.Println("good")
	case 10:
		fmt.Println("excellent")
	default:
		fmt.Println("invalid score")
	}
	*/
	score = 2
	switch score {
	case 0, 1, 2, 3:
		fmt.Println("terrible")
	case 4, 5, 6, 7:
		fmt.Println("not bad")
	case 8, 9:
		fmt.Println("good")
	case 10:
		fmt.Println("excellent")
	default:
		fmt.Println("invalid score")
	}

	num := 21
	switch {
	case num%2 == 0:
		fmt.Println("num is even")
	case num%2 == 1:
		fmt.Println("num is odd")
	}

	//fallthrough
	n := 4
	switch n {
	case 0:
		fmt.Println("is zero")
	case 1:
		fmt.Println("n <= 1")
		fallthrough
	case 2:
		fmt.Println("n <= 2")
		fallthrough
	case 3:
		fmt.Println("n <= 3")
		fallthrough
	case 4:
		fmt.Println("n <= 4")
		fallthrough
	case 5:
		fmt.Println("n <= 5")
		fallthrough
	case 6:
		fmt.Println("n <= 6")
	case 7:
		fmt.Println("n <= 7")
		fallthrough
	case 8:
		fmt.Println("n <= 8")
	}
}
