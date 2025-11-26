package main

import "fmt"

func main() {
	// Explicit type
	var name string = "Ramesh"
	fmt.Println(name)
	var age int = 30
	fmt.Println(age)
	var isEmployed bool = true
	fmt.Println(isEmployed)
	var salary float64 = 75000.50
	fmt.Println(salary)

	// Inferred type
	var city = "New York"
	fmt.Println(city)

	// Short declaration
	country := "USA"
	fmt.Println(country)

	var myName string
	myName = "Ramesh"
	fmt.Println(myName)
}
