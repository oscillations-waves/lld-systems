package main

import "fmt"

type Windows struct {
	printer Printer
}

func (w *Windows) SetupPrinter(p Printer) {
	fmt.Println("Windows: Setting up printer.")
	w.printer = p
}
func (w *Windows) Print() {
	fmt.Println("Windows: Printing document.")
	w.printer.PrintFile()
}
