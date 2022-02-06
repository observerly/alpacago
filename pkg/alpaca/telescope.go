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

/*
	IsAtHome()

	@returns true if the mount is stopped in the Home position. This property will alwayts be false if the telescope does not support homing.
	@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/get_telescope__device_number__athome
*/
func (t *Telescope) IsAtHome() (bool, error) {
	return t.Alpaca.GetBooleanResponse("telescope", t.DeviceNumber, "athome")
}

/*
	IsAtPark()

	@returns true if the telescope has been put into the parked state by the seee Park() method.
	@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/get_telescope__device_number__atpark
*/
func (t *Telescope) IsAtPark() (bool, error) {
	return t.Alpaca.GetBooleanResponse("telescope", t.DeviceNumber, "atpark")
}

/*
	GetAzimuth()

	@returns the azimuth at the local horizon of the mount's current position (degrees, North-referenced, positive East/clockwise).
	@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/get_telescope__device_number__azimuth
*/
func (t *Telescope) GetAzimuth() (float64, error) {
	return t.Alpaca.GetFloat64Response("telescope", t.DeviceNumber, "azimuth")
}

/*
	CanFindHome()

	@returns true if this telescope is capable of programmed finding its home position (FindHome() method).
	@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/get_telescope__device_number__canfindhome
*/
func (t *Telescope) CanFindHome() (bool, error) {
	return t.Alpaca.GetBooleanResponse("telescope", t.DeviceNumber, "canfindhome")
}

/*
	CanPark()

	@returns true if this telescope is capable of programmed parking (Park() method)
	@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/get_telescope__device_number__canpark
*/
func (t *Telescope) CanPark() (bool, error) {
	return t.Alpaca.GetBooleanResponse("telescope", t.DeviceNumber, "canpark")
}

/*
	CanPulseGuide()

	@returns true if this telescope is capable of software-pulsed guiding (via the PulseGuide(GuideDirections, Int32) method)
	@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/get_telescope__device_number__canpulseguide
*/
func (t *Telescope) CanPulseGuide() (bool, error) {
	return t.Alpaca.GetBooleanResponse("telescope", t.DeviceNumber, "canpulseguide")
}