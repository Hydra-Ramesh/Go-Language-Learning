package main

import "fmt"

func main() {
	age := 20
	if age < 18 {
		fmt.Println("Minor")
	} else {
		fmt.Println("Adult")
	}

	newAge := 16
	if newAge < 18 {
		fmt.Println("Minor")
	} else if newAge == 18 {
		fmt.Println("Just became an Adult")
	} else {
		fmt.Println("Adult")
	}

	var role = "admin"
	var hasAccess = true
	if role == "admin" && hasAccess {
		fmt.Println("Access granted to admin")
	}

	// We can also initialize a variable in the if statement
	if num := 10; num%2 == 0 {
		fmt.Println("Even number")
	} else {
		fmt.Println("Odd number")
	}

	// Note: Go does not have a ternary operator like some other languages.
}
