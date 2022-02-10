package alpaca

type AlignmentMode int32

type AxisType int

type EquatorialSystem int

type PierPointingMode int32

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
	Topocentric EquatorialSystem = iota
	J2000
)

const (
	PierUnknown PierPointingMode = -1
	PierEast    PierPointingMode = 0
	PierWest    PierPointingMode = 1
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

/*
	CanSetDeclinationRate()

	@returns true if the DeclinationRate property can be changed to provide offset tracking in the declination axis.
	@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/get_telescope__device_number__cansetdeclinationrate
*/
func (t *Telescope) CanSetDeclinationRate() (bool, error) {
	return t.Alpaca.GetBooleanResponse("telescope", t.DeviceNumber, "cansetdeclinationrate")
}

/*
	CanSetGuideRates()

	@returns true if the guide rate properties used for PulseGuide(GuideDirections, Int32) can ba adjusted.
	@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/get_telescope__device_number__cansetguiderates
*/
func (t *Telescope) CanSetGuideRates() (bool, error) {
	return t.Alpaca.GetBooleanResponse("telescope", t.DeviceNumber, "cansetguiderates")
}

/*
	CanSetPark()

	@returns true if this telescope is capable of programmed parking (Park() method)
	@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/get_telescope__device_number__cansetpark
*/
func (t *Telescope) CanSetPark() (bool, error) {
	return t.Alpaca.GetBooleanResponse("telescope", t.DeviceNumber, "cansetpark")
}

/*
	CanSetPierSide()

	@returns true if the SideOfPier property can be set, meaning that the mount can be forced to flip.
	@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/get_telescope__device_number__cansetpierside
*/
func (t *Telescope) CanSetPierSide() (bool, error) {
	return t.Alpaca.GetBooleanResponse("telescope", t.DeviceNumber, "cansetpierside")
}

/*
	CanSetRightAscensionRate()

	@returns true if the RightAscensionRate property can be changed to provide offset tracking in the right ascension axis.
	@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/get_telescope__device_number__cansetrightascensionrate
*/
func (t *Telescope) CanSetRightAscensionRate() (bool, error) {
	return t.Alpaca.GetBooleanResponse("telescope", t.DeviceNumber, "cansetrightascensionrate")
}

/*
	CanSetTracking()

	@returns true if the Tracking property can be changed, turning telescope sidereal tracking on and off.
	@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/get_telescope__device_number__cansettracking
*/
func (t *Telescope) CanSetTracking() (bool, error) {
	return t.Alpaca.GetBooleanResponse("telescope", t.DeviceNumber, "cansettracking")
}

/*
	CanSlew()

	@returns true if this telescope is capable of programmed slewing (synchronous or asynchronous) to equatorial coordinates
	@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/get_telescope__device_number__canslew
*/
func (t *Telescope) CanSlew() (bool, error) {
	return t.Alpaca.GetBooleanResponse("telescope", t.DeviceNumber, "canslew")
}

/*
	CanSlewAltAz()

	@returns true if this telescope is capable of programmed slewing (synchronous or asynchronous) to local horizontal coordinates
	@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/get_telescope__device_number__canslewaltaz
*/
func (t *Telescope) CanSlewAltAz() (bool, error) {
	return t.Alpaca.GetBooleanResponse("telescope", t.DeviceNumber, "canslewaltaz")
}

/*
	CanSlewAltAzAsync()

	@returns true if this telescope is capable of programmed asynchronous slewing to local horizontal coordinates
	@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/get_telescope__device_number__canslewaltazasync
*/
func (t *Telescope) CanSlewAltAzAsync() (bool, error) {
	return t.Alpaca.GetBooleanResponse("telescope", t.DeviceNumber, "canslewaltazasync")
}

/*
	CanSlewAsync()

	@returns true if this telescope is capable of programmed asynchronous slewing to equatorial coordinates
	@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/get_telescope__device_number__canslewasync
*/
func (t *Telescope) CanSlewAsync() (bool, error) {
	return t.Alpaca.GetBooleanResponse("telescope", t.DeviceNumber, "canslewasync")
}

/*
	CanSync()

	@returns true if this telescope is capable of programmed synching to equatorial coordinates
	@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/get_telescope__device_number__cansync
*/
func (t *Telescope) CanSync() (bool, error) {
	return t.Alpaca.GetBooleanResponse("telescope", t.DeviceNumber, "cansync")
}

/*
	CanSyncAltAz()

	@returns true if this telescope is capable of programmed synching to local horizontal coordinates
	@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/get_telescope__device_number__cansyncaltaz
*/
func (t *Telescope) CanSyncAltAz() (bool, error) {
	return t.Alpaca.GetBooleanResponse("telescope", t.DeviceNumber, "cansyncaltaz")
}

/*
	CanUnPark()

	@returns true if this telescope is capable of programmed unparking (UnPark() method)
	@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/get_telescope__device_number__canunpark
*/
func (t *Telescope) CanUnPark() (bool, error) {
	return t.Alpaca.GetBooleanResponse("telescope", t.DeviceNumber, "canunpark")
}

/*
	GetDeclination()

	@returns the declination (degrees) of the mount's current equatorial coordinates, in the coordinate
	system given by the EquatorialSystem property. Reading the property will raise an error if the value
	is unavailable.
	@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/get_telescope__device_number__declination
*/
func (t *Telescope) GetDeclination() (float64, error) {
	return t.Alpaca.GetFloat64Response("telescope", t.DeviceNumber, "declination")
}

/*
	GetDeclinationRate()

	@returns the declination tracking rate (arcseconds per second, default = 0.0)
	@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/get_telescope__device_number__declinationrate
*/
func (t *Telescope) GetDeclinationRate() (float64, error) {
	return t.Alpaca.GetFloat64Response("telescope", t.DeviceNumber, "declinationrate")
}

/*
	DoesRefraction()

	@returns true if the telescope or driver applies atmospheric refraction to coordinates
	@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/get_telescope__device_number__doesrefraction
*/
func (t *Telescope) DoesRefraction() (bool, error) {
	return t.Alpaca.GetBooleanResponse("telescope", t.DeviceNumber, "doesrefraction")
}

/*
	GetEquatorialSystem()

	@returns the current equatorial coordinate system used by this telescope (e.g. Topocentric or J2000).
	@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/get_telescope__device_number__equatorialsystem
*/
func (t *Telescope) GetEquatorialSystem() (EquatorialSystem, error) {
	system, err := t.Alpaca.GetInt32Response("telescope", t.DeviceNumber, "equatorialsystem")
	return EquatorialSystem(system), err
}

/*
	GetFocalLength()

	@returns the telescope's focal length in meters
	@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/get_telescope__device_number__focallength
*/
func (t *Telescope) GetFocalLength() (float64, error) {
	return t.Alpaca.GetFloat64Response("telescope", t.DeviceNumber, "focallength")
}

/*
	IsPulseGuiding()

	@returns true if a PulseGuide(GuideDirections, Int32) command is in progress, false otherwise
	@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/get_telescope__device_number__ispulseguiding
*/
func (t *Telescope) IsPulseGuiding() (bool, error) {
	return t.Alpaca.GetBooleanResponse("telescope", t.DeviceNumber, "ispulseguiding")
}

/*
	GetRightAscension()

	@returns the right ascension (hours) of the mount's current equatorial coordinates, in the coordinate
	system given by the EquatorialSystem property
	@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/get_telescope__device_number__rightascension
*/
func (t *Telescope) GetRightAscension() (float64, error) {
	return t.Alpaca.GetFloat64Response("telescope", t.DeviceNumber, "rightascension")
}

/*
	GetRightAscensionRate()

	@returns the right ascension tracking rate (arcseconds per second, default = 0.0)
	@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/get_telescope__device_number__rightascensionrate
*/
func (t *Telescope) GetRightAscensionRate() (float64, error) {
	return t.Alpaca.GetFloat64Response("telescope", t.DeviceNumber, "rightascensionrate")
}

/*
	GetSideOfPier()

	@returns the pointing state of the mount. 0 = pierEast, 1 = pierWest, -1= pierUnknown
	@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/get_telescope__device_number__sideofpier
*/
func (t *Telescope) GetSideOfPier() (PierPointingMode, error) {
	mode, err := t.Alpaca.GetInt32Response("telescope", t.DeviceNumber, "sideofpier")
	return PierPointingMode(mode), err
}

/*
	GetSiderealTime()

	@returns the local apparent sidereal time from the telescope's internal clock (hours, sidereal).
	@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/get_telescope__device_number__siderealtime
*/
func (t *Telescope) GetSiderealTime() (float64, error) {
	return t.Alpaca.GetFloat64Response("telescope", t.DeviceNumber, "siderealtime")
}

/*
	GetSiteElevation()

	@returns the elevation above mean sea level (meters) of the site at which the telescope is located.
	@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/get_telescope__device_number__siteelevation
*/
func (t *Telescope) GetSiteElevation() (float64, error) {
	return t.Alpaca.GetFloat64Response("telescope", t.DeviceNumber, "siteelevation")
}

/*
	GetSiteLatitude()

	@returns the geodetic(map) latitude (degrees, positive North, WGS84) of the site at which the telescope is located.
	@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/get_telescope__device_number__sitelatitude
*/
func (t *Telescope) GetSiteLatitude() (float64, error) {
	return t.Alpaca.GetFloat64Response("telescope", t.DeviceNumber, "sitelatitude")
}

/*
	GetSiteLongitude()

	@returns the geodetic(map) longitude (degrees, positive East, WGS84) of the site at which the telescope is located.
	@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/get_telescope__device_number__sitelongitude
*/
func (t *Telescope) GetSiteLongitude() (float64, error) {
	return t.Alpaca.GetFloat64Response("telescope", t.DeviceNumber, "sitelongitude")
}

/*
	IsSlewing()

	@returns true  if telescope is currently moving in response to one of the Slew methods or the MoveAxis(TelescopeAxes,
	Double) method, false at all other times.
	@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/get_telescope__device_number__slewing
*/
func (t *Telescope) IsSlewing() (bool, error) {
	return t.Alpaca.GetBooleanResponse("telescope", t.DeviceNumber, "slewing")
}

/*
	GetSlewSettleTime()

	@returns the post-slew settling time (in seconds).
	@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/get_telescope__device_number__slewsettletime
*/
func (t *Telescope) GetSlewSettleTime() (int32, error) {
	return t.Alpaca.GetInt32Response("telescope", t.DeviceNumber, "slewsettletime")
}

/*
	GetTargetDeclination()

	@returns the declination (degrees, positive North) for the target of an equatorial slew or sync operation.
	@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/get_telescope__device_number__targetdeclination
*/
func (t *Telescope) GetTargetDeclination() (float64, error) {
	return t.Alpaca.GetFloat64Response("telescope", t.DeviceNumber, "targetdeclination")
}

/*
	GetTargetRightAscension()

	@returns the right ascension (hours) for the target of an equatorial slew or sync operation
	@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/get_telescope__device_number__targetrightascension
*/
func (t *Telescope) GetTargetRightAscension() (float64, error) {
	return t.Alpaca.GetFloat64Response("telescope", t.DeviceNumber, "targetrightascension")
}
