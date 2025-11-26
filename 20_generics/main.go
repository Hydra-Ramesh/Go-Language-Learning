package main

import "fmt"

func printSlice[T any](item []T) {
	for _, v := range item {
		fmt.Println(v)
	}
}

func limitPrintSlice[T int | string](item []T) {
	for _, v := range item {
		fmt.Println(v)
	}
}

type Stack[T any] struct {
	items []T
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	printSlice(nums)
	fmt.Println("Slice printed successfully")

	names := []string{"Alice", "Bob", "Charlie"}
	printSlice(names)
	fmt.Println("Slice printed successfully")
	limitPrintSlice(nums)
	limitPrintSlice(names)

	// This will cause a compile-time error
	// 	valus:= []bool {true, false, true}

	myStack := Stack[string]{
		items: []string{"first", "second", "third"},
	}
	fmt.Println("Stack items:")
	printSlice(myStack.items)
}
