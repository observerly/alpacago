package alpaca

import "fmt"

type Rotator struct {
	Alpaca       *ASCOMAlpacaAPIClient
	DeviceNumber uint
}

func NewRotator(clientId uint32, secure bool, domain string, ip string, port int32, deviceNumber uint) *Rotator {
	alpaca := NewAlpacaAPI(clientId, secure, domain, ip, port)

	rotator := Rotator{
		Alpaca:       alpaca,
		DeviceNumber: deviceNumber,
	}

	return &rotator
}

/*
	SetConnected() common method to all ASCOM Alpaca compliant devices

	@param connected bool (set True to connect to the device hardware, set false to disconnect from the device hardware)
	@returns the connected state of the device
	@see https://ascom-standards.org/api/#/ASCOM%20Methods%20Common%20To%20All%20Devices/put__device_type___device_number__connected
*/
func (r *Rotator) SetConnected(connected bool) error {
	r.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		// Set True to connect to the device hardware, set False to disconnect from the device hardware
		"Connected":           fmt.Sprintf("%t", connected),
		"ClientID":            fmt.Sprintf("%d", r.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", r.Alpaca.TransactionId),
	}

	return r.Alpaca.Put("rotator", r.DeviceNumber, "connected", form)
}
