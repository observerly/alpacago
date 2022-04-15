package alpacago

import "fmt"

type OperationalState int32

const (
	CameraIdle OperationalState = iota
	CameraWaiting
	CameraExposing
	CameraReading
	CameraDownload
	CameraError
)

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
	GetBinY()

	@returns the binning factor for the Y axis.
	@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__biny
*/
func (c *Camera) GetBinY() (int32, error) {
	return c.Alpaca.GetInt32Response("camera", c.DeviceNumber, "biny")
}

/*
	GetOperationalState()

	@returns the current camera operational state (CameraIdle, CameraWaiting, CameraExposing, CameraReading, CameraDownload, CameraError)
	The operational state is specified as an integer value from the OperationalState Enum.
	@see https://ascom-standards.org/api/#/Camera%20Specific%20Methods/get_camera__device_number__camerastate
*/
func (c *Camera) GetOperationalState() (OperationalState, error) {
	state, err := c.Alpaca.GetInt32Response("camera", c.DeviceNumber, "camerastate")
	return OperationalState(state), err
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
