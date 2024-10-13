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
