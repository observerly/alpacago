package alpaca

type Focuser struct {
	Alpaca       *ASCOMAlpacaAPIClient
	DeviceNumber uint
}

func NewFocuser(clientId uint32, secure bool, domain string, ip string, port int32, deviceNumber uint) *Focuser {
	alpaca := NewAlpacaAPI(clientId, secure, domain, ip, port)

	focuser := Focuser{
		Alpaca:       alpaca,
		DeviceNumber: deviceNumber,
	}

	return &focuser
}
