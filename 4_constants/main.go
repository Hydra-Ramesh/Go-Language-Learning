package main

import "fmt"

const x = 42

// We cant use := to define constants and outside functions
// y:= "This is a string"

var y string = "This is a string"

// Grouped constant declaration
const (
	port = 8080
	host = "localhost"
)

func main() {
	// We define constants using the 'const' keyword and they cannot be changed later in the code.
	const pi = 3.14
	const greeting = "Hello, World!"
	fmt.Println("Value of pi:", pi)
	fmt.Println(greeting)

	fmt.Println("Value of x:", x)
	fmt.Println("Value of y:", y)
	fmt.Println("Server is running on", host, "at port", port)
}
