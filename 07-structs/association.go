package main

import "fmt"

type Organization struct {
	Id   string
	Name string
	City string
}

type Employee struct {
	Id   int
	Name string
	Org  *Organization
}

func main() {
	salesForce := &Organization{
		Id:   "Org-101",
		Name: "Salesforce",
		City: "Bangalore",
	}
	emp1 := Employee{
		Id:   100,
		Name: "Magesh",
		Org:  salesForce,
	}
	emp2 := Employee{
		Id:   100,
		Name: "Magesh",
		Org:  salesForce,
	}
	//salesForce.City = "New York"
	//emp1.Org.City = "NewYork"
	fmt.Printf("%#v\n", emp1)
	fmt.Printf("%#v\n", emp2)
	emp1.Org.City = "New York"
	fmt.Println(emp1.Org.City)
	fmt.Println(emp2.Org.City)
}
