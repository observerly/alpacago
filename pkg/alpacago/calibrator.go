package alpacago

import "fmt"

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

/*
	SetConnected() common method to all ASCOM Alpaca compliant devices

	@param connected bool (set True to connect to the device hardware, set false to disconnect from the device hardware)
	@returns the connected state of the device
	@see https://ascom-standards.org/api/#/ASCOM%20Methods%20Common%20To%20All%20Devices/put__device_type___device_number__connected
*/
func (c *CoverCalibrator) SetConnected(connected bool) error {
	c.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		// Set True to connect to the device hardware, set False to disconnect from the device hardware
		"Connected":           fmt.Sprintf("%t", connected),
		"ClientID":            fmt.Sprintf("%d", c.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", c.Alpaca.TransactionId),
	}

	return c.Alpaca.Put("covercalibrator", c.DeviceNumber, "connected", form)
}

/*
	GetBrightness()

	@returns the current calibrator brightness in the range 0 (completely off) to MaxBrightness (fully on)
	@see https://ascom-standards.org/api/#/CoverCalibrator%20Specific%20Methods/get_covercalibrator__device_number__brightness
*/
func (c *CoverCalibrator) GetBrightness() (float64, error) {
	return c.Alpaca.GetFloat64Response("covercalibrator", c.DeviceNumber, "brightness")
}
