package main

import "fmt"

type Document struct {
	Content string
}

type MultifunctionalPrinter interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type MultifunctionalPrinterMachine struct{}

func (m MultifunctionalPrinterMachine) Print(d Document) {
	fmt.Println("Printing... ", d.Content)
}

func (m MultifunctionalPrinterMachine) Fax(d Document) {
	fmt.Println("Faxing... ", d.Content)
}

func (m MultifunctionalPrinterMachine) Scan(d Document) {
	fmt.Println("Scanning... ", d.Content)
}

type OldPrinterMachine struct{}

func (m OldPrinterMachine) Print(d Document) {
	fmt.Println("Printing... ", d.Content)
}
func (m OldPrinterMachine) Fax(d Document) {
	panic("Operation not supported")
}
func (m OldPrinterMachine) Scan(d Document) {
	panic("Operation not supported")
}

func main() {
	d := Document{"My document content"}

	mpm := MultifunctionalPrinterMachine{}
	mpm.Print(d)
	mpm.Scan(d)
	mpm.Fax(d)

	opm := OldPrinterMachine{}
	opm.Print(d)
}
