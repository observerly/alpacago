package alpacago

import "fmt"

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

/*
	SetConnected() common method to all ASCOM Alpaca compliant devices

	@param connected bool (set True to connect to the device hardware, set false to disconnect from the device hardware)
	@returns the connected state of the device
	@see https://ascom-standards.org/api/#/ASCOM%20Methods%20Common%20To%20All%20Devices/put__device_type___device_number__connected
*/
func (f *FilterWheel) SetConnected(connected bool) error {
	f.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		// Set True to connect to the device hardware, set False to disconnect from the device hardware
		"Connected":           fmt.Sprintf("%t", connected),
		"ClientID":            fmt.Sprintf("%d", f.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", f.Alpaca.TransactionId),
	}

	return f.Alpaca.Put("filterwheel", f.DeviceNumber, "connected", form)
}

/*
	GetFocusOffsets()

	@returns an integer array of filter focus offsets.
	@see https://ascom-standards.org/api/#/FilterWheel%20Specific%20Methods/get_filterwheel__device_number__focusoffsets
*/
func (f *FilterWheel) GetFocusOffsets() ([]uint32, error) {
	return f.Alpaca.GetUInt32ListResponse("filterwheel", f.DeviceNumber, "focusoffsets")
}

/*
	GetNames()

	@returns the names of the filters.
	@see https://ascom-standards.org/api/#/FilterWheel%20Specific%20Methods/get_filterwheel__device_number__names
*/
func (f *FilterWheel) GetNames() ([]string, error) {
	return f.Alpaca.GetStringListResponse("filterwheel", f.DeviceNumber, "names")
}

/*
	GetPosition()

	@returns the current filter wheel position.
	@see https://ascom-standards.org/api/#/FilterWheel%20Specific%20Methods/get_filterwheel__device_number__position
*/
func (f *FilterWheel) GetPosition() (int32, error) {
	return f.Alpaca.GetInt32Response("filterwheel", f.DeviceNumber, "position")
}

/*
	SetPosition()

	@returns an error or nil, if nil it sets the filter wheel position
	@see https://ascom-standards.org/api/#/FilterWheel%20Specific%20Methods/put_filterwheel__device_number__position
*/
func (f *FilterWheel) SetPosition(position int32) error {
	f.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		"Position":            fmt.Sprintf("%d", position),
		"ClientID":            fmt.Sprintf("%d", f.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", f.Alpaca.TransactionId),
	}

	return f.Alpaca.Put("filterwheel", f.DeviceNumber, "position", form)
}
