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
