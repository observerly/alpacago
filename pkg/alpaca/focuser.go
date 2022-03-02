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
