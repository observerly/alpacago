package alpaca

import "fmt"

type ObservingConditions struct {
	Alpaca       *ASCOMAlpacaAPIClient
	DeviceNumber uint
}

func NewObservingConditions(clientId uint32, secure bool, domain string, ip string, port int32, deviceNumber uint) *ObservingConditions {
	alpaca := NewAlpacaAPI(clientId, secure, domain, ip, port)

	conditions := ObservingConditions{
		Alpaca:       alpaca,
		DeviceNumber: deviceNumber,
	}

	return &conditions
}

/*
	IsConnected() common method to all ASCOM Alpaca compliant devices

	@returns the connected state of the device
	@see https://ascom-standards.org/api/#/ASCOM%20Methods%20Common%20To%20All%20Devices/get__device_type___device_number__connected
*/
func (c *ObservingConditions) IsConnected() (bool, error) {
	return c.Alpaca.GetBooleanResponse("observingconditions", c.DeviceNumber, "connected")
}

/*
	SetConnected() common method to all ASCOM Alpaca compliant devices

	@param connected bool (set True to connect to the device hardware, set false to disconnect from the device hardware)
	@returns the connected state of the device
	@see https://ascom-standards.org/api/#/ASCOM%20Methods%20Common%20To%20All%20Devices/put__device_type___device_number__connected
*/
func (c *ObservingConditions) SetConnected(connected bool) error {
	c.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		// Set True to connect to the device hardware, set False to disconnect from the device hardware
		"Connected":           fmt.Sprintf("%t", connected),
		"ClientID":            fmt.Sprintf("%d", c.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", c.Alpaca.TransactionId),
	}

	return c.Alpaca.Put("observingconditions", c.DeviceNumber, "connected", form)
}

/*
	GetAveragePeriod()

	@returns the time period over which observations will be averaged.
	@see https://ascom-standards.org/api/#/ObservingConditions%20Specific%20Methods/get_observingconditions__device_number__averageperiod
*/
func (c *ObservingConditions) GetAveragePeriod() (int32, error) {
	return c.Alpaca.GetInt32Response("observingconditions", c.DeviceNumber, "averageperiod")
}

/*
	GetCloudCover()

	@returns the percentage of the sky obscured by cloud
	@see https://ascom-standards.org/api/#/ObservingConditions%20Specific%20Methods/get_observingconditions__device_number__cloudcover
*/
func (c *ObservingConditions) GetCloudCover() (float64, error) {
	return c.Alpaca.GetFloat64Response("observingconditions", c.DeviceNumber, "cloudcover")
}

/*
	GetDewPoint()

	@returns the atmospheric dew point at the observatory reported in °C.
	@see https://ascom-standards.org/api/#/ObservingConditions%20Specific%20Methods/get_observingconditions__device_number__dewpoint
*/
func (c *ObservingConditions) GetDewPoint() (float64, error) {
	return c.Alpaca.GetFloat64Response("observingconditions", c.DeviceNumber, "dewpoint")
}

/*
	GetHumidity()

	@returns the atmospheric humidity (%) at the observatory
	@see https://ascom-standards.org/api/#/ObservingConditions%20Specific%20Methods/get_observingconditions__device_number__humidity
*/
func (c *ObservingConditions) GetHumidity() (float64, error) {
	return c.Alpaca.GetFloat64Response("observingconditions", c.DeviceNumber, "humidity")
}

/*
	GetPressure()

	@returns the atmospheric pressure in hectoPascals at the observatory's altitude - NOT reduced to sea level.
	@see https://ascom-standards.org/api/#/ObservingConditions%20Specific%20Methods/get_observingconditions__device_number__pressure
*/
func (c *ObservingConditions) GetPressure() (float64, error) {
	return c.Alpaca.GetFloat64Response("observingconditions", c.DeviceNumber, "pressure")
}
