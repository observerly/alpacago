package alpacago

import (
	"testing"
)

var camera = NewCamera(65535, false, "100.69.47.32", "", -1, 0)

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
	var want string = "http://100.69.47.32"

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

func TestNewCameraSetBinX(t *testing.T) {
	camera.SetBinX(1)

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

func TestNewCameraSetBinY(t *testing.T) {
	camera.SetBinY(1)

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

	var want = "idle"

	if err != nil {
		t.Errorf("got %q, wanted %q", err, want)
	}

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q, wanted %q", camera.Alpaca.ErrorMessage, want)
	}
}

func TestNewCameraGetOperationalStateToStringRepresentation(t *testing.T) {
	var state = OperationalState.String(CameraIdle)

	if state == "" {
		t.Errorf("got %q, wanted %q", state, "the operational status to represnet an iota in range 0 to 5")
	}
}

func TestNewCameraGetCCDSizeX(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.GetCCDSizeX()

	var want int32 = 1463

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

	var want int32 = 1168

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

func TestNewCameraTurnCoolerOff(t *testing.T) {
	camera.SetConnected(true)

	camera.TurnCoolerOff()

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

func TestNewCameraIsCoolerOff(t *testing.T) {
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

func TestNewCameraTurnCoolerOn(t *testing.T) {
	camera.SetConnected(true)

	camera.TurnCoolerOn()

	var got, err = camera.IsCoolerOn()

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

func TestNewCameraIsCoolerOn(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.IsCoolerOn()

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

func TestNewCameraDisableFastReadout(t *testing.T) {
	camera.SetConnected(true)

	camera.DisableFastReadout()

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

func TestNewCameraEnableFastReadout(t *testing.T) {
	camera.SetConnected(true)

	camera.EnableFastReadout()

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

func TestNewCameraSetGain(t *testing.T) {
	camera.SetConnected(true)

	camera.SetGain(50)

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

	var want = []string{}

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

func TestNewCameraGetLastExposureStartTime(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.GetLastExposureStartTime()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != nil && got.Unix() < 1621033200 {
		t.Errorf("got %v, but expected the last exposure start time to be a realistic value", got)
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

	if got < 0 && got > 3600 {
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

func TestNewCameraSetSubFrameWidth(t *testing.T) {
	camera.SetConnected(true)

	camera.SetSubFrameWidth(800)

	var got, err = camera.GetSubFrameWidth()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < 0 && got > 800 {
		t.Errorf("got %v, but expected the subframe width value to be a realistic value. The valid range is: 1 to 800.", got)
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

	if got < 0 && got > 800 {
		t.Errorf("got %v, but expected the subframe width value to be a realistic value. The valid range is: 1 to 800.", got)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraSetSubFrameHeight(t *testing.T) {
	camera.SetConnected(true)

	camera.SetSubFrameHeight(600)

	var got, err = camera.GetSubFrameHeight()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < 0 && got > 600 {
		t.Errorf("got %v, but expected the subframe height value to be a realistic value. The valid range is: 1 to 600.", got)
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

	if got < 0 && got > 600 {
		t.Errorf("got %v, but expected the subframe height value to be a realistic value. The valid range is: 1 to 600.", got)
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

func TestNewCameraSetReadOutMode(t *testing.T) {
	camera.SetConnected(true)

	camera.SetReadOutMode(0)

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

func TestNewCameraGetReadOutModes(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.GetReadOutModes()

	var want = []string{"Default"}

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got[0] != want[0] {
		t.Errorf("got %v, but expected the read out mode to be the default value", got)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraGetSensorName(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.GetSensorName()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != "" {
		t.Errorf("got %v, but expected the sensor name to be a realistic value", got)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraGetSensorType(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.GetSensorType()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != Monochrome {
		t.Errorf("got %v, but expected the sensor type to be default to Monochrome", got)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraGetSensorTypeString(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.GetSensorType()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got.String() != "Monochrome" {
		t.Errorf("got %v, but expected the sensor type to be default to Monochrome", got)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraGetSensorTypeStringIota(t *testing.T) {
	if SensorType.String(0) != "Monochrome" {
		t.Errorf("Expected the sensor type to be default to Monochrome")
	}

	if SensorType.String(1) != "Color" {
		t.Errorf("Expected the sensor type to be default to Color")
	}

	if SensorType.String(2) != "RGGB" {
		t.Errorf("Expected the sensor type to be default to RGGB")
	}

	if SensorType.String(3) != "CMYG" {
		t.Errorf("Expected the sensor type to be default to CMYG")
	}

	if SensorType.String(4) != "CMYG2" {
		t.Errorf("Expected the sensor type to be default to CMYG2")
	}

	if SensorType.String(5) != "LRGB" {
		t.Errorf("Expected the sensor type to be default to LRGB")
	}
}

func TestNewCameraGetCCDTemperatureCoolerSetPoint(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.GetCCDTemperatureCoolerSetPoint()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < -273.15 && got > 1000 {
		t.Errorf("got %v, but expected the CCD temperature cooler set point to be a realistic value", got)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraSetCCDTemperatureCoolerSetPoint(t *testing.T) {
	camera.SetConnected(true)

	camera.SetCCDTemperatureCoolerSetPoint(4.5)

	var got, err = camera.GetCCDTemperatureCoolerSetPoint()

	var want float64 = 4.5

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < -273.15 && got > 1000 {
		t.Errorf("got %v, but expected the CCD temperature cooler set point to be a realistic value", got)
	}

	if got != want {
		t.Errorf("got %v, but expected the CCD temperature cooler set point to be set correctly", got)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraSetStartX(t *testing.T) {
	camera.SetConnected(true)

	camera.SetStartX(400)

	var got, err = camera.GetStartX()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < 0 && got > 800 {
		t.Errorf("got %v, but expected the subframe start position for the X axis value to be a realistic value. The valid range is: 0 to 800.", got)
	}

	if got != 400 {
		t.Errorf("got %v, but expected the subframe start position for the X axis value to be set correctly", got)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraGetStartX(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.GetStartX()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < 0 && got > 800 {
		t.Errorf("got %v, but expected the subframe start position for the X axis value to be a realistic value. The valid range is: 0 to 800.", got)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraSetStartY(t *testing.T) {
	camera.SetConnected(true)

	camera.SetStartY(300)

	var got, err = camera.GetStartY()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < 0 && got > 600 {
		t.Errorf("got %v, but expected the subframe start position for the Y axis value to be a realistic value. The valid range is: 0 to 600.", got)
	}

	if got != 300 {
		t.Errorf("got %v, but expected the subframe start position for the Y axis value to be set correctly", got)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraGetStartY(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.GetStartY()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < 0 && got > 600 {
		t.Errorf("got %v, but expected the subframe start position for the Y axis value to be a realistic value. The valid range is: 0 to 600.", got)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraGetSubExposureDuration(t *testing.T) {
	camera.SetConnected(true)

	var got, err = camera.GetSubExposureDuration()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < 0 && got > 10000 {
		t.Errorf("got %v, but expected the subframe exposure duration to be a realistic value", got)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraSetSubExposureDuration(t *testing.T) {
	camera.SetConnected(true)

	camera.SetSubExposureDuration(90.0)

	var got, err = camera.GetSubExposureDuration()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != 90.0 {
		t.Errorf("got %v, but expected the subframe exposure duration to be set correctly", got)
	}

	if got < 0 && got > 10000 {
		t.Errorf("got %v, but expected the subframe exposure duration to be a realistic value", got)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraAbortExposure(t *testing.T) {
	camera.SetConnected(true)

	var err = camera.AbortExposure()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

// func TestNewCameraSetPulseGuide(t *testing.T) {
// 	camera.SetConnected(true)

// 	var err = camera.SetPulseGuide(1, 50)

// 	if err != nil {
// 		t.Errorf("got %q", err)
// 	}

// 	if camera.Alpaca.ErrorNumber != 0 {
// 		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
// 	}
// }

func TestNewCameraStartExposure(t *testing.T) {
	camera.SetConnected(true)

	camera.SetStartX(0)
	camera.SetStartY(0)

	var err = camera.StartExposure(20, true)

	if err != nil {
		t.Errorf("got %q", err)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraStopExposure(t *testing.T) {
	camera.SetConnected(true)

	var err = camera.StopExposure()

	if err != nil {
		t.Errorf("got %q", err)
	}

	var got, _ = camera.GetOperationalState()

	var want = "idle"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}

func TestNewCameraGetExposure(t *testing.T) {
	camera.SetConnected(true)

	var got, rank, err = camera.GetExposure()

	ysize := len(got[0])

	xsize := len(got)

	if xsize != 800 {
		t.Errorf("got %v, but expected the CCD width to be 800 pixels", xsize)
	}

	if ysize != 600 {
		t.Errorf("got %v, but expected the CCD height to be 600 pixels", ysize)
	}

	if err != nil {
		t.Errorf("got %q", err)
	}

	if rank > 3 {
		t.Errorf("got %v, but expected the rank of the array to be a realiostic value", rank)
	}

	if camera.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", camera.Alpaca.ErrorMessage)
	}
}
