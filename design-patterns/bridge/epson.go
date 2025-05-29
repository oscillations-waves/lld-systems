package main

import "fmt"

type Epson struct{}

func (e *Epson) PrintFile() {
	fmt.Println("Epson: Printing by an Epson printer.")
}
