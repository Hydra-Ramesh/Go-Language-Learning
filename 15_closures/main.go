package main

import "fmt"

// counter returns a closure that increments and returns an integer value each time it is called.
func counter() func() int {
	var count int = 0
	return func() int {
		count += 1
		return count
	}
}
func main() {
	increment := counter()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}
