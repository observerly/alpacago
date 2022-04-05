package alpacago

import "fmt"

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
