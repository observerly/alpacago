package alpacago

type CoverCalibrator struct {
	Alpaca       *ASCOMAlpacaAPIClient
	DeviceNumber uint
}

func NewCoverCalibrator(clientId uint32, secure bool, domain string, ip string, port int32, deviceNumber uint) *CoverCalibrator {
	alpaca := NewAlpacaAPI(clientId, secure, domain, ip, port)

	calibrator := CoverCalibrator{
		Alpaca:       alpaca,
		DeviceNumber: deviceNumber,
	}

	return &calibrator
}

/*
	IsConnected() common method to all ASCOM Alpaca compliant devices

	@returns the connected state of the device
	@see https://ascom-standards.org/api/#/ASCOM%20Methods%20Common%20To%20All%20Devices/get__device_type___device_number__connected
*/
func (c *CoverCalibrator) IsConnected() (bool, error) {
	return c.Alpaca.GetBooleanResponse("covercalibrator", c.DeviceNumber, "connected")
}
