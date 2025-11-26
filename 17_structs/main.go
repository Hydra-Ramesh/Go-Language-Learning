package main

import (
	"fmt"
	"time"
)

// Defining a struct to represent a customer
type Customer struct {
	name  string
	phone string
}

// Defining a struct to represent an order
type Order struct {
	id        string
	amount    float64
	status    string
	createdAt time.Time // nanosecond precision
	// Nested struct
	customer Customer
}

// Function to create a new order and return its pointer
func newOrder(id string, amount float64, status string) *Order {
	myOrders := Order{
		id:     id,
		amount: amount,
		status: status,
	}
	return &myOrders
}

// Receiver method to change the status of the order
func (o *Order) changeStatus(status string) {
	o.status = status
}
func main() {
	// If you did not set any value, Go will assign the zero value for that type and for string it is ""
	order := Order{
		id:     "12345",
		amount: 250.75,
		status: "Processing",
	}
	order.createdAt = time.Now()
	fmt.Println(order)
	fmt.Println("Order created at:", order.createdAt)

	myOrder := Order{
		id:        "67890",
		amount:    99.99,
		status:    "Shipped",
		createdAt: time.Date(2024, time.March, 15, 10, 30, 0, 0, time.UTC),
	}
	myOrder.changeStatus("Delivered")
	fmt.Println(myOrder)
	fmt.Println("My order created at:", myOrder.createdAt)
	fmt.Println("My order status:", myOrder.status)

	newOrd := newOrder("54321", 150.00, "Pending")
	fmt.Println(newOrd)

	language := struct {
		name   string
		isGood bool
	}{name: "Go", isGood: true}
	fmt.Println(language)

	myOrder2 := Order{
		id:     "11223",
		amount: 300.00,
		status: "In Transit",
		customer: Customer{
			name:  "Alice Smith",
			phone: "555-1234",
		},
	}
	fmt.Println(myOrder2)
}
