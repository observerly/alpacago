package alpacago

import "fmt"

type OperationalState int32

const (
	CameraIdle OperationalState = iota
	CameraWaiting
	CameraExposing
	CameraReading
	CameraDownload
	CameraError
)

type Camera struct {
	Alpaca       *ASCOMAlpacaAPIClient
	DeviceNumber uint
}

func NewCamera(clientId uint32, secure bool, domain string, ip string, port int32, deviceNumber uint) *Camera {
	alpaca := NewAlpacaAPI(clientId, secure, domain, ip, port)

	camera := Camera{
		Alpaca:       alpaca,
		DeviceNumber: deviceNumber,
	}

	return &camera
}

/*
	IsConnected() common method to all ASCOM Alpaca compliant devices

	@returns the connected state of the device
	@see https://ascom-standards.org/api/#/ASCOM%20Methods%20Common%20To%20All%20Devices/get__device_type___device_number__connected
*/
func (c *Camera) IsConnected() (bool, error) {
	return c.Alpaca.GetBooleanResponse("camera", c.DeviceNumber, "connected")
}

/*
	SetConnected() common method to all ASCOM Alpaca compliant devices

	@param connected bool (set True to connect to the device hardware, set false to disconnect from the device hardware)
	@returns the connected state of the device
	@see https://ascom-standards.org/api/#/ASCOM%20Methods%20Common%20To%20All%20Devices/put__device_type___device_number__connected
*/
func (c *Camera) SetConnected(connected bool) error {
	c.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		// Set True to connect to the device hardware, set False to disconnect from the device hardware
		"Connected":           fmt.Sprintf("%t", connected),
		"ClientID":            fmt.Sprintf("%d", c.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", c.Alpaca.TransactionId),
	}

	return c.Alpaca.Put("camera", c.DeviceNumber, "connected", form)
}

/*
	GetBayerOffsetX()

	@returns the X offset of the Bayer matrix, as defined in SensorType.
	@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__bayeroffsetx
*/
func (c *Camera) GetBayerOffsetX() (int32, error) {
	return c.Alpaca.GetInt32Response("camera", c.DeviceNumber, "bayeroffsetx")
}

/*
	GetBayerOffsetY()

	@returns the Y offset of the Bayer matrix, as defined in SensorType.
	@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__bayeroffsety
*/
func (c *Camera) GetBayerOffsetY() (int32, error) {
	return c.Alpaca.GetInt32Response("camera", c.DeviceNumber, "bayeroffsety")
}

/*
	GetBinX()

	@returns the binning factor for the X axis.
	@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__binx
*/
func (c *Camera) GetBinX() (int32, error) {
	return c.Alpaca.GetInt32Response("camera", c.DeviceNumber, "binx")
}

/*
	GetBinY()

	@returns the binning factor for the Y axis.
	@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__biny
*/
func (c *Camera) GetBinY() (int32, error) {
	return c.Alpaca.GetInt32Response("camera", c.DeviceNumber, "biny")
}

/*
	GetOperationalState()

	@returns the current camera operational state (CameraIdle, CameraWaiting, CameraExposing, CameraReading, CameraDownload, CameraError)
	The operational state is specified as an integer value from the OperationalState Enum.
	@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__camerastate
*/
func (c *Camera) GetOperationalState() (OperationalState, error) {
	state, err := c.Alpaca.GetInt32Response("camera", c.DeviceNumber, "camerastate")
	return OperationalState(state), err
}

/*
	GetCCDSizeX()

	@returns the width of the CCD camera chip in unbinned pixels.
	@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__cameraxsize
*/
func (c *Camera) GetCCDSizeX() (int32, error) {
	return c.Alpaca.GetInt32Response("camera", c.DeviceNumber, "cameraxsize")
}

/*
	GetCCDSizeY()

	@returns the width of the CCD camera chip in unbinned pixels.
	@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__cameraysize
*/
func (c *Camera) GetCCDSizeY() (int32, error) {
	return c.Alpaca.GetInt32Response("camera", c.DeviceNumber, "cameraysize")
}