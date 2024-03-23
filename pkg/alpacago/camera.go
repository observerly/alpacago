package alpacago

import (
	"fmt"
	"strconv"
	"time"
)

type OperationalState int32

const (
	CameraIdle OperationalState = iota
	CameraWaiting
	CameraExposing
	CameraReading
	CameraDownload
	CameraError
)

func (s OperationalState) String() string {
	name := []string{"idle", "waiting", "exposing", "reading", "download", "error"}

	i := uint8(s)

	switch {
	case i <= uint8(Error):
		return name[i]
	default:
		return strconv.Itoa(int(i))
	}
}

type SensorType int32

const (
	Monochrome SensorType = iota
	ColourNotRequiringDecoding
	RGGBBayerEncoding
	CMYGBayerEncoding
	CMYG2BayerEncoding
	LRGBTRUESENSEBayerEncoding
)

func (s SensorType) String() string {
	name := []string{"Monochrome", "Color", "RGGB", "CMYG", "CMYG2", "LRGB"}

	i := uint8(s)

	switch {
	case i <= uint8(LRGBTRUESENSEBayerEncoding):
		return name[i]
	default:
		return strconv.Itoa(int(i))
	}
}

type Camera struct {
	Alpaca       *ASCOMAlpacaAPIClient
	DeviceNumber uint
}

func NewCamera(clientId uint32, secure bool, domain string, ip string, port int32, deviceNumber uint) *Camera {
	alpaca := NewAlpacaAPI(clientId, secure, domain, ip, port)

	camera := Camera{
		Alpaca:       alpaca,
		DeviceNumber: deviceNumber,
	}

	return &camera
}

/*
IsConnected() common method to all ASCOM Alpaca compliant devices

@returns the connected state of the device
@see https://ascom-standards.org/api/#/ASCOM%20Methods%20Common%20To%20All%20Devices/get__device_type___device_number__connected
*/
func (c *Camera) IsConnected() (bool, error) {
	return c.Alpaca.GetBooleanResponse("camera", c.DeviceNumber, "connected")
}

/*
SetConnected() common method to all ASCOM Alpaca compliant devices

@param connected bool (set True to connect to the device hardware, set false to disconnect from the device hardware)
@returns the connected state of the device
@see https://ascom-standards.org/api/#/ASCOM%20Methods%20Common%20To%20All%20Devices/put__device_type___device_number__connected
*/
func (c *Camera) SetConnected(connected bool) error {
	c.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		// Set True to connect to the device hardware, set False to disconnect from the device hardware
		"Connected":           fmt.Sprintf("%t", connected),
		"ClientID":            fmt.Sprintf("%d", c.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", c.Alpaca.TransactionId),
	}

	return c.Alpaca.Put("camera", c.DeviceNumber, "connected", form)
}

/*
GetBayerOffsetX()

@returns the X offset of the Bayer matrix, as defined in SensorType.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__bayeroffsetx
*/
func (c *Camera) GetBayerOffsetX() (int32, error) {
	return c.Alpaca.GetInt32Response("camera", c.DeviceNumber, "bayeroffsetx")
}

/*
GetBayerOffsetY()

@returns the Y offset of the Bayer matrix, as defined in SensorType.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__bayeroffsety
*/
func (c *Camera) GetBayerOffsetY() (int32, error) {
	return c.Alpaca.GetInt32Response("camera", c.DeviceNumber, "bayeroffsety")
}

/*
GetBinX()

@returns the binning factor for the X axis.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__binx
*/
func (c *Camera) GetBinX() (int32, error) {
	return c.Alpaca.GetInt32Response("camera", c.DeviceNumber, "binx")
}

/*
SetBinX()

@returns an error or nil, if nil it sets the binning factor for the X axis.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/put_camera__device_number__binx
*/
func (c *Camera) SetBinX(binX int32) error {
	c.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		"BinX":                fmt.Sprintf("%d", binX),
		"ClientID":            fmt.Sprintf("%d", c.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", c.Alpaca.TransactionId),
	}

	return c.Alpaca.Put("camera", c.DeviceNumber, "binx", form)
}

/*
GetBinY()

@returns the binning factor for the Y axis.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__biny
*/
func (c *Camera) GetBinY() (int32, error) {
	return c.Alpaca.GetInt32Response("camera", c.DeviceNumber, "biny")
}

/*
SetBinY()

@returns an error or nil, if nil it sets the binning factor for the Y axis.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/put_camera__device_number__biny
*/
func (c *Camera) SetBinY(binY int32) error {
	c.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		"BinY":                fmt.Sprintf("%d", binY),
		"ClientID":            fmt.Sprintf("%d", c.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", c.Alpaca.TransactionId),
	}

	return c.Alpaca.Put("camera", c.DeviceNumber, "biny", form)
}

/*
GetOperationalState()

@returns the current camera operational state (CameraIdle, CameraWaiting, CameraExposing, CameraReading, CameraDownload, CameraError)
The operational state is specified as an integer value from the OperationalState Enum.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__camerastate
*/
func (c *Camera) GetOperationalState() (string, error) {
	state, err := c.Alpaca.GetInt32Response("camera", c.DeviceNumber, "camerastate")
	return OperationalState(state).String(), err
}

/*
GetCCDSizeX()

@returns the width of the CCD camera chip in unbinned pixels.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__cameraxsize
*/
func (c *Camera) GetCCDSizeX() (int32, error) {
	return c.Alpaca.GetInt32Response("camera", c.DeviceNumber, "cameraxsize")
}

/*
GetCCDSizeY()

@returns the width of the CCD camera chip in unbinned pixels.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__cameraysize
*/
func (c *Camera) GetCCDSizeY() (int32, error) {
	return c.Alpaca.GetInt32Response("camera", c.DeviceNumber, "cameraysize")
}

/*
CanAbortExposure()

@returns true if the camera can abort exposures; false if not.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__canabortexposure
*/
func (c *Camera) CanAbortExposure() (bool, error) {
	return c.Alpaca.GetBooleanResponse("camera", c.DeviceNumber, "canabortexposure")
}

/*
CanFastReadout()

@returns true if the camera has a fast readout mode.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__canfastreadout
*/
func (c *Camera) CanFastReadout() (bool, error) {
	return c.Alpaca.GetBooleanResponse("camera", c.DeviceNumber, "canfastreadout")
}

/*
CanAsymmetricBin()

@returns true if the camera supports asymmetric binning
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__canasymmetricbin
*/
func (c *Camera) CanAsymmetricBin() (bool, error) {
	return c.Alpaca.GetBooleanResponse("camera", c.DeviceNumber, "canasymmetricbin")
}

/*
CanGetCoolerPower()

@returns true if the camera's cooler power setting can be read.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__cangetcoolerpower
*/
func (c *Camera) CanGetCoolerPower() (bool, error) {
	return c.Alpaca.GetBooleanResponse("camera", c.DeviceNumber, "cangetcoolerpower")
}

/*
CanPulseGuide()

@returns true if the camera supports pulse guiding.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__canpulseguide
*/
func (c *Camera) CanPulseGuide() (bool, error) {
	return c.Alpaca.GetBooleanResponse("camera", c.DeviceNumber, "canpulseguide")
}

/*
CanSetCCDTemperature()

@returns true if the camera supports setting the CCD temperature.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__cansetccdtemperature
*/
func (c *Camera) CanSetCCDTemperature() (bool, error) {
	return c.Alpaca.GetBooleanResponse("camera", c.DeviceNumber, "cansetccdtemperature")
}

/*
CanStopExposure()

@returns true if the camera can stop an exposure that is in progress
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__canstopexposure
*/
func (c *Camera) CanStopExposure() (bool, error) {
	return c.Alpaca.GetBooleanResponse("camera", c.DeviceNumber, "canstopexposure")
}

/*
GetCCDTemperature()

@returns true if the camera can stop an exposure that is in progress
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__ccdtemperature
*/
func (c *Camera) GetCCDTemperature() (float64, error) {
	return c.Alpaca.GetFloat64Response("camera", c.DeviceNumber, "ccdtemperature")
}

/*
IsCoolerOn()

@returns true if the camera cooler is one, else false
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__cooleron
*/
func (c *Camera) IsCoolerOn() (bool, error) {
	return c.Alpaca.GetBooleanResponse("camera", c.DeviceNumber, "cooleron")
}

/*
TurnCoolerOn()

@returns error if there was a problem turning the cooler on
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/put_camera__device_number__cooleron
*/
func (c *Camera) TurnCoolerOn() error {
	c.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		// Set True to turn the camera cooler on:
		"CoolerOn":            fmt.Sprintf("%t", true),
		"ClientID":            fmt.Sprintf("%d", c.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", c.Alpaca.TransactionId),
	}

	return c.Alpaca.Put("camera", c.DeviceNumber, "cooleron", form)
}

/*
TurnCoolerOff()

@returns error if there was a problem turning the cooler off
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/put_camera__device_number__cooleron
*/
func (c *Camera) TurnCoolerOff() error {
	c.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		// Set True to turn the camera cooler on:
		"CoolerOn":            fmt.Sprintf("%t", false),
		"ClientID":            fmt.Sprintf("%d", c.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", c.Alpaca.TransactionId),
	}

	return c.Alpaca.Put("camera", c.DeviceNumber, "cooleron", form)
}

/*
GetCoolerPowerLevel()

@returns the present cooler power level, in percent.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__coolerpower
*/
func (c *Camera) GetCoolerPowerLevel() (float64, error) {
	return c.Alpaca.GetFloat64Response("camera", c.DeviceNumber, "coolerpower")
}

/*
GetGainInElectronsPerADUnit()

@returns the gain of the camera in photoelectrons per A/D unit.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__electronsperadu
*/
func (c *Camera) GetGainInElectronsPerADUnit() (float64, error) {
	return c.Alpaca.GetFloat64Response("camera", c.DeviceNumber, "electronsperadu")
}

/*
GetExposure()

@returns Returns an array of 32bit integers containing the pixel values from the last exposure. This
call can return either a 2 dimension (monochrome images) or 3 dimension (colour or multi-plane images)
array of size NumX * NumY or NumX * NumY * NumPlanes. Where applicable, the size of NumPlanes has
to be determined by inspection of the returned Array.

Since 32bit integers are always returned by this call, the returned JSON Type value (0 = Unknown, 1 = short(16bit),
2 = int(32bit), 3 = Double) is always 2. The number of planes is given in the returned Rank value.

When de-serialising to an object it is essential to know the array Rank beforehand so that the correct data
class can be used. This can be achieved through a regular expression or by direct parsing of the returned
JSON string to extract the Type and Rank values before de-serialising.

@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__imagearray
*/
func (c *Camera) GetExposure() ([][]uint32, uint32, error) {
	return c.Alpaca.GetUInt32Rank2ArrayResponse("camera", c.DeviceNumber, "imagearray")
}

/*
GetExposureMax()

@returns the maximum exposure time supported by StartExposure.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__exposuremax
*/
func (c *Camera) GetExposureMax() (float64, error) {
	return c.Alpaca.GetFloat64Response("camera", c.DeviceNumber, "exposuremax")
}

/*
GetExposureMin()

@returns the maximum exposure time supported by StartExposure.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__exposuremin
*/
func (c *Camera) GetExposureMin() (float64, error) {
	return c.Alpaca.GetFloat64Response("camera", c.DeviceNumber, "exposuremin")
}

/*
GetExposureResolution()

@returns the smallest increment in exposure time supported by StartExposure.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__exposureresolution
*/
func (c *Camera) GetExposureResolution() (float64, error) {
	return c.Alpaca.GetFloat64Response("camera", c.DeviceNumber, "exposureresolution")
}

/*
IsFastReadoutEnabled()

@returns whenther Fast Readout Mode is enabled.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__fastreadout
*/
func (c *Camera) IsFastReadoutEnabled() (bool, error) {
	return c.Alpaca.GetBooleanResponse("camera", c.DeviceNumber, "fastreadout")
}

/*
EnableFastReadout()

@returns an error or nil, if nil it enables the fast readout mode
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/put_camera__device_number__fastreadout
*/
func (c *Camera) EnableFastReadout() error {
	c.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		// Set True to enable fast readout mode:
		"FastReadout":         fmt.Sprintf("%t", true),
		"ClientID":            fmt.Sprintf("%d", c.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", c.Alpaca.TransactionId),
	}

	return c.Alpaca.Put("camera", c.DeviceNumber, "fastreadout", form)
}

/*
DisableFastReadout()

@returns an error or nil, if nil it disables the fast readout mode
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/put_camera__device_number__fastreadout
*/
func (c *Camera) DisableFastReadout() error {
	c.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		// Set False to disable fast readout mode:
		"FastReadout":         fmt.Sprintf("%t", false),
		"ClientID":            fmt.Sprintf("%d", c.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", c.Alpaca.TransactionId),
	}

	return c.Alpaca.Put("camera", c.DeviceNumber, "fastreadout", form)
}

/*
GetFullWellCapacity()

@returns the full well capacity of the camera in electrons, at the current camera settings (binning, SetupDialog settings, etc.).
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__fullwellcapacity
*/
func (c *Camera) GetFullWellCapacity() (float64, error) {
	return c.Alpaca.GetFloat64Response("camera", c.DeviceNumber, "fullwellcapacity")
}

/*
GetGain()

@returns the camera's gain (GAIN VALUE MODE) OR the index of the selected camera gain description in the Gains array (GAINS INDEX MODE).
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__gain
*/
func (c *Camera) GetGain() (int32, error) {
	return c.Alpaca.GetInt32Response("camera", c.DeviceNumber, "gain")
}

/*
SetGain()

@returns an error or nil, if nil it sets the gain to the specified value.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/put_camera__device_number__gain
*/
func (c *Camera) SetGain(gain int32) error {
	c.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		// Set the gain (GAIN VALUE MODE) OR the index of the selected camera gain description in the Gains array (GAINS INDEX MODE).
		"Gain":                fmt.Sprintf("%d", gain),
		"ClientID":            fmt.Sprintf("%d", c.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", c.Alpaca.TransactionId),
	}

	return c.Alpaca.Put("camera", c.DeviceNumber, "gain", form)
}

/*
GetGainMax()

@returns the maximum value of Gain.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__gainmax
*/
func (c *Camera) GetGainMax() (int32, error) {
	return c.Alpaca.GetInt32Response("camera", c.DeviceNumber, "gainmax")
}

/*
GetGainMin()

@returns the minimum value of Gain.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__gainmin
*/
func (c *Camera) GetGainMin() (int32, error) {
	return c.Alpaca.GetInt32Response("camera", c.DeviceNumber, "gainmin")
}

/*
GetGains()

@returns the Gains supported by the camera.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__gains
*/
func (c *Camera) GetGains() ([]string, error) {
	return c.Alpaca.GetStringListResponse("camera", c.DeviceNumber, "gains")
}

/*
HasShutter()

@returns the true if this camera has a mechanical shutter.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__hasshutter
*/
func (c *Camera) HasShutter() (bool, error) {
	return c.Alpaca.GetBooleanResponse("camera", c.DeviceNumber, "hasshutter")
}

/*
GetHeatSinkTemperature()

@returns the current heat sink temperature (called "ambient temperature" by some manufacturers) in degrees Celsius.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__heatsinktemperature
*/
func (c *Camera) GetHeatSinkTemperature() (float64, error) {
	return c.Alpaca.GetFloat64Response("camera", c.DeviceNumber, "heatsinktemperature")
}

/*
IsImageReady()

@returns true if the image is ready to be downloaded from the camera
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__imageready
*/
func (c *Camera) IsImageReady() (bool, error) {
	return c.Alpaca.GetBooleanResponse("camera", c.DeviceNumber, "imageready")
}

/*
IsPulseGuiding()

@returns true if the the camera is currrently in a PulseGuide operation.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__ispulseguiding
*/
func (c *Camera) IsPulseGuiding() (bool, error) {
	return c.Alpaca.GetBooleanResponse("camera", c.DeviceNumber, "ispulseguiding")
}

/*
GetLastExposureStartTime()

@returns the actual exposure start in the FITS-standard CCYY-MM-DDThh:mm:ss[.sss...] format.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__lastexposurestarttime
*/
func (c *Camera) GetLastExposureStartTime() (*time.Time, error) {
	starttime, err := c.Alpaca.GetStringResponse("camera", c.DeviceNumber, "lastexposurestarttime")

	if err != nil {
		return nil, err
	}

	if starttime == "" {
		return nil, nil
	}

	t, err := time.Parse("2006-01-02T15:04:05", starttime)

	if err != nil {
		return nil, err
	}

	return &t, nil
}

/*
GetLastExposureDuration()

@returns the actual exposure duration in seconds (i.e. shutter open time).
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__lastexposureduration
*/
func (c *Camera) GetLastExposureDuration() (float64, error) {
	return c.Alpaca.GetFloat64Response("camera", c.DeviceNumber, "lastexposureduration")
}

/*
GetMaxADU()

@returns the maximum ADU value the camera can produce.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__maxadu
*/
func (c *Camera) GetMaxADU() (int32, error) {
	return c.Alpaca.GetInt32Response("camera", c.DeviceNumber, "maxadu")
}

/*
GetMaxBinX()

@returns the maximum allowed binning for the X camera axis.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__maxbinx
*/
func (c *Camera) GetMaxBinX() (int32, error) {
	return c.Alpaca.GetInt32Response("camera", c.DeviceNumber, "maxbinx")
}

/*
GetMaxBinY()

@returns the maximum allowed binning for the Y camera axis.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__maxbiny
*/
func (c *Camera) GetMaxBinY() (int32, error) {
	return c.Alpaca.GetInt32Response("camera", c.DeviceNumber, "maxbiny")
}

/*
GetSubFrameWidth()

@returns the current subframe width, if binning is active, value is in binned pixels.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__numx
*/
func (c *Camera) GetSubFrameWidth() (int32, error) {
	return c.Alpaca.GetInt32Response("camera", c.DeviceNumber, "numx")
}

/*
SetSubFrameWidth()

@returns an error or nil, if the subframe width can be set.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__numx
*/
func (c *Camera) SetSubFrameWidth(numX int32) error {
	c.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		// Set the subframe width in pixels.
		"NumX":                fmt.Sprintf("%d", numX),
		"ClientID":            fmt.Sprintf("%d", c.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", c.Alpaca.TransactionId),
	}

	return c.Alpaca.Put("camera", c.DeviceNumber, "numx", form)
}

/*
GetSubFrameHeight()

@returns the current subframe height, if binning is active, value is in binned pixels.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__numy
*/
func (c *Camera) GetSubFrameHeight() (int32, error) {
	return c.Alpaca.GetInt32Response("camera", c.DeviceNumber, "numy")
}

/*
SetSubFrameHeight()

@returns an error or nil, if the subframe height can be set.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__numy
*/
func (c *Camera) SetSubFrameHeight(numY int32) error {
	c.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		// Set the subframe height in pixels.
		"NumY":                fmt.Sprintf("%d", numY),
		"ClientID":            fmt.Sprintf("%d", c.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", c.Alpaca.TransactionId),
	}

	return c.Alpaca.Put("camera", c.DeviceNumber, "numy", form)
}

/*
GetCurrentOperationPercentageComplete()

@returns the percentage of the current operation that is complete. If valid, returns an integer between 0 and
100, where 0 indicates 0% progress (function just started) and 100 indicates 100% progress (i.e. completion).
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__percentcompleted
*/
func (c *Camera) GetCurrentOperationPercentageComplete() (int32, error) {
	return c.Alpaca.GetInt32Response("camera", c.DeviceNumber, "percentcompleted")
}

/*
GetPixelSizeX()

@returns the width of the CCD chip pixels in microns.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__pixelsizex
*/
func (c *Camera) GetPixelSizeX() (float64, error) {
	return c.Alpaca.GetFloat64Response("camera", c.DeviceNumber, "pixelsizex")
}

/*
GetPixelSizeY()

@returns the width of the CCD chip pixels in microns.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__pixelsizey
*/
func (c *Camera) GetPixelSizeY() (float64, error) {
	return c.Alpaca.GetFloat64Response("camera", c.DeviceNumber, "pixelsizey")
}

/*
GetReadOutMode()

@returns an index into the array ReadoutModes and returns the desired readout mode for the camera. Defaults to 0 if not set.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__readoutmode
*/
func (c *Camera) GetReadOutMode() (int32, error) {
	return c.Alpaca.GetInt32Response("camera", c.DeviceNumber, "readoutmode")
}

/*
SetReadOutMode()

@returns an error or nil, if nil sets the readout mode.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/put_camera__device_number__readoutmode
*/
func (c *Camera) SetReadOutMode(readOutMode int32) error {
	c.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		// Set the readout mode for the camera.
		"ReadoutMode":         fmt.Sprintf("%d", readOutMode),
		"ClientID":            fmt.Sprintf("%d", c.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", c.Alpaca.TransactionId),
	}

	return c.Alpaca.Put("camera", c.DeviceNumber, "readoutmode", form)
}

/*
GetReadOutModes()

@returns an array of strings, each of which describes an available readout mode of the camera. At least one string must be present in the list.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__readoutmodes
*/
func (c *Camera) GetReadOutModes() ([]string, error) {
	return c.Alpaca.GetStringListResponse("camera", c.DeviceNumber, "readoutmodes")
}

/*
GetSensorName()

@returns the name of the sensor used within the camera.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__sensorname
*/
func (c *Camera) GetSensorName() (string, error) {
	return c.Alpaca.GetStringResponse("camera", c.DeviceNumber, "sensorname")
}

/*
GetSensorType()

@returns the sensor type, e.g., whether the sensor is monochrome, or what Bayer matrix it encodes. Where:
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__sensortype
*/
func (c *Camera) GetSensorType() (SensorType, error) {
	sensor, err := c.Alpaca.GetInt32Response("camera", c.DeviceNumber, "sensortype")
	return SensorType(sensor), err
}

/*
GetCCDTemperatureCoolerSetPoint()

@returns the current temperature setpoint of the CCD cooler in degrees Celsius.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__setccdtemperature
*/
func (c *Camera) GetCCDTemperatureCoolerSetPoint() (float64, error) {
	return c.Alpaca.GetFloat64Response("camera", c.DeviceNumber, "setccdtemperature")
}

/*
SetCCDTemperatureCoolerSetPoint()

@returns an error or nil, if nil it sets the camera's cooler setpoint in degrees Celsius.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/put_camera__device_number__setccdtemperature
*/
func (c *Camera) SetCCDTemperatureCoolerSetPoint(temperature float64) error {
	c.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		"SetCCDTemperature":   fmt.Sprintf("%f", temperature),
		"ClientID":            fmt.Sprintf("%d", c.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", c.Alpaca.TransactionId),
	}

	return c.Alpaca.Put("camera", c.DeviceNumber, "setccdtemperature", form)
}

/*
GetStartX()

@returns the current subframe start X coordinate, if binning is active, value is in binned pixels.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__startx
*/
func (c *Camera) GetStartX() (int32, error) {
	return c.Alpaca.GetInt32Response("camera", c.DeviceNumber, "startx")
}

/*
SetStartX()

@returns an error or nil, if nil sets the subframe start X coordinate, if binning is active, value is in binned pixels.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__startx
*/
func (c *Camera) SetStartX(startX int32) error {
	c.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		"StartX":              fmt.Sprintf("%d", startX),
		"ClientID":            fmt.Sprintf("%d", c.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", c.Alpaca.TransactionId),
	}

	return c.Alpaca.Put("camera", c.DeviceNumber, "startx", form)
}

/*
GetStartY()

@returns the current subframe start Y coordinate, if binning is active, value is in binned pixels.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__starty
*/
func (c *Camera) GetStartY() (int32, error) {
	return c.Alpaca.GetInt32Response("camera", c.DeviceNumber, "starty")
}

/*
SetStartY()

@returns an error or nil, if nil sets the subframe start Y coordinate, if binning is active, value is in binned pixels.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__starty
*/
func (c *Camera) SetStartY(startY int32) error {
	c.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		"StartY":              fmt.Sprintf("%d", startY),
		"ClientID":            fmt.Sprintf("%d", c.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", c.Alpaca.TransactionId),
	}

	return c.Alpaca.Put("camera", c.DeviceNumber, "starty", form)
}

/*
GetSubExposureDuration()

@returns the sub exposure duration in seconds, *only available in Camera Interface Version 3 and later.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__subexposureduration
*/
func (c *Camera) GetSubExposureDuration() (float64, error) {
	return c.Alpaca.GetFloat64Response("camera", c.DeviceNumber, "subexposureduration")
}

/*
SetSubExposureDuration()

@returns an error or nil, if nil sets the sub exposure duration in seconds, *only available in Camera Interface Version 3 and later.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/put_camera__device_number__subexposureduration
*/
func (c *Camera) SetSubExposureDuration(subExposureDuration float64) error {
	c.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		"SubExposureDuration": fmt.Sprintf("%f", subExposureDuration),
		"ClientID":            fmt.Sprintf("%d", c.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", c.Alpaca.TransactionId),
	}

	return c.Alpaca.Put("camera", c.DeviceNumber, "subexposureduration", form)
}

/*
AbortExposure()

@returns an error or nil, if nil aborts the current exposure, if any, and returns the camera to Idle state.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/put_camera__device_number__abortexposure
*/
func (c *Camera) AbortExposure() error {
	c.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		"ClientID":            fmt.Sprintf("%d", c.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", c.Alpaca.TransactionId),
	}

	return c.Alpaca.Put("camera", c.DeviceNumber, "abortexposure", form)
}

/*
PulseGuide()

@returns an error or nil, if nil starts a pulse guide operation i.e., it activates the Camera's mount control system to instruct the mount to move in a particular direction for a given period of time
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/put_camera__device_number__pulseguide
*/
func (c *Camera) SetPulseGuide(direction Direction, duration int32) error {
	c.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		"Direction":           fmt.Sprintf("%d", direction),
		"Duration":            fmt.Sprintf("%d", duration),
		"ClientID":            fmt.Sprintf("%d", c.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", c.Alpaca.TransactionId),
	}

	return c.Alpaca.Put("camera", c.DeviceNumber, "pulseguide", form)
}

/*
StartExposure()

@returns an error or nil, if nil starts an exposure (use IsImageReady to check when the exposure is complete.)
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/put_camera__device_number__startexposure
*/
func (c *Camera) StartExposure(duration float64, light bool) error {
	c.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		"Duration":            fmt.Sprintf("%f", duration),
		"Light":               fmt.Sprintf("%t", light),
		"ClientID":            fmt.Sprintf("%d", c.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", c.Alpaca.TransactionId),
	}

	return c.Alpaca.Put("camera", c.DeviceNumber, "startexposure", form)
}

/*
StopExposure()

@returns an error or nil, if nil stops the current exposure, if any. If an exposure is in progress, the readout process is initiated. Ignored if readout is already in process.
@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/put_camera__device_number__stopexposure
*/
func (c *Camera) StopExposure() error {
	c.Alpaca.TransactionId++

	var form map[string]string = map[string]string{
		"ClientID":            fmt.Sprintf("%d", c.Alpaca.ClientId),
		"ClientTransactionID": fmt.Sprintf("%d", c.Alpaca.TransactionId),
	}

	return c.Alpaca.Put("camera", c.DeviceNumber, "stopexposure", form)
}
