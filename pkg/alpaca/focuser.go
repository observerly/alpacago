package alpaca

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
func (f *Focuser) GetStepSize() (int32, error) {
	return f.Alpaca.GetInt32Response("focuser", f.DeviceNumber, "stepsize")
}

/*
	GetTemperatureCompensation()

	@returns the state of temperature compensation mode (if available), else always false.
	@see https://ascom-standards.org/api/#/Focuser%20Specific%20Methods/get_focuser__device_number__tempcomp
*/
func (f *Focuser) GetTemperatureCompensation() (bool, error) {
	return f.Alpaca.GetBooleanResponse("focuser", f.DeviceNumber, "tempcomp")
}
