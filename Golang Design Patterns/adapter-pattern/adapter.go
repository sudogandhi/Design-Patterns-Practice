package main

import "fmt"

// There is a client
type Client struct {
}

// The client wants to invoke insert a lightning port (func) on a computer (struct)
func (c *Client) InsertLightningConnectorToAComputer(computer Computer) {
	computer.InsertLightningConnector()
}

// Computer is an iterface because there can be multiple computers of different types.
// Both mac and windows can implement this interface
type Computer interface {
	InsertLightningConnector()
}

type Mac struct {
}

func (mac *Mac) InsertLightningConnector() {
	fmt.Println("Lightning Connector is connected to MAC")
}

type Windows struct {
}

// But lightning connector cannot be inserted in windows
// Windows accepts only USB port.
// So we need to implement an adapter which will be a wrapper around windows
// and can connect lightning.
func (windows *Windows) InsertUSBPort() {
	fmt.Println("USB port is connected to windows")
}

type WindowsAdapter struct {
	windows *Windows
}

func (adapter *WindowsAdapter) InsertLightningConnector() {
	fmt.Println("Adapter converts lightning to USB signals")
	adapter.windows.InsertUSBPort()
}

func main() {
	// Create a client
	client := &Client{}
	// Create a mac computer
	mac := &Mac{}
	// Create a windows computer
	windows := &Windows{}
	// Create a windows adapter
	adapter := &WindowsAdapter{windows: windows}

	client.InsertLightningConnectorToAComputer(mac)

	client.InsertLightningConnectorToAComputer(adapter)
}
