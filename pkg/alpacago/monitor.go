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

/*
	IsConnected() common method to all ASCOM Alpaca compliant devices

	@returns the connected state of the device
	@see https://ascom-standards.org/api/#/ASCOM%20Methods%20Common%20To%20All%20Devices/get__device_type___device_number__connected
*/
func (m *SafetyMonitor) IsConnected() (bool, error) {
	return m.Alpaca.GetBooleanResponse("safetymonitor", m.DeviceNumber, "connected")
}
