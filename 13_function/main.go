package main

import "fmt"

// Function with parameters and return value
func add(a int, b int) int {
	return a + b
}

// Function returning multiple values
func getLang() (string, string, string) {
	return "Go", "Java", "Python"
}

// Function as parameter
func processIt(fn func(a int) int) {
	fn(5)
}

// Function returning a function
func peocess() func(a int) int {
	return func(a int) int {
		return 2
	}
}
func main() {
	result := add(2, 5)
	fmt.Println(result)
	lang1, lang2, lang3 := getLang()
	fmt.Println(lang1, lang2, lang3)

	// Anonymous function
	fn := func(a int) int {
		return 2
	}

	processIt(fn)

	// Function returning a function
	f := peocess()
	f(6)
}
