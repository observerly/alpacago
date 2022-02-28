package alpaca

type FilterWheel struct {
	Alpaca       *ASCOMAlpacaAPIClient
	DeviceNumber uint
}

func NewFilterWheel(clientId uint32, secure bool, domain string, ip string, port int32, deviceNumber uint) *FilterWheel {
	alpaca := NewAlpacaAPI(clientId, secure, domain, ip, port)

	filterwheel := FilterWheel{
		Alpaca:       alpaca,
		DeviceNumber: deviceNumber,
	}

	return &filterwheel
}
