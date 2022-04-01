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
