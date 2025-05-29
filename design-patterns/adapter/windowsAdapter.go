package main

import "fmt"

type WindowsAdapter struct {
	windowMachine *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("WindowsAdapter: Lightning connector inserted into Windows machine.")
	w.windowMachine.insertIntoUSBPort()
}
