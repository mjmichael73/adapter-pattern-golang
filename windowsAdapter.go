package main

import "fmt"

type WindowsAdapter struct {
	windowsMachine *Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts lightning signal to USB.")
	w.windowsMachine.insertIntoUSBPort()
}