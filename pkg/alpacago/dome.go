package alpacago

type Dome struct {
	Alpaca       *ASCOMAlpacaAPIClient
	DeviceNumber uint
}

func NewDome(clientId uint32, secure bool, domain string, ip string, port int32, deviceNumber uint) *Dome {
	alpaca := NewAlpacaAPI(clientId, secure, domain, ip, port)

	dome := Dome{
		Alpaca:       alpaca,
		DeviceNumber: deviceNumber,
	}

	return &dome
}
