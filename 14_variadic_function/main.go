package main

import "fmt"

func sum(num ...int) int {
	total := 0
	for _, v := range num {
		total += v
	}
	return total
}
func main() {
	fmt.Println(sum(1, 2, 3, 4))
	res := sum(10, 20, 30)
	fmt.Println(res)

	arr := []int{5, 10, 15}
	fmt.Println(sum(arr...))
}
