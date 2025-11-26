package main

import "fmt"

// for -> only loop in Go
func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// Classic for loop
	for j := 1; j <= 10; j++ {
		fmt.Println(j)
	}

	// for range loop
	for i := range 10 {
		fmt.Println(i)
	}
}
