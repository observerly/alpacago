package alpaca

type AlignmentMode int32

type AxisType int

type TrackingMode int

const (
	AlignmentAltAz AlignmentMode = iota
	AlignmentPolar
	AlignmentGermanPolar
)

const (
	AxisAzmRa AxisType = iota
	AxisAltDec
	AxisTertiary
)

const (
	NotTracking TrackingMode = iota
	Alt_Az
	EQ_North
	EQ_South
)

type Telescope struct {
	Alpaca       *ASCOMAlpacaAPIClient
	DeviceNumber uint
	Tracking     TrackingMode
}

func NewTelescope(clientId uint32, secure bool, domain string, ip string, port int32, deviceNumber uint, tm TrackingMode) *Telescope {
	alpaca := NewAlpacaAPI(clientId, secure, domain, ip, port)

	telescope := Telescope{
		Alpaca:       alpaca,
		DeviceNumber: deviceNumber,
		Tracking:     tm,
	}

	return &telescope
}

/*
	GetAlignmentMode()

	@returns the alignment mode of the mount (Alt/Az, Polar, German Polar).
	The alignment mode is specified as an integer value from the
	AlignmentModes Enum.
	@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/get_telescope__device_number__alignmentmode
*/
func (t *Telescope) GetAlignmentMode() (AlignmentMode, error) {
	mode, err := t.Alpaca.GetInt32Response("telescope", t.DeviceNumber, "alignmentmode")
	return AlignmentMode(mode), err
}

/*
	GetAltitude()

	@returns the altitude above the local horizon of the mount's current position (degrees, positive up)
	@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/get_telescope__device_number__altitude
*/
func (t *Telescope) GetAltitude() (float64, error) {
	return t.Alpaca.GetFloat64Response("telescope", t.DeviceNumber, "altitude")
}

/*
	GetApertureArea()

	@returns the area of the telescope's aperture, taking into account any obstructions (square meters)
	@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/get_telescope__device_number__aperturearea
*/
func (t *Telescope) GetApertureArea() (float64, error) {
	return t.Alpaca.GetFloat64Response("telescope", t.DeviceNumber, "aperturearea")
}

/*
	GetApertureDiameter()

	@returns the telescope's effective aperture diameter (meters)
	@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/get_telescope__device_number__aperturediameter
*/
func (t *Telescope) GetApertureDiameter() (float64, error) {
	return t.Alpaca.GetFloat64Response("telescope", t.DeviceNumber, "aperturearea")
}
