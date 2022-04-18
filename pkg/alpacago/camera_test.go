package alpacago

import "testing"

var camera = NewCamera(65535, true, "alpaca.observerly.com", "", -1, 0)

func TestNewCameraBaseURL(t *testing.T) {
	camera := NewCamera(65535, false, "", "0.0.0.0", 8000, 0)

	var got string = camera.Alpaca.UrlBase
	var want string = "http://0.0.0.0:8000"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewCameraBaseURLForHost(t *testing.T) {
	var got string = camera.Alpaca.UrlBase
	var want string = "https://alpaca.observerly.com"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewCameraClientID(t *testing.T) {
	camera := NewCamera(65535, false, "", "0.0.0.0", 8000, 0)

	var got uint32 = camera.Alpaca.ClientId
	var want uint32 = 65535

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewCameraTransactionID(t *testing.T) {
	camera := NewCamera(65535, false, "", "0.0.0.0", 8000, 0)

	var got uint32 = camera.Alpaca.TransactionId
	var want uint32 = 0

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewCameraDeviceNumber(t *testing.T) {
	camera := NewCamera(65535, false, "", "0.0.0.0", 8000, 0)

	var got uint = camera.DeviceNumber
	var want uint = 0

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewCameraSetConnectedOn(t *testing.T) {
	var err = camera.SetConnected(true)

	if err != nil {
		t.Errorf("got %q", err)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraIsConnectedOn(t *testing.T) {
	var got, err = camera.IsConnected()

	var want = true

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraSetConnectedOff(t *testing.T) {
	var err = camera.SetConnected(false)

	if err != nil {
		t.Errorf("got %q", err)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraIsConnectedOff(t *testing.T) {
	var got, err = camera.IsConnected()

	var want = false

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %t wanted %t", got, want)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraGetBayerOffsetX(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.GetBayerOffsetX()

	var want int32 = 0

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %q wanted %q", got, want)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraGetBayerOffsetY(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.GetBayerOffsetY()

	var want int32 = 0

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %q wanted %q", got, want)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraGetBinX(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.GetBinX()

	var want int32 = 1

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %v wanted %v", got, want)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraGetBinY(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.GetBinY()

	var want int32 = 1

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %v wanted %v", got, want)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraGetOperationalState(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.GetOperationalState()

	var want = CameraIdle

	if err != nil {
		t.Errorf("got %q, wanted %q", err, want)
	}

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if telescope.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %d", telescope.Alpaca.ErrorMessage, want)
	}
}

func TestNewCameraGetCCDSizeX(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.GetCCDSizeX()

	var want int32 = 800

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %v wanted %v", got, want)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraGetCCDSizeY(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.GetCCDSizeY()

	var want int32 = 600

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %v wanted %v", got, want)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraCanAbortExposure(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.CanAbortExposure()

	var want bool = true

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %v wanted %v", got, want)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraCanAsymmetricBin(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.CanAsymmetricBin()

	var want bool = true

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %v wanted %v", got, want)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraCanFastReadout(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.CanFastReadout()

	var want bool = false

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %v wanted %v", got, want)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraCanGetCoolerPower(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.CanGetCoolerPower()

	var want bool = false

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %v wanted %v", got, want)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraCanPulseGuide(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.CanPulseGuide()

	var want bool = false

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %v wanted %v", got, want)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraCanSetCCDTemperature(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.CanSetCCDTemperature()

	var want bool = false

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %v wanted %v", got, want)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraCanStopExposure(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.CanStopExposure()

	var want bool = true

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %v wanted %v", got, want)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraGetCCDTemperature(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.GetCCDTemperature()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < -273.15 && got < 100 {
		t.Errorf("got %v, but expected the CCD temperature to be a realistic temperature", got)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraIsCoolerOn(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.IsCoolerOn()

	var want bool = false

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %v wanted %v", got, want)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraGetCoolerPowerLevel(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.GetCoolerPowerLevel()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < 0 && got > 100 {
		t.Errorf("got %v, but expected the CCD power level to be a percentage", got)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraGetGainInElectronsPerADUnit(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.GetGainInElectronsPerADUnit()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < 0 && got > 10 {
		t.Errorf("got %v, but expected the gain in photoelectrons per A/D unit to be a physically realistic value", got)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraGetExposureMax(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.GetExposureMax()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < 0 && got > 10000 {
		t.Errorf("got %v, but expected the maximum exposure value to be a realistic value", got)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraGetExposureMin(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.GetExposureMin()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < 0 && got > 1 {
		t.Errorf("got %v, but expected the minimum exposure value to be a realistic value", got)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraGetExposureResolution(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.GetExposureResolution()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < 0 && got > 1 {
		t.Errorf("got %v, but expected the exposure resolution value to be a realistic value", got)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraIsFastReadoutEnabled(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.IsFastReadoutEnabled()

	var want = false

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraGetFullWellCapacity(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.GetFullWellCapacity()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < 0 && got > 1 {
		t.Errorf("got %v, but expected the full-well capacity value to be a realistic value", got)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraGetGain(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.GetGain()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < 0 && got > 100 {
		t.Errorf("got %v, but expected the gain value to be a realistic value", got)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraGetGainMax(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.GetGainMax()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < 0 && got > 100 {
		t.Errorf("got %v, but expected the maximum gain value to be a realistic value", got)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraGetGainMin(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.GetGainMin()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < 0 && got > 10 {
		t.Errorf("got %v, but expected the minimum gain value to be a realistic value", got)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraGetGainMinMax(t *testing.T) {
	camera.SetConnected(true)

	var max, errMax = camera.GetGainMax()

	var min, errMin = camera.GetGainMin()

	if errMax != nil {
		t.Errorf("got %q", errMax)
	}

	if errMin != nil {
		t.Errorf("got %q", errMin)
	}

	if min > max {
		t.Errorf("got %v, but expected the minimum gain value to be less than the maximum gain value", min)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraGetGains(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.GetGains()

	var want = []string{"ISO 100", "ISO 200", "ISO 400", "ISO 800", "ISO 1600"}

	if err != nil {
		t.Errorf("got %q", err)
	}

	if len(got) != len(want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraHasShutter(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.HasShutter()

	var want = true

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraGetHeatSinkTemperature(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.GetHeatSinkTemperature()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < 100 && got > 100 {
		t.Errorf("got %v, but expected the heat sink temperature to be a realistic physical value", got)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraIsImageReady(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.IsImageReady()

	var want bool = false

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %v wanted %v", got, want)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraIsPulseGuiding(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.IsPulseGuiding()

	var want bool = false

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %v wanted %v", got, want)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraGetLastExposureDuration(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.GetLastExposureDuration()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < 0 && got > 1 {
		t.Errorf("got %v, but expected exposure duration value to be a realistic value", got)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraGetMaxADU(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.GetMaxADU()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < 0 && got > 10000 {
		t.Errorf("got %v, but expected the maximum ADU value to be a realistic value", got)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraGetMaxBinX(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.GetMaxBinX()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < 0 && got > 10000 {
		t.Errorf("got %v, but expected the maximum binx value to be a realistic value", got)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraGetMaxBinY(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.GetMaxBinY()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < 0 && got > 10000 {
		t.Errorf("got %v, but expected the maximum biny value to be a realistic value", got)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraGetSubFrameWidth(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.GetSubFrameWidth()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < 0 && got > 10000 {
		t.Errorf("got %v, but expected the subframe width value to be a realistic value", got)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraGetSubFrameHeight(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.GetSubFrameHeight()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < 0 && got > 10000 {
		t.Errorf("got %v, but expected the subframe height value to be a realistic value", got)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraGetCurrentOperationPercentageComplete(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.GetCurrentOperationPercentageComplete()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < 0 && got > 100 {
		t.Errorf("got %v, but expected the percentage completion to be between 0 and 100", got)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraGetPixelSizeX(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.GetPixelSizeX()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < 0 && got > 10000 {
		t.Errorf("got %v, but expected the maximum binx value to be a realistic value", got)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraGetPixelSizeY(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.GetPixelSizeY()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < 0 && got > 10000 {
		t.Errorf("got %v, but expected the maximum biny value to be a realistic value", got)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraGetReadOutMode(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.GetReadOutMode()

	var want int32 = 0

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %v, but expected the read out mode to be the default value", got)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}
