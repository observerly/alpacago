package alpacago

import "fmt"

type Dome struct {
	Alpaca       *ASCOMAlpacaAPIClient
	DeviceNumber uint
}

func NewDome(clientId uint32, secure bool, domain string, ip string, port int32, deviceNumber uint) *Dome {
	alpaca := NewAlpacaAPI(clientId, secure, domain, ip, port)

	dome := Dome{
		Alpaca:       alpaca,
		DeviceNumber: deviceNumber,
	}

	return &dome
}

/*
	SetConnected() common method to all ASCOM Alpaca compliant devices

	@param connected bool (set True to connect to the device hardware, set false to disconnect from the device hardware)
	@returns the connected state of the device
	@see https://ascom-standards.org/api/#/ASCOM%20Methods%20Common%20To%20All%20Devices/put__device_type___device_number__connected
*/
func (d *Dome) SetConnected(connected bool) error {
	d.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		// Set True to connect to the device hardware, set False to disconnect from the device hardware
		"Connected":           fmt.Sprintf("%t", connected),
		"ClientID":            fmt.Sprintf("%d", d.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", d.Alpaca.TransactionId),
	}

	return d.Alpaca.Put("dome", d.DeviceNumber, "connected", form)
}
