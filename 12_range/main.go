package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for i, num := range nums {
		sum += num
		fmt.Println(num, i)
	}
	fmt.Println("Sum:", sum)

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
