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
