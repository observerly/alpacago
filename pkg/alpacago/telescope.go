package alpacago

import (
	"errors"
	"fmt"
	"time"
)

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

// String returns the string representation of the AlignmentSystem value.
func (am AlignmentMode) String() string {
	switch am {
	case AlignmentAltAz:
		return "Altitude-Azimuth"
	case AlignmentPolar:
		return "Polar"
	case AlignmentGermanPolar:
		return "German Polar"
	default:
		return fmt.Sprintf("Unknown AlignmentSystem value: %d", am)
	}
}

const (
	AxisAzmRa AxisType = iota
	AxisAltDec
	AxisTertiary
)

const (
	Topocentric EquatorialSystem = iota
	J2000
)

// String returns the string representation of the EquatorialSystem value.
func (es EquatorialSystem) String() string {
	switch es {
	case Topocentric:
		return "Topocentric"
	case J2000:
		return "J2000"
	default:
		return fmt.Sprintf("Unknown EquatorialSystem value: %d", es)
	}
}

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

type AxisRatesResponse struct {
	Value               []map[string]float64 `json:"Value"`
	ClientTransactionID int32                `json:"ClientTransactionID"`
	ServerTransactionID int32                `json:"ServerTransactionID"`
	ErrorNumber         int32                `json:"ErrorNumber"`
	ErrorMessage        string               `json:"ErrorMessage"`
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
GetDescription() common method to all ASCOM Alpaca compliant devices

@returns the description of the device
@see https://ascom-standards.org/api/#/ASCOM%20Methods%20Common%20To%20All%20Devices/get__device_type___device_number__description
*/
func (t *Telescope) GetDescription() (string, error) {
	return t.Alpaca.GetDescription("telescope", t.DeviceNumber)
}

/*
IsConnected() common method to all ASCOM Alpaca compliant devices

@returns the connected state of the device
@see https://ascom-standards.org/api/#/ASCOM%20Methods%20Common%20To%20All%20Devices/get__device_type___device_number__connected
*/
func (t *Telescope) IsConnected() (bool, error) {
	return t.Alpaca.GetBooleanResponse("telescope", t.DeviceNumber, "connected")
}

/*
SetConnected() common method to all ASCOM Alpaca compliant devices

@param connected bool (set True to connect to the device hardware, set false to disconnect from the device hardware)
@returns the connected state of the device
@see https://ascom-standards.org/api/#/ASCOM%20Methods%20Common%20To%20All%20Devices/put__device_type___device_number__connected
*/
func (t *Telescope) SetConnected(connected bool) error {
	t.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		// Set True to connect to the device hardware, set False to disconnect from the device hardware
		"Connected":           fmt.Sprintf("%t", connected),
		"ClientID":            fmt.Sprintf("%d", t.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", t.Alpaca.TransactionId),
	}

	return t.Alpaca.Put("telescope", t.DeviceNumber, "connected", form)
}

/*
SetAbortSlew()

@returns an error or nil, if nil immediately Stops a slew in progress.
@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/put_telescope__device_number__abortslew
*/
func (t *Telescope) SetAbortSlew() error {
	t.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		"ClientID":            fmt.Sprintf("%d", t.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", t.Alpaca.TransactionId),
	}

	return t.Alpaca.Put("telescope", t.DeviceNumber, "abortslew", form)
}

/*
GetAlignmentMode()

@returns the alignment mode of the mount (Alt/Az, Polar, German Polar).
The alignment mode is specified as an integer value from the
AlignmentModes Enum.
@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/get_telescope__device_number__alignmentmode
*/
func (t *Telescope) GetAlignmentMode() (string, error) {
	mode, err := t.Alpaca.GetInt32Response("telescope", t.DeviceNumber, "alignmentmode")
	return AlignmentMode(mode).String(), err
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
GetAxisRates()

@returns the rates at which the telescope may be moved about the specified axis by the MoveAxis(TelescopeAxes, Double) method.
@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/get_telescope__device_number__axisrates
*/
func (t *Telescope) GetAxisRates(axis AxisType) (map[string]float64, error) {
	url := t.Alpaca.getEndpoint("telescope", t.DeviceNumber, "axisrates")

	querystring := fmt.Sprintf("axis=%d&%s", axis, t.Alpaca.getQueryString())

	// Setup the resty client:
	resp, err := t.Alpaca.Client.R().SetResult(&AxisRatesResponse{}).SetQueryString(querystring).SetHeader("Accept", "application/json").Get(url)

	// If the response object has a REST error:
	if err != nil {
		return map[string]float64{}, err
	}

	// Return the result:
	result := (resp.Result().(*AxisRatesResponse))

	return result.Value[0], nil
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
FindHome()

@returns an error or nil, if nil it puts the telescope into the homed state.
@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/put_telescope__device_number__findhome
*/
func (t *Telescope) FindHome() error {
	t.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		"ClientID":            fmt.Sprintf("%d", t.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", t.Alpaca.TransactionId),
	}

	return t.Alpaca.Put("telescope", t.DeviceNumber, "findhome", form)
}

/*
CanMoveAxis()

@returns true if this telescope can move the requested axis.
@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/get_telescope__device_number__canmoveaxis
*/
func (t *Telescope) CanMoveAxis(axis AxisType) (bool, error) {
	url := t.Alpaca.getEndpoint("telescope", t.DeviceNumber, "canmoveaxis")

	querystring := fmt.Sprintf("axis=%d&%s", axis, t.Alpaca.getQueryString())

	// Setup the resty client:
	resp, err := t.Alpaca.Client.R().SetResult(&booleanResponse{}).SetQueryString(querystring).SetHeader("Accept", "application/json").Get(url)

	if err != nil {
		return false, err
	}

	// If the response object has a REST error:
	if resp.IsError() {
		t.Alpaca.ErrorNumber = resp.StatusCode()
		t.Alpaca.ErrorMessage = resp.String()
	}

	// Return the result:
	result := (resp.Result().(*booleanResponse))

	return result.Value, nil
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
SetPark()

@returns an error or nil, if nil it takes telescope out of the Parked state.
@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/put_telescope__device_number__unpark
*/
func (t *Telescope) SetPark() error {
	t.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		"ClientID":            fmt.Sprintf("%d", t.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", t.Alpaca.TransactionId),
	}

	return t.Alpaca.Put("telescope", t.DeviceNumber, "park", form)
}

/*
/*
SetUnPark()

@returns an error or nil, if nil it takes telescope out of the Parked state.
@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/put_telescope__device_number__unpark
*/
func (t *Telescope) SetUnPark() error {
	t.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		"ClientID":            fmt.Sprintf("%d", t.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", t.Alpaca.TransactionId),
	}

	return t.Alpaca.Put("telescope", t.DeviceNumber, "unpark", form)
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
SetDeclinationRate()

@returns an error or nil, if nil it sets the declination tracking rate (arcseconds per second)
@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/put_telescope__device_number__declinationrate
*/
func (t *Telescope) SetDeclinationRate(declinationRate float64) error {
	t.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		"DeclinationRate":     fmt.Sprintf("%f", declinationRate),
		"ClientID":            fmt.Sprintf("%d", t.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", t.Alpaca.TransactionId),
	}

	return t.Alpaca.Put("telescope", t.DeviceNumber, "declinationrate", form)
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
SetDoesRefraction()

Determines whether atmospheric refraction is applied to coordinates.

@returns an error or nil, if nil causes the rotator to move Position degrees relative to the current position value.
@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/put_telescope__device_number__doesrefraction
*/
func (t *Telescope) SetDoesRefraction(doesRefraction bool) error {
	t.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		"DoesRefraction":      fmt.Sprintf("%t", doesRefraction),
		"ClientID":            fmt.Sprintf("%d", t.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", t.Alpaca.TransactionId),
	}

	return t.Alpaca.Put("telescope", t.DeviceNumber, "doesrefraction", form)
}

/*
GetEquatorialSystem()

@returns the current equatorial coordinate system used by this telescope (e.g. Topocentric or J2000).
@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/get_telescope__device_number__equatorialsystem
*/
func (t *Telescope) GetEquatorialSystem() (string, error) {
	system, err := t.Alpaca.GetInt32Response("telescope", t.DeviceNumber, "equatorialsystem")
	return EquatorialSystem(system).String(), err
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
SetRightAscensionRate()

@returns an error or nil, if nil it sets the right ascension tracking rate (arcseconds per second)
@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/put_telescope__device_number__rightascensionrate
*/
func (t *Telescope) SetRightAscensionRate(rightAscensionRate float64) error {
	t.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		"RightAscensionRate":  fmt.Sprintf("%f", rightAscensionRate),
		"ClientID":            fmt.Sprintf("%d", t.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", t.Alpaca.TransactionId),
	}

	return t.Alpaca.Put("telescope", t.DeviceNumber, "rightascensionrate", form)
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
SetSideOfPier()

@returns an error or nil, if nil it sets the pointing state of the mount. 0 = pierEast, 1 = pierWest
@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/put_telescope__device_number__sideofpier
*/
func (t *Telescope) SetSideOfPier(sideOfPier PierPointingMode) error {
	t.Alpaca.TransactionId++

	if sideOfPier != 1 && sideOfPier != 0 {
		return errors.New("please provide a valid pointing state for the mount e.g., eiher 0 = pierEast, 1 = pierWest")
	}

	var form map[string]string = map[string]string{
		"SideOfPier":          fmt.Sprintf("%d", sideOfPier),
		"ClientID":            fmt.Sprintf("%d", t.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", t.Alpaca.TransactionId),
	}

	return t.Alpaca.Put("telescope", t.DeviceNumber, "sideofpier", form)
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
SetSiteElevation()

@params siteElevation - the site elevation above mean sea level (metres).
@returns an error or nil, if nil it sets the elevation above mean sea level (metres) of the site at which the telescope is located.
@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/put_telescope__device_number__siteelevation
*/
func (t *Telescope) SetSiteElevation(siteElevation float64) error {
	t.Alpaca.TransactionId++

	if siteElevation < -1000 || siteElevation > 10000 {
		return errors.New("please provide a realistic site elevation, e.g., greater than or equal to -1000m, but less than 10000m relative to mean sea level")
	}

	var form map[string]string = map[string]string{
		"SiteElevation":       fmt.Sprintf("%f", siteElevation),
		"ClientID":            fmt.Sprintf("%d", t.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", t.Alpaca.TransactionId),
	}

	return t.Alpaca.Put("telescope", t.DeviceNumber, "siteelevation", form)
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
SetSiteLatitude()

@returns an error or nil, if nil it sets the observing site's latitude (degrees).
@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/put_telescope__device_number__sitelatitude
*/
func (t *Telescope) SetSiteLatitude(siteLatitude float64) error {
	t.Alpaca.TransactionId++

	if siteLatitude <= -90 || siteLatitude >= 90 {
		return errors.New("please provide a valid latitude between -90° and +90°")
	}

	var form map[string]string = map[string]string{
		"SiteLatitude":        fmt.Sprintf("%f", siteLatitude),
		"ClientID":            fmt.Sprintf("%d", t.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", t.Alpaca.TransactionId),
	}

	return t.Alpaca.Put("telescope", t.DeviceNumber, "sitelatitude", form)
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
SetSiteLongitude()

@returns an error or nil, if nil it sets the observing site's longitude (degrees, positive East, WGS84).
@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/put_telescope__device_number__sitelongitude
*/
func (t *Telescope) SetSiteLongitude(siteLongitude float64) error {
	t.Alpaca.TransactionId++

	if siteLongitude <= -180 || siteLongitude >= 180 {
		return errors.New("please provide a valid longitude between -180° and +180°")
	}

	var form map[string]string = map[string]string{
		"SiteLongitude":       fmt.Sprintf("%f", siteLongitude),
		"ClientID":            fmt.Sprintf("%d", t.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", t.Alpaca.TransactionId),
	}

	return t.Alpaca.Put("telescope", t.DeviceNumber, "sitelongitude", form)
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
SetSlewSettleTime()

@returns an error or nil, if nil it sets the post-slew settling time (integer sec.).
@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/put_telescope__device_number__slewsettletime
*/
func (t *Telescope) SetSlewSettleTime(slewSettleTime int32) error {
	t.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		"SlewSettleTime":      fmt.Sprintf("%d", slewSettleTime),
		"ClientID":            fmt.Sprintf("%d", t.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", t.Alpaca.TransactionId),
	}

	return t.Alpaca.Put("telescope", t.DeviceNumber, "slewsettletime", form)
}

/*
SetSlewToAltAz

@returns an error or nil, if nil it moves the telescope to the given local horizontal coordinates, return when slew is complete
@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/put_telescope__device_number__slewtoaltaz
*/
func (t *Telescope) SetSlewToAltAz(altitude float64, azimuth float64) error {
	t.Alpaca.TransactionId++

	t.SetTracking(false)

	if altitude < -90 || altitude > 90 {
		return errors.New("please provide a valid altitude between -90° and +90°")
	}

	if azimuth < 0 || azimuth > 360 {
		return errors.New("please provide a valid azimuth between 0° and +360°")
	}

	var form map[string]string = map[string]string{
		"Altitude":            fmt.Sprintf("%f", altitude),
		"Azimuth":             fmt.Sprintf("%f", azimuth),
		"ClientID":            fmt.Sprintf("%d", t.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", t.Alpaca.TransactionId),
	}

	return t.Alpaca.Put("telescope", t.DeviceNumber, "slewtoaltaz", form)
}

/*
SetSlewToAltAzAsync

@returns an error or nil, if nil it moves the telescope to the given local horizontal coordinates, return immediately after
the slew starts. The client can poll the Slewing method to determine when the mount reaches the intended coordinates.
@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/put_telescope__device_number__slewtoaltazasync
*/
func (t *Telescope) SetSlewToAltAzAsync(altitude float64, azimuth float64) error {
	t.Alpaca.TransactionId++

	t.SetTracking(false)

	if altitude < -90 || altitude > 90 {
		return errors.New("please provide a valid altitude between -90° and +90°")
	}

	if azimuth < 0 || azimuth > 360 {
		return errors.New("please provide a valid azimuth between 0° and +360°")
	}

	var form map[string]string = map[string]string{
		"Altitude":            fmt.Sprintf("%f", altitude),
		"Azimuth":             fmt.Sprintf("%f", azimuth),
		"ClientID":            fmt.Sprintf("%d", t.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", t.Alpaca.TransactionId),
	}

	return t.Alpaca.Put("telescope", t.DeviceNumber, "slewtoaltazasync", form)
}

/*
SetSlewToCoordinates

@returns an error or nil, if nil it moves the telescope to the given local horizontal coordinates, return immediately after
the slew starts. The client can poll the Slewing method to determine when the mount reaches the intended coordinates.
@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/put_telescope__device_number__slewtocoordinates
*/
func (t *Telescope) SetSlewToCoordinates(rightAscension float64, declination float64) error {
	t.Alpaca.TransactionId++

	t.SetTracking(true)

	if declination < -90 || declination > 90 {
		return errors.New("please provide a valid altitude between -90° and +90°")
	}

	if rightAscension < 0 || rightAscension > 360 {
		return errors.New("please provide a valid azimuth between 0° and +360°")
	}

	rightAscension /= 15

	var form map[string]string = map[string]string{
		"RightAscension":      fmt.Sprintf("%f", rightAscension),
		"Declination":         fmt.Sprintf("%f", declination),
		"ClientID":            fmt.Sprintf("%d", t.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", t.Alpaca.TransactionId),
	}

	return t.Alpaca.Put("telescope", t.DeviceNumber, "slewtocoordinates", form)
}

/*
SetSlewToCoordinatesAsync

@returns an error or nil, if nil it moves the telescope to the given equatorial coordinates, return immediately after
the slew starts. The client can poll the Slewing method to determine when the mount reaches the intended coordinates.
@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/put_telescope__device_number__slewtocoordinatesasync
*/
func (t *Telescope) SetSlewToCoordinatesAsync(rightAscension float64, declination float64) error {
	t.Alpaca.TransactionId++

	t.SetTracking(true)

	if declination < -90 || declination > 90 {
		return errors.New("please provide a valid altitude between -90° and +90°")
	}

	if rightAscension < 0 || rightAscension > 360 {
		return errors.New("please provide a valid azimuth between 0° and +360°")
	}

	rightAscension /= 15

	var form map[string]string = map[string]string{
		"RightAscension":      fmt.Sprintf("%f", rightAscension),
		"Declination":         fmt.Sprintf("%f", declination),
		"ClientID":            fmt.Sprintf("%d", t.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", t.Alpaca.TransactionId),
	}

	return t.Alpaca.Put("telescope", t.DeviceNumber, "slewtocoordinatesasync", form)
}

/*
SetSlewToTarget

@returns an error or nil, if nil it moves the telescope to the TargetRightAscension and TargetDeclination
equatorial coordinates, return when slew is complete
@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/put_telescope__device_number__slewtotarget
*/
func (t *Telescope) SetSlewToTarget() error {
	t.Alpaca.TransactionId++

	t.SetTracking(true)

	var form map[string]string = map[string]string{
		"ClientID":            fmt.Sprintf("%d", t.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", t.Alpaca.TransactionId),
	}

	return t.Alpaca.Put("telescope", t.DeviceNumber, "slewtotarget", form)
}

/*
SetSlewToTargetAsync

@returns an error or nil, if nil it moves the telescope to the TargetRightAscension and TargetDeclination equatorial coordinates,
return immediately after the slew starts. The client can poll the Slewing method to determine when the mount reaches the
intended coordinates.
@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/put_telescope__device_number__slewtotargetasync
*/
func (t *Telescope) SetSlewToTargetAsync() error {
	t.Alpaca.TransactionId++

	t.SetTracking(true)

	var form map[string]string = map[string]string{
		"ClientID":            fmt.Sprintf("%d", t.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", t.Alpaca.TransactionId),
	}

	return t.Alpaca.Put("telescope", t.DeviceNumber, "slewtotargetasync", form)
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
SetTargetDeclination()

@returns an error or nil, if nil it sets the declination (degrees, positive North) for the target of an equatorial slew or sync operation.
@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/put_telescope__device_number__targetdeclination
*/
func (t *Telescope) SetTargetDeclination(targetDeclination float64) error {
	t.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		"TargetDeclination":   fmt.Sprintf("%f", targetDeclination),
		"ClientID":            fmt.Sprintf("%d", t.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", t.Alpaca.TransactionId),
	}

	return t.Alpaca.Put("telescope", t.DeviceNumber, "targetdeclination", form)
}

/*
GetTargetRightAscension()

@returns the right ascension (hours) for the target of an equatorial slew or sync operation
@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/get_telescope__device_number__targetrightascension
*/
func (t *Telescope) GetTargetRightAscension() (float64, error) {
	return t.Alpaca.GetFloat64Response("telescope", t.DeviceNumber, "targetrightascension")
}

/*
SetTargetRightAscension()

@returns  an error or nil, if nil it the right ascension (hours) for the target of an equatorial slew or sync operation
@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/put_telescope__device_number__targetrightascension
*/
func (t *Telescope) SetTargetRightAscension(targetRightAscension float64) error {
	t.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		"TargetRightAscension": fmt.Sprintf("%f", targetRightAscension),
		"ClientID":             fmt.Sprintf("%d", t.Alpaca.ClientId),
		"ClientTransactionID":  fmt.Sprintf("%d", t.Alpaca.TransactionId),
	}

	return t.Alpaca.Put("telescope", t.DeviceNumber, "targetrightascension", form)
}

/*
IsTracking()

@returns the state of the telescope's sidereal tracking drive.
@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/get_telescope__device_number__tracking
*/
func (t *Telescope) IsTracking() (bool, error) {
	return t.Alpaca.GetBooleanResponse("telescope", t.DeviceNumber, "tracking")
}

/*
SetTracking()

@returns an error or nil, if nil it sets the state of the telescope's sidereal tracking drive.
@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/put_telescope__device_number__tracking
*/
func (t *Telescope) SetTracking(tracking bool) error {
	t.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		"Tracking":            fmt.Sprintf("%t", tracking),
		"ClientID":            fmt.Sprintf("%d", t.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", t.Alpaca.TransactionId),
	}

	return t.Alpaca.Put("telescope", t.DeviceNumber, "tracking", form)
}

/*
GetTrackingRate()

@returns the current tracking rate of the telescope's sidereal drive.
@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/get_telescope__device_number__trackingrate
*/
func (t *Telescope) GetTrackingRate() (int32, error) {
	return t.Alpaca.GetInt32Response("telescope", t.DeviceNumber, "trackingrate")
}

/*
GetUTCDate()

@returns the UTC date/time of the telescope's internal clock in ISO 8601 format including fractional
seconds. The general format (in Microsoft custom date format style) is yyyy-MM-ddTHH:mm:ss.fffffffZ
e.g. 2016-03-04T17:45:31.1234567Z or 2016-11-14T07:03:08.1234567Z Please note the compulsary trailing
Z indicating the 'Zulu', UTC time zone.
@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/get_telescope__device_number__utcdate
*/
func (t *Telescope) GetUTCDate() (time.Time, error) {
	utc, err := t.Alpaca.GetStringResponse("telescope", t.DeviceNumber, "utcdate")
	date, _ := time.Parse("yyyy-MM-ddTHH:mm:ss.fffffffZ", utc)
	return date, err
}

/*
SetUTCDate()

@returns an error or nil, if nil it sets the UTC date/time of the telescope's internal clock in ISO 8601 format including fractional
seconds. The general format (in Microsoft custom date format style) is yyyy-MM-ddTHH:mm:ss.fffffffZ
e.g. 2016-03-04T17:45:31.1234567Z or 2016-11-14T07:03:08.1234567Z Please note the compulsary trailing
Z indicating the 'Zulu', UTC time zone.
@see https://ascom-standards.org/api/#/Telescope%20Specific%20Methods/put_telescope__device_number__utcdate
*/
func (t *Telescope) SetUTCDate(UTCDate time.Time) error {
	t.Alpaca.TransactionId++

	// Don't ask, just read: https://go.dev/src/time/format.go
	date := UTCDate.Format("2006-01-02T15:04:05.000000000Z")

	var form map[string]string = map[string]string{
		"UTCDate":             fmt.Sprintf("%q", date),
		"ClientID":            fmt.Sprintf("%d", t.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", t.Alpaca.TransactionId),
	}

	return t.Alpaca.Put("telescope", t.DeviceNumber, "utcdate", form)
}
