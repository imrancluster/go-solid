package main

import "fmt"

// Base interface
type PaymentProcessor interface {
	ProcessPayment(amount float64) string
}

// CashPayment implements the base interface
type CashPayment struct{}

func (c CashPayment) ProcessPayment(amount float64) string {
	return fmt.Sprintf("Processing cash payment of %f", amount)
}

// CardPayment also implements the same interface
type CardPayment struct{}

func (c CardPayment) ProcessPayment(amount float64) string {
	return fmt.Sprintf("Processing card payment of %f", amount)
}

func main() {
	var paymentProcessor PaymentProcessor

	// Using CashPayment
	paymentProcessor = CashPayment{}
	fmt.Println(paymentProcessor.ProcessPayment(500))

	// Using CardPayment
	paymentProcessor = CardPayment{}
	fmt.Println(paymentProcessor.ProcessPayment(1000))
}
