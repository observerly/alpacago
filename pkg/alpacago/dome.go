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

/*
	IsAtHome()

	@returns returns true if the dome is in the home position. This is normally used following a FindHome() operation.
	The value is reset with any azimuth slew operation that moves the dome away from the home position. AtHome may
	also become true durng normal slew operations, if the dome passes through the home position and the dome controller
	hardware is capable of detecting that; or at the end of a slew operation if the dome comes to rest at the home position.
	@see https://ascom-standards.org/api/#/Dome%20Specific%20Methods/get_dome__device_number__athome
*/
func (d *Dome) IsAtHome() (bool, error) {
	return d.Alpaca.GetBooleanResponse("dome", d.DeviceNumber, "athome")
}

/*
	IsAtPark()

	@returns true if the dome is in the programmed park position. Set only following a Park() operation and reset with any slew operation.
	@see https://ascom-standards.org/api/#/Dome%20Specific%20Methods/get_dome__device_number__atpark
*/
func (d *Dome) IsAtPark() (bool, error) {
	return d.Alpaca.GetBooleanResponse("dome", d.DeviceNumber, "atpark")
}

/*
	GetAzimuth()

	@returns the dome azimuth (degrees, North-referenced, positive East/clockwise, i.e., 90 East, 180 South, 270 West).
	@see https://ascom-standards.org/api/#/Dome%20Specific%20Methods/get_dome__device_number__azimuth
*/
func (d *Dome) GetAzimuth() (float64, error) {
	return d.Alpaca.GetFloat64Response("dome", d.DeviceNumber, "azimuth")
}

/*
	CanFindHome()

	@returns true if the dome can move to the home position.
	@see https://ascom-standards.org/api/#/Dome%20Specific%20Methods/get_dome__device_number__canfindhome
*/
func (d *Dome) CanFindHome() (bool, error) {
	return d.Alpaca.GetBooleanResponse("dome", d.DeviceNumber, "canfindhome")
}

/*
	CanPark()

	@returns true  if the dome is capable of programmed parking (Park() method)
	@see https://ascom-standards.org/api/#/Dome%20Specific%20Methods/get_dome__device_number__canpark
*/
func (d *Dome) CanPark() (bool, error) {
	return d.Alpaca.GetBooleanResponse("dome", d.DeviceNumber, "canpark")
}

/*
	CanSetAltitude()

	@returns true if driver is capable of setting the dome altitude.
	@see https://ascom-standards.org/api/#/Dome%20Specific%20Methods/get_dome__device_number__cansetaltitude
*/
func (d *Dome) CanSetAltitude() (bool, error) {
	return d.Alpaca.GetBooleanResponse("dome", d.DeviceNumber, "cansetaltitude")
}
