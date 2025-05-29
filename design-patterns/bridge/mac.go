package main

import "fmt"

type Mac struct {
	printer Printer
}

func (m *Mac) Print() {
	fmt.Println("Mac: Printing document.")
	m.printer.PrintFile()

}

func (m *Mac) SetupPrinter(p Printer) {
	fmt.Println("Mac: Setting up printer.")
	m.printer = p
}
