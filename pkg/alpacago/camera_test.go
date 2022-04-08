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
