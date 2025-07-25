package adapter

import "l1/21/windows"

import "fmt"

type WindowsAdapter struct {
	WindowMachine *windows.Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.WindowMachine.InsertIntoUSBPort()
}
