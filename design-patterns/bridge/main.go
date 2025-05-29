package main

import "fmt"

func main() {
	// Create an instance of the Epson printer
	epsonPrinter := &Epson{}
	HpPrinter := &Hp{}

	// Create an instance of the Mac computer
	macComputer := &Mac{}

	macComputer.SetupPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()
	// Create an instance of the Windows computer
	windowsComputer := &Windows{}
	windowsComputer.SetupPrinter(HpPrinter)
	windowsComputer.Print()
	fmt.Println()

	windowsComputer.SetupPrinter(epsonPrinter)
	windowsComputer.Print()
	fmt.Println()

}
