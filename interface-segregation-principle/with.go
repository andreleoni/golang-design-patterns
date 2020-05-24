package main

import "fmt"

type Document struct {
	Content string
}

type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

type PrinterMachine struct{}

func (m PrinterMachine) Print(d Document) {
	fmt.Println("Printing... ", d.Content)
}

type ScannerMachine struct{}

func (m ScannerMachine) Scan(d Document) {
	fmt.Println("Scanning... ", d.Content)
}

// Interfaces aggregation if it makes sense
type MultifunctionDevice interface {
	Printer
	Scanner
}

type MultifunctionalMachine struct{}

func (m MultifunctionalMachine) Print(d Document) {
	fmt.Println("Printing... ", d.Content)
}

func (m MultifunctionalMachine) Scan(d Document) {
	fmt.Println("Scanning... ", d.Content)
}

func main() {
	var p Printer
	var s Scanner
	var mfd MultifunctionDevice

	d := Document{"My document content"}

	p = PrinterMachine{}
	p.Print(d)

	s = ScannerMachine{}
	s.Scan(d)

	mfd = MultifunctionalMachine{}
	mfd.Print(d)
	mfd.Scan(d)
}
