package alpacago

import "fmt"

type Focuser struct {
	Alpaca       *ASCOMAlpacaAPIClient
	DeviceNumber uint
}

func NewFocuser(clientId uint32, secure bool, domain string, ip string, port int32, deviceNumber uint) *Focuser {
	alpaca := NewAlpacaAPI(clientId, secure, domain, ip, port)

	focuser := Focuser{
		Alpaca:       alpaca,
		DeviceNumber: deviceNumber,
	}

	return &focuser
}

/*
	IsConnected() common method to all ASCOM Alpaca compliant devices

	@returns the connected state of the device
	@see https://ascom-standards.org/api/#/ASCOM%20Methods%20Common%20To%20All%20Devices/get__device_type___device_number__connected
*/
func (f *Focuser) IsConnected() (bool, error) {
	return f.Alpaca.GetBooleanResponse("focuser", f.DeviceNumber, "connected")
}

/*
	SetConnected() common method to all ASCOM Alpaca compliant devices

	@param connected bool (set True to connect to the device hardware, set false to disconnect from the device hardware)
	@returns the connected state of the device
	@see https://ascom-standards.org/api/#/ASCOM%20Methods%20Common%20To%20All%20Devices/put__device_type___device_number__connected
*/
func (f *Focuser) SetConnected(connected bool) error {
	f.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		// Set True to connect to the device hardware, set False to disconnect from the device hardware
		"Connected":           fmt.Sprintf("%t", connected),
		"ClientID":            fmt.Sprintf("%d", f.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", f.Alpaca.TransactionId),
	}

	return f.Alpaca.Put("focuser", f.DeviceNumber, "connected", form)
}

/*
	IsAbsolute()

	@returns true if the focuser is capable of absolute position; that is, being commanded to a specific step location.
	@see https://ascom-standards.org/api/#/Focuser%20Specific%20Methods/get_focuser__device_number__absolute
*/
func (f *Focuser) IsAbsolute() (bool, error) {
	return f.Alpaca.GetBooleanResponse("focuser", f.DeviceNumber, "absolute")
}

/*
	IsMoving()

	@returns true if the focuser is currently moving to a new position and false if the focuser is stationary.
	@see https://ascom-standards.org/api/#/Focuser%20Specific%20Methods/get_focuser__device_number__ismoving
*/
func (f *Focuser) IsMoving() (bool, error) {
	return f.Alpaca.GetBooleanResponse("focuser", f.DeviceNumber, "ismoving")
}

/*
	GetMaxIncrement()

	@returns the maximum increment size allowed by the focuser; i.e. the maximum number of steps allowed in one move operation.
	@see https://ascom-standards.org/api/#/Focuser%20Specific%20Methods/get_focuser__device_number__maxincrement
*/
func (f *Focuser) GetMaxIncrement() (int32, error) {
	return f.Alpaca.GetInt32Response("focuser", f.DeviceNumber, "maxincrement")
}

/*
	GetMaxStep()

	@returns the maximum step size allowed by the focuser; i.e. the maximum step position permitted.
	@see https://ascom-standards.org/api/#/Focuser%20Specific%20Methods/get_focuser__device_number__maxstep
*/
func (f *Focuser) GetMaxStep() (int32, error) {
	return f.Alpaca.GetInt32Response("focuser", f.DeviceNumber, "maxstep")
}

/*
	GetPosition()

	@returns the focuser's current position.
	@see https://ascom-standards.org/api/#/Focuser%20Specific%20Methods/get_focuser__device_number__position
*/
func (f *Focuser) GetPosition() (int32, error) {
	return f.Alpaca.GetInt32Response("focuser", f.DeviceNumber, "position")
}

/*
	GetStepSize()

	@returns the step size (microns) for the focuser.
	@see https://ascom-standards.org/api/#/Focuser%20Specific%20Methods/get_focuser__device_number__stepsize
*/
func (f *Focuser) GetStepSize() (float64, error) {
	return f.Alpaca.GetFloat64Response("focuser", f.DeviceNumber, "stepsize")
}

/*
	GetTemperatureCompensation()

	@returns the state of temperature compensation mode (if available), else always false.
	@see https://ascom-standards.org/api/#/Focuser%20Specific%20Methods/get_focuser__device_number__tempcomp
*/
func (f *Focuser) GetTemperatureCompensation() (bool, error) {
	return f.Alpaca.GetBooleanResponse("focuser", f.DeviceNumber, "tempcomp")
}

/*
	SetTemperatureCompensation()

	@params tempComp bool (set true to enable the focuser's temperature compensation mode, otherwise false for normal operation.)
	@returns an error or nil, if nil it sets the state of temperature compensation mode.
	@see https://ascom-standards.org/api/#/Focuser%20Specific%20Methods/put_focuser__device_number__tempcomp
*/
func (f *Focuser) SetTemperatureCompensation(tempComp bool) error {
	f.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		// Set true to enable the focuser's temperature compensation mode, otherwise false for normal operation.
		"TempComp":            fmt.Sprintf("%t", tempComp),
		"ClientID":            fmt.Sprintf("%d", f.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", f.Alpaca.TransactionId),
	}

	return f.Alpaca.Put("focuser", f.DeviceNumber, "tempcomp", form)
}

/*
	IsTemperatureCompensationAvailable()

	@returns true if focuser has temperature compensation available., else always false.
	@see https://ascom-standards.org/api/#/Focuser%20Specific%20Methods/get_focuser__device_number__tempcompavailable
*/
func (f *Focuser) IsTemperatureCompensationAvailable() (bool, error) {
	return f.Alpaca.GetBooleanResponse("focuser", f.DeviceNumber, "tempcompavailable")
}

/*
	GetTemperature()

	@returns the current ambient temperature as measured by the focuser.
	@see https://ascom-standards.org/api/#/Focuser%20Specific%20Methods/get_focuser__device_number__temperature
*/
func (f *Focuser) GetTemperature() (float64, error) {
	return f.Alpaca.GetFloat64Response("focuser", f.DeviceNumber, "temperature")
}

/*
	SetHalt

	@returns an error or nil, if nil it immediately stop any focuser motion due to a previous move() method call.
	@see https://ascom-standards.org/api/#/Focuser%20Specific%20Methods/put_focuser__device_number__halt
*/
func (f *Focuser) SetHalt() error {
	f.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		"ClientID":            fmt.Sprintf("%d", f.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", f.Alpaca.TransactionId),
	}

	return f.Alpaca.Put("focuser", f.DeviceNumber, "halt", form)
}

/*
	SetMove

	@params position int32 (step distance or absolute position, depending on the value of the absolute property.)
	@returns an error or nil, if nil it moves the focuser by the specified amount or to the specified
	position depending on the value of the Absolute property.
	@see https://ascom-standards.org/api/#/Focuser%20Specific%20Methods/put_focuser__device_number__move
*/
func (f *Focuser) SetMove(position int32) error {
	f.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		// Step distance or absolute position, depending on the value of the Absolute property
		"Position":            fmt.Sprintf("%d", position),
		"ClientID":            fmt.Sprintf("%d", f.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", f.Alpaca.TransactionId),
	}

	return f.Alpaca.Put("focuser", f.DeviceNumber, "move", form)
}
