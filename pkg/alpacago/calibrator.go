package alpacago

import "fmt"

type CalibratorState int32

type CoverState int32

const (
	// This device does not have a calibration capability
	CalibratorNotPresent CalibratorState = iota
	// The calibrator is off
	CalibratorOff
	// The calibrator is stabilising or is not yet in the commanded state
	CalibratorNotReady
	// The calibrator is ready for use
	CalibratorReady
	// The calibrator state is unknown
	CalibratorUnknown
	// The calibrator encountered an error when changing state
	CalibratorError
)

const (
	// This device does not have a cover that can be closed independently
	CoverNotPresent CoverState = iota
	// The cover is closed
	CoverClosed
	// The cover is moving to a new position
	CoverMoving
	// The Cover is open
	CoverOpen
	// The Cover state is unknown
	CoverUnknown
	// The Cover encountered an error when changing state
	CoverError
)

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

/*
	GetStatus()

	@returns the state of the calibration device, if present, otherwise returns "NotPresent". The calibrator state mode is specified as an integer value from the CalibratorStatus Enum.
	@see https://ascom-standards.org/api/#/ASCOM%20Methods%20Common%20To%20All%20Devices/get__device_type___device_number__calibratorstate
	@see https://ascom-standards.org/Help/Platform/html/T_ASCOM_DeviceInterface_CalibratorStatus.htm
*/
func (c *CoverCalibrator) GetStatus() (CalibratorState, error) {
	status, err := c.Alpaca.GetInt32Response("covercalibrator", c.DeviceNumber, "calibratorstate")
	return CalibratorState(status), err
}

/*
	GetCoverStatus()

	@returns the state of the device cover, if present, otherwise returns "NotPresent". The cover state mode is specified as an integer value from the CoverStatus Enum.
	@see https://ascom-standards.org/api/#/CoverCalibrator%20Specific%20Methods/get_covercalibrator__device_number__coverstate
	@see https://ascom-standards.org/Help/Platform/html/T_ASCOM_DeviceInterface_CoverStatus.htm
*/
func (c *CoverCalibrator) GetCoverStatus() (CoverState, error) {
	status, err := c.Alpaca.GetInt32Response("covercalibrator", c.DeviceNumber, "coverstate")
	return CoverState(status), err
}

/*
	GetMaxBrightness()

	@returns the brightness value that makes the calibrator deliver its maximum illumination.
	@see https://ascom-standards.org/api/#/CoverCalibrator%20Specific%20Methods/get_covercalibrator__device_number__maxbrightness
*/
func (c *CoverCalibrator) GetMaxBrightness() (int32, error) {
	return c.Alpaca.GetInt32Response("covercalibrator", c.DeviceNumber, "maxbrightness")
}

/*
	SetCalibratorOn()

	@returns and error, or nil, if nil the calibrator on if the device has calibration capability.
	@see https://ascom-standards.org/api/#/CoverCalibrator%20Specific%20Methods/put_covercalibrator__device_number__calibratoron
*/
func (c *CoverCalibrator) SetCalibratorOn(brightness int32) error {
	c.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		// The brightness value that makes the calibrator deliver its maximum illumination.
		"Brightness":          fmt.Sprintf("%d", brightness),
		"ClientID":            fmt.Sprintf("%d", c.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", c.Alpaca.TransactionId),
	}

	return c.Alpaca.Put("covercalibrator", c.DeviceNumber, "calibratoron", form)
}

/*
	SetCalibratorOff()

	@turns the calibrator off if the device has calibration capability.
	@see https://ascom-standards.org/api/#/CoverCalibrator%20Specific%20Methods/put_covercalibrator__device_number__calibratoroff
*/
func (c *CoverCalibrator) SetCalibratorOff() error {
	c.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		"ClientID":            fmt.Sprintf("%d", c.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", c.Alpaca.TransactionId),
	}

	return c.Alpaca.Put("covercalibrator", c.DeviceNumber, "calibratoroff", form)
}

/*
	CloseCover()

	@returns and error, or nil, if nil initiates cover closing if a cover is present.
	@see https://ascom-standards.org/api/#/CoverCalibrator%20Specific%20Methods/put_covercalibrator__device_number__closecover
*/
func (c *CoverCalibrator) CloseCover() error {
	c.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		"ClientID":            fmt.Sprintf("%d", c.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", c.Alpaca.TransactionId),
	}

	return c.Alpaca.Put("covercalibrator", c.DeviceNumber, "closecover", form)
}

/*
	HaltCover()

	@returns and error, or nil, if nil stops any cover movement that may be in progress if a cover is present and cover movement can be interrupted.
	@see https://ascom-standards.org/api/#/CoverCalibrator%20Specific%20Methods/put_covercalibrator__device_number__haltcover
*/
func (c *CoverCalibrator) HaltCover() error {
	c.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		"ClientID":            fmt.Sprintf("%d", c.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", c.Alpaca.TransactionId),
	}

	return c.Alpaca.Put("covercalibrator", c.DeviceNumber, "haltcover", form)
}
