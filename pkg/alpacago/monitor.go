package alpacago

type SafetyMonitor struct {
	Alpaca       *ASCOMAlpacaAPIClient
	DeviceNumber uint
}

func NewSafetyMonitor(clientId uint32, secure bool, domain string, ip string, port int32, deviceNumber uint) *SafetyMonitor {
	alpaca := NewAlpacaAPI(clientId, secure, domain, ip, port)

	monitor := SafetyMonitor{
		Alpaca:       alpaca,
		DeviceNumber: deviceNumber,
	}

	return &monitor
}
