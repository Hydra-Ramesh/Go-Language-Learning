package main

import "fmt"

func main() {
	// Using continue and break in loops
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println("Iteration:", i)
	}

	for j := 0; j < 5; j++ {
		if j == 3 {
			break
		}
		fmt.Println("Count:", j)
	}
}
