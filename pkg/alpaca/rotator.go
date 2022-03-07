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

/*
	CanReverse()

	@returns true if the Rotator supports the Reverse method.
	@see https://ascom-standards.org/api/#/Rotator%20Specific%20Methods/get_rotator__device_number__canreverse
*/
func (r *Rotator) CanReverse() (bool, error) {
	return r.Alpaca.GetBooleanResponse("rotator", r.DeviceNumber, "canreverse")
}

/*
	IsMoving()

	@returns true if the rotator is currently moving to a new position. False if the focuser is stationary.
	@see https://ascom-standards.org/api/#/Rotator%20Specific%20Methods/get_rotator__device_number__ismoving
*/
func (r *Rotator) IsMoving() (bool, error) {
	return r.Alpaca.GetBooleanResponse("rotator", r.DeviceNumber, "ismoving")
}

/*
	GetMechanicalPosition()

	@returns the raw mechanical position of the rotator in degrees.
	@see https://ascom-standards.org/api/#/Rotator%20Specific%20Methods/get_rotator__device_number__mechanicalposition
*/
func (r *Rotator) GetMechanicalPosition() (float64, error) {
	return r.Alpaca.GetFloat64Response("rotator", r.DeviceNumber, "mechanicalposition")
}

/*
	GetPosition()

	@returns the current instantaneous Rotator position, in degrees.
	@see https://ascom-standards.org/api/#/Rotator%20Specific%20Methods/get_rotator__device_number__position
*/
func (r *Rotator) GetPosition() (float64, error) {
	return r.Alpaca.GetFloat64Response("rotator", r.DeviceNumber, "mechanicalposition")
}

/*
	GetReverse()

	@returns the rotator’s reverse state.
	@see https://ascom-standards.org/api/#/Rotator%20Specific%20Methods/get_rotator__device_number__reverse
*/
func (r *Rotator) GetReverse() (bool, error) {
	return r.Alpaca.GetBooleanResponse("rotator", r.DeviceNumber, "reverse")
}

/*
	SetReverse

	@returns an error or nil, if nil it sets the rotator’s reverse state.
	@see https://ascom-standards.org/api/#/Rotator%20Specific%20Methods/put_rotator__device_number__reverse
*/
func (r *Rotator) SetReverse(reverse bool) error {
	r.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		"Reverse":             fmt.Sprintf("%t", reverse),
		"ClientID":            fmt.Sprintf("%d", r.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", r.Alpaca.TransactionId),
	}

	return r.Alpaca.Put("rotator", r.DeviceNumber, "reverse", form)
}
