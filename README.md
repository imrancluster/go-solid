Let's dive into SOLID principles with practical GoLang examples. The SOLID principles are:

1. **Single Responsibility Principle (SRP)**
2. **Open/Closed Principle (OCP)**
3. **Liskov Substitution Principle (LSP)**
4. **Interface Segregation Principle (ISP)**
5. **Dependency Inversion Principle (DIP)**

We'll go through each with real-life scenarios and Go code examples.

### 1. Single Responsibility Principle (SRP)

**Definition**: A class (or function/module) should have only one reason to change, meaning it should only have one responsibility.

#### Go Example:

Imagine a billing system. A struct `Invoice` should handle invoice data only, while another struct handles printing.

```go
package main

import (
	"fmt"
)

type Invoice struct {
	ID     int
	Amount float64
}

func (i Invoice) CalculateTax() float64 {
	return i.Amount * 0.15 // 15% tax calculation
}

// Separate responsibility for printing the invoice
type InvoicePrinter struct{}

func (p InvoicePrinter) PrintInvoice(invoice Invoice) {
	fmt.Printf("Invoice ID: %d, Amount: %f\n", invoice.ID, invoice.Amount)
}

func main() {
	invoice := Invoice{ID: 1, Amount: 1000}
	printer := InvoicePrinter{}
	printer.PrintInvoice(invoice)
}
```

Here, the `Invoice` struct handles the tax calculation, while `InvoicePrinter` handles printing, following SRP.

### 2. Open/Closed Principle (OCP)

**Definition**: Software entities should be open for extension but closed for modification.

#### Go Example:

Let's extend an invoice system where we want different discount strategies without modifying the original `Invoice` structure.

```go
package main

import "fmt"

type Invoice struct {
	Amount float64
}

// Base discount interface
type Discount interface {
	ApplyDiscount(amount float64) float64
}

// Specific discount implementation for holiday offers
type HolidayDiscount struct{}

func (h HolidayDiscount) ApplyDiscount(amount float64) float64 {
	return amount * 0.9 // 10% off
}

// New discount type for loyalty members
type LoyaltyDiscount struct{}

func (l LoyaltyDiscount) ApplyDiscount(amount float64) float64 {
	return amount * 0.85 // 15% off
}

func main() {
	invoice := Invoice{Amount: 1000}

	// Apply a holiday discount
	holidayDiscount := HolidayDiscount{}
	fmt.Println("Holiday Discount:", holidayDiscount.ApplyDiscount(invoice.Amount))

	// Apply a loyalty discount
	loyaltyDiscount := LoyaltyDiscount{}
	fmt.Println("Loyalty Discount:", loyaltyDiscount.ApplyDiscount(invoice.Amount))
}
```

By introducing new types that implement `Discount`, we can extend the behavior without changing the original code.

### 3. Liskov Substitution Principle (LSP)

**Definition**: Objects of a superclass should be replaceable with objects of a subclass without affecting the correctness of the program.

#### Go Example:

Assume we have a system for payments with both cash and card payments.

```go
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
```

Both `CashPayment` and `CardPayment` can be used interchangeably, adhering to LSP.

### 4. Interface Segregation Principle (ISP)

**Definition**: Clients should not be forced to implement interfaces they don't use. Instead, create smaller, specific interfaces.

#### Go Example:

Imagine different types of printers that only need certain functionalities.

```go
package main

import "fmt"

// Split the functionalities into smaller interfaces
type Printer interface {
	Print()
}

type Scanner interface {
	Scan()
}

// A normal printer implements only the Printer interface
type SimplePrinter struct{}

func (p SimplePrinter) Print() {
	fmt.Println("Printing document")
}

// A multifunction printer implements both Printer and Scanner
type MultifunctionPrinter struct{}

func (m MultifunctionPrinter) Print() {
	fmt.Println("Printing document")
}

func (m MultifunctionPrinter) Scan() {
	fmt.Println("Scanning document")
}

func main() {
	printer := SimplePrinter{}
	printer.Print()

	mfp := MultifunctionPrinter{}
	mfp.Print()
	mfp.Scan()
}
```

Here, `SimplePrinter` implements only the functionality it needs, and `MultifunctionPrinter` can implement both without forcing either one to implement unwanted methods.

### 5. Dependency Inversion Principle (DIP)

**Definition**: High-level modules should not depend on low-level modules. Both should depend on abstractions.

#### Go Example:

Let's decouple payment methods from the payment processing system by using interfaces.

```go
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
```

The `PaymentProcessor` depends on the abstraction (`PaymentMethod`) instead of concrete implementations. This allows us to easily switch or extend payment methods.

---

These principles help keep code clean, maintainable, and extensible, particularly in larger systems. Would you like to dive deeper into any specific principle or concept?
