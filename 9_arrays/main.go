package main

import "fmt"

func main() {
	var arr [5]int
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	arr[3] = 40
	arr[4] = 50
	fmt.Println("Array elements:", arr)

	var boolArr [3]bool
	boolArr[0] = true
	boolArr[1] = false
	boolArr[2] = true
	fmt.Println("Boolean array elements:", boolArr)

	var name [3]string
	name[0] = "Go"
	name[1] = "Python"
	name[2] = "Java"
	fmt.Println("String array elements:", name)

	// Default values in all types are zero for numeric types, false for boolean, and empty string for strings
	// Intializing an array without assigning values
	nums := [4]int{1, 2, 3, 4}
	fmt.Println("Initialized array elements:", nums)

	//2D Array
	matrix := [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("2D array elements:", matrix)
}
