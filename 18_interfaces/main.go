package main

import "fmt"

// paymenter defines the behavior for processing payments
type paymenter interface {
	ProcessPayment(amount float64)
}

// Payment struct uses a paymenter to process payments
type Payment struct {
	gateway paymenter
}

// ProcessPayment processes the payment using the configured gateway
func (p Payment) ProcessPayment(amount float64) {
	if p.gateway == nil {
		fmt.Println("no payment gateway configured")
		return
	}
	p.gateway.ProcessPayment(amount)
	fmt.Println("Payment processed")
}

// Razorpay struct implements the paymenter interface
type Razorpay struct{}

// ProcessPayment processes payment using Razorpay gateway
func (r Razorpay) ProcessPayment(amount float64) {
	fmt.Printf("Processing payment of %.2f using Razorpay gateway\n", amount)
}

// stripe struct implements the paymenter interface
type stripe struct{}

// fakepayment struct implements the paymenter interface
type fakepayment struct{}

// ProcessPayment processes payment using Fake Payment gateway
func (f fakepayment) ProcessPayment(amount float64) {
	fmt.Printf("Processing payment of %.2f using Fake Payment gateway\n", amount)
}

// ProcessPayment processes payment using Stripe gateway
func (s stripe) ProcessPayment(amount float64) {
	fmt.Printf("Processing payment of %.2f using Stripe gateway\n", amount)
}

// paypal struct implements the paymenter interface
type paypal struct{}

// ProcessPayment processes payment using PayPal gateway
func (p paypal) ProcessPayment(amount float64) {
	fmt.Printf("Processing payment of %.2f using PayPal gateway\n", amount)
}
func main() {
	razorpayPayment := Payment{gateway: Razorpay{}}
	razorpayPayment.ProcessPayment(100.0)
	stripePayment := Payment{gateway: stripe{}}
	stripePayment.ProcessPayment(200.0)
	fakePayment := Payment{gateway: fakepayment{}}
	fakePayment.ProcessPayment(50.0)
	paypalPayment := Payment{gateway: paypal{}}
	paypalPayment.ProcessPayment(150.0)
}
