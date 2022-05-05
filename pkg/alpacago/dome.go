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
	IsConnected() common method to all ASCOM Alpaca compliant devices

	@returns the connected state of the device
	@see https://ascom-standards.org/api/#/ASCOM%20Methods%20Common%20To%20All%20Devices/get__device_type___device_number__connected
*/
func (d *Dome) IsConnected() (bool, error) {
	return d.Alpaca.GetBooleanResponse("dome", d.DeviceNumber, "connected")
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

/*
	GetAltitude()

	@returns The dome altitude (degrees, horizon zero and increasing positive to 90 zenith).
	@see https://ascom-standards.org/api/#/Dome%20Specific%20Methods/get_dome__device_number__altitude
*/
func (d *Dome) GetAltitude() (float64, error) {
	return d.Alpaca.GetFloat64Response("dome", d.DeviceNumber, "altitude")
}
