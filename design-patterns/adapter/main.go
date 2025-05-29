package main

func main() {
	client := &Client{}
	mac := &Mac{}

	client.InsertLightningConnectorIntoComputer(mac)
	windowsMachine := &Windows{}
	windowsAdapter := &WindowsAdapter{
		windowMachine: windowsMachine,
	}
	client.InsertLightningConnectorIntoComputer(windowsAdapter)
}
