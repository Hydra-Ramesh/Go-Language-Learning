package main

import "fmt"

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
func main() {
	num1 := 6
	num2 := 8
	swap(&num1, &num2)
	fmt.Println("num1:", num1)
	fmt.Println("num2:", num2)
	fmt.Println(&num1, &num2)
}
