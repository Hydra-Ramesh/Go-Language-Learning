package main

import (
	"fmt"
	"time"
)

func main() {
	//Simple switch statement
	// We should not break after each case in Go like other languages
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}

	// Multiple values in a case
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday.")
	}

	// Type switch
	whoAmI := func(i interface{}) {
		switch v := i.(type) {
		case int:
			fmt.Println("I am an int")
		case string:
			fmt.Println("I am a string")
		default:
			fmt.Println("I am something else", v)
		}
	}
	whoAmI(42)
	whoAmI("hello")
	whoAmI(3.14)
}
