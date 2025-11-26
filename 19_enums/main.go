package main

import "fmt"

type OrderStatus int

const (
	Pending OrderStatus = iota
	Shipped
	Delivered
	Canceled
)

func changeStatus(status OrderStatus) {
	fmt.Println("Changing order status to:", status)
}
func main() {
	changeStatus(Shipped)
}
