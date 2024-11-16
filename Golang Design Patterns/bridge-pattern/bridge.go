package main

import "fmt"

// We have an interface which represents computer and an interface to reprsent printers
// Every struct which implements a computer needs to have a variable of printer type interface
// the print method of computer interface needs to invoke the printfile method of printer interface.
// this makes it easy for us to introduce new printers or computers

type Computer interface {
	print()
	setPrinter()
}

type Windows struct {
	printer Printer
}

func (windows *Windows) print() {
	fmt.Println("Printing file in window")
	windows.printer.printFile()
}

func (windows *Windows) setPrinter(printer Printer) {
	windows.printer = printer
}

type Mac struct {
	printer Printer
}

func (mac *Mac) print() {
	fmt.Println("Printing file in mac")
	mac.printer.printFile()
}

func (mac *Mac) setPrinter(printer Printer) {
	mac.printer = printer
}

type Printer interface {
	printFile()
}

type Epson struct{}

func (epson *Epson) printFile() {
	fmt.Println("Printing file using EPSON printer")
}

type Canon struct{}

func (canon *Canon) printFile() {
	fmt.Println("Printing file using Canon printer")
}

func main() {
	// Create a printer
	epson := Epson{}
	// Create a printer
	canon := Canon{}

	// Create a computer
	windows := Windows{}
	windows.setPrinter(&epson)
	windows.print()

	// Create a computer
	mac := Mac{}
	mac.setPrinter(&canon)
	mac.print()
}
