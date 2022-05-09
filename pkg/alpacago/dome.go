package alpacago

import "fmt"

type Dome struct {
	Alpaca       *ASCOMAlpacaAPIClient
	DeviceNumber uint
}

type ShutterStatus int32

const (
	Open ShutterStatus = iota
	Closed
	Opening
	Closing
	Error
)

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

/*
	CanSetAzimuth()

	@returns true if driver is capable of setting the dome azimuth.
	@see https://ascom-standards.org/api/#/Dome%20Specific%20Methods/get_dome__device_number__cansetazimuth
*/
func (d *Dome) CanSetAzimuth() (bool, error) {
	return d.Alpaca.GetBooleanResponse("dome", d.DeviceNumber, "cansetazimuth")
}

/*
	CanSetPark()

	@returns true if driver is capable of setting the dome park position.
	@see https://ascom-standards.org/api/#/Dome%20Specific%20Methods/get_dome__device_number__cansetpark
*/
func (d *Dome) CanSetPark() (bool, error) {
	return d.Alpaca.GetBooleanResponse("dome", d.DeviceNumber, "cansetpark")
}

/*
	CanSetShutter()

	@returns true if driver is capable of setting the shutter state.
	@see https://ascom-standards.org/api/#/Dome%20Specific%20Methods/get_dome__device_number__cansetshutter
*/
func (d *Dome) CanSetShutter() (bool, error) {
	return d.Alpaca.GetBooleanResponse("dome", d.DeviceNumber, "cansetshutter")
}

/*
	CanSlave()

	@returns true if driver is capable of slaving to a telescope.
	@see https://ascom-standards.org/api/#/Dome%20Specific%20Methods/get_dome__device_number__canslave
*/
func (d *Dome) CanSlave() (bool, error) {
	return d.Alpaca.GetBooleanResponse("dome", d.DeviceNumber, "canslave")
}

/*
	CanSyncAzimuth()

	@returns true if driver is capable of synchronizing the dome azimuth position using the SyncToAzimuth(Double) method.
	@see https://ascom-standards.org/api/#/Dome%20Specific%20Methods/get_dome__device_number__cansyncazimuth
*/
func (d *Dome) CanSyncAzimuth() (bool, error) {
	return d.Alpaca.GetBooleanResponse("dome", d.DeviceNumber, "cansyncazimuth")
}

/*
	GetShutterStatus()

	@returns the status of the dome shutter or roll-off roof. 0 = Open, 1 = Closed, 2 = Opening, 3 = Closing, 4 = Shutter status error
	@see https://ascom-standards.org/api/#/Dome%20Specific%20Methods/get_dome__device_number__shutterstatus
*/
func (d *Dome) GetShutterStatus() (ShutterStatus, error) {
	status, err := d.Alpaca.GetInt32Response("dome", d.DeviceNumber, "shutterstatus")
	return ShutterStatus(status), err
}

/*
	IsSlaved()

	@returns true if the dome is slaved to the telescope in its hardware, else false.
	@see https://ascom-standards.org/api/#/Dome%20Specific%20Methods/get_dome__device_number__slaved
*/
func (d *Dome) IsSlaved() (bool, error) {
	return d.Alpaca.GetBooleanResponse("dome", d.DeviceNumber, "slaved")
}

/*
	SetSlaved()

	@returns error, or nil, if nil it sets the salved state of the dome.
	@see https://ascom-standards.org/api/#/Dome%20Specific%20Methods/put_dome__device_number__slaved
*/
func (d *Dome) SetSlaved(slaved bool) error {
	d.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		// Set True if telescope is slaved to dome, otherwise False
		"Slaved":              fmt.Sprintf("%t", slaved),
		"ClientID":            fmt.Sprintf("%d", d.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", d.Alpaca.TransactionId),
	}

	return d.Alpaca.Put("dome", d.DeviceNumber, "slaved", form)
}

/*
	IsSlewing()

	@return true if any part of the dome is currently moving, False if all dome components are steady.
	@see https://ascom-standards.org/api/#/Dome%20Specific%20Methods/get_dome__device_number__slewing
*/
func (d *Dome) IsSlewing() (bool, error) {
	return d.Alpaca.GetBooleanResponse("dome", d.DeviceNumber, "slewing")
}

/*
	AbortSlew()

	@returs error, or nil, if nil it aborts the current slew operation.
	@see https://ascom-standards.org/api/#/Dome%20Specific%20Methods/put_dome__device_number__abortslew
*/
func (d *Dome) AbortSlew() error {
	d.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		"ClientID":            fmt.Sprintf("%d", d.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", d.Alpaca.TransactionId),
	}

	return d.Alpaca.Put("dome", d.DeviceNumber, "abortslew", form)
}

/*
	CloseShutter()

	@returns error, or nil, if nil it closes the shutter or otherwise shield telescope from the sky.
	@see https://ascom-standards.org/api/#/Dome%20Specific%20Methods/put_dome__device_number__closeshutter
*/
func (d *Dome) CloseShutter() error {
	d.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		"ClientID":            fmt.Sprintf("%d", d.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", d.Alpaca.TransactionId),
	}

	return d.Alpaca.Put("dome", d.DeviceNumber, "closeshutter", form)
}

/*
	FindHome()

	@returns error, or nil, if nil it starts operation to search for the dome home position.
	@effects after home position is established initializes azimuth to the default value and sets the AtHome flag.
	@see https://ascom-standards.org/api/#/Dome%20Specific%20Methods/put_dome__device_number__findhome
*/
func (d *Dome) FindHome() error {
	d.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		"ClientID":            fmt.Sprintf("%d", d.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", d.Alpaca.TransactionId),
	}

	return d.Alpaca.Put("dome", d.DeviceNumber, "findhome", form)
}

/*
	OpenShutter()

	@returns error, or nil, if nil it opens the shutter or otherwise expose telescope to the sky.
	@see https://ascom-standards.org/api/#/Dome%20Specific%20Methods/put_dome__device_number__openshutter
*/
func (d *Dome) OpenShutter() error {
	d.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		"ClientID":            fmt.Sprintf("%d", d.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", d.Alpaca.TransactionId),
	}

	return d.Alpaca.Put("dome", d.DeviceNumber, "openshutter", form)
}
