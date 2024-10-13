package main

import "fmt"

// PaymentMethod interface (abstraction)
type PaymentMethod interface {
	Pay(amount float64) string
}

// CreditCard struct (low-level module)
type CreditCard struct{}

func (cc CreditCard) Pay(amount float64) string {
	return fmt.Sprintf("Paid %f using Credit Card", amount)
}

// PayPal struct (low-level module)
type PayPal struct{}

func (pp PayPal) Pay(amount float64) string {
	return fmt.Sprintf("Paid %f using PayPal", amount)
}

// PaymentProcessor struct (high-level module)
type PaymentProcessor struct {
	Method PaymentMethod
}

func (p PaymentProcessor) Process(amount float64) {
	fmt.Println(p.Method.Pay(amount))
}

func main() {
	// Process payment using Credit Card
	processor := PaymentProcessor{Method: CreditCard{}}
	processor.Process(100)

	// Process payment using PayPal
	processor = PaymentProcessor{Method: PayPal{}}
	processor.Process(200)
}
