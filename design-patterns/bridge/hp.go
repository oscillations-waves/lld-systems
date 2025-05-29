package main

import "fmt"

type Hp struct {
}

func (h *Hp) PrintFile() {
	fmt.Println("HP: Printing by a HP printer.")
}
