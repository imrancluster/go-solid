package main

import "fmt"

const (
	HOLIDAY_DISCOUNT_PERCENTAGE = 0.9
	ROYALTY_DISCOUNT_PERCENTAGE = 0.85
)

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
	return amount * HOLIDAY_DISCOUNT_PERCENTAGE // 10% off
}

// New discount type for the loyalty members
type LoyaltyDiscount struct{}

func (l LoyaltyDiscount) ApplyDiscount(amount float64) float64 {
	return amount * ROYALTY_DISCOUNT_PERCENTAGE // 15% off
}

func main() {
	invoice := Invoice{Amount: 1000}

	// Apply a holiday discount
	holidayDiscount := HolidayDiscount{}
	fmt.Println("Holiday Discount: ", holidayDiscount.ApplyDiscount(invoice.Amount))

	// Apply a loyalty discount
	loyaltyDiscount := LoyaltyDiscount{}
	fmt.Println("Holiday Discount: ", loyaltyDiscount.ApplyDiscount(invoice.Amount))

}
