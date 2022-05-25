package alpacago

import "fmt"

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

/*
	SetConnected() common method to all ASCOM Alpaca compliant devices

	@param connected bool (set True to connect to the device hardware, set false to disconnect from the device hardware)
	@returns the connected state of the device
	@see https://ascom-standards.org/api/#/ASCOM%20Methods%20Common%20To%20All%20Devices/put__device_type___device_number__connected
*/
func (m *SafetyMonitor) SetConnected(connected bool) error {
	m.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		// Set True to connect to the device hardware, set False to disconnect from the device hardware
		"Connected":           fmt.Sprintf("%t", connected),
		"ClientID":            fmt.Sprintf("%d", m.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", m.Alpaca.TransactionId),
	}

	return m.Alpaca.Put("safetymonitor", m.DeviceNumber, "connected", form)
}

/*
	IsSafe()

	@returns true if the state is safe, false if it is unsafe. Indicates whether the monitored state is safe for use.
	@see https://ascom-standards.org/api/#/SafetyMonitor%20Specific%20Methods/get_safetymonitor__device_number__issafe
*/
func (m *SafetyMonitor) IsSafe() (bool, error) {
	return m.Alpaca.GetBooleanResponse("safetymonitor", m.DeviceNumber, "issafe")
}
