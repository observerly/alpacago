package alpacago

type CoverCalibrator struct {
	Alpaca       *ASCOMAlpacaAPIClient
	DeviceNumber uint
}

func NewCoverCalibrator(clientId uint32, secure bool, domain string, ip string, port int32, deviceNumber uint) *CoverCalibrator {
	alpaca := NewAlpacaAPI(clientId, secure, domain, ip, port)

	calibrator := CoverCalibrator{
		Alpaca:       alpaca,
		DeviceNumber: deviceNumber,
	}

	return &calibrator
}
