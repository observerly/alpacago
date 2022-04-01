package alpaca

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
	GetAveragePeriod()

	@returns the time period over which observations will be averaged.
	@see https://ascom-standards.org/api/#/ObservingConditions%20Specific%20Methods/get_observingconditions__device_number__averageperiod
*/
func (c *ObservingConditions) GetAveragePeriod() (int32, error) {
	return c.Alpaca.GetInt32Response("observingconditions", c.DeviceNumber, "averageperiod")
}
