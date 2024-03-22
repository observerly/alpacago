package alpacago

import "testing"

var calibrator = NewCoverCalibrator(65535, false, "100.69.47.32", "", -1, 0)

func TestNewCoverCalibratorBaseURL(t *testing.T) {
	calibrator := NewCoverCalibrator(65535, false, "", "0.0.0.0", 8000, 0)

	var got string = calibrator.Alpaca.UrlBase
	var want string = "http://0.0.0.0:8000"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewCoverCalibratorBaseURLForHost(t *testing.T) {
	var got string = calibrator.Alpaca.UrlBase
	var want string = "http://100.69.47.32"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewCoverCalibratorClientID(t *testing.T) {
	calibrator := NewCoverCalibrator(65535, false, "", "0.0.0.0", 8000, 0)

	var got uint32 = calibrator.Alpaca.ClientId
	var want uint32 = 65535

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewCoverCalibratorTransactionID(t *testing.T) {
	calibrator := NewCoverCalibrator(65535, false, "", "0.0.0.0", 8000, 0)

	var got uint32 = calibrator.Alpaca.TransactionId
	var want uint32 = 0

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewCoverCalibratorDeviceNumber(t *testing.T) {
	calibrator := NewCoverCalibrator(65535, false, "", "0.0.0.0", 8000, 0)

	var got uint = calibrator.DeviceNumber
	var want uint = 0

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewCoverCalibratorIsConnected(t *testing.T) {
	var got, err = calibrator.IsConnected()

	var want = true

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	if calibrator.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", calibrator.Alpaca.ErrorMessage)
	}
}

func TestNewCoverCalibratorSetConnected(t *testing.T) {
	calibrator.SetConnected(true)

	var got, err = calibrator.IsConnected()

	var want = true

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	if calibrator.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", calibrator.Alpaca.ErrorMessage)
	}
}

func TestNewCoverCalibratorGetBrightness(t *testing.T) {
	var got, err = calibrator.GetBrightness()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < 0 {
		t.Errorf("got %v, but expected a physically realistic brightness value", got)
	}

	if calibrator.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", calibrator.Alpaca.ErrorMessage)
	}
}

func TestNewCalibratorCoverGetStatus(t *testing.T) {
	calibrator.SetConnected(true)

	calibrator.SetCalibratorOn(90)

	var got, err = calibrator.GetStatus()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != CalibratorReady && got != CalibratorNotReady {
		t.Errorf("got %v, but expected the calibrator to be ready", got)
	}

	if calibrator.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", calibrator.Alpaca.ErrorMessage)
	}
}

func TestNewCalibratorCoverGetCoverStatus(t *testing.T) {
	calibrator.SetConnected(true)

	calibrator.SetCalibratorOn(90)

	var got, err = calibrator.GetCoverStatus()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != CoverClosed && got != CoverOpen && got != CoverMoving && got != CoverUnknown {
		t.Errorf("got %v, but expected the calibrator to be either open, close or moving", got)
	}

	if calibrator.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", calibrator.Alpaca.ErrorMessage)
	}
}

func TestNewCalibratorCoverGetMaxBrightness(t *testing.T) {
	calibrator.SetConnected(true)

	var got, err = calibrator.GetMaxBrightness()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < 0 {
		t.Errorf("got %v, but expected a physically realistic brightness value", got)
	}

	if calibrator.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", calibrator.Alpaca.ErrorMessage)
	}
}

func TestNewCalibratorCoverSetCalibratorOn(t *testing.T) {
	calibrator.SetConnected(true)

	var err = calibrator.SetCalibratorOn(90)

	if err != nil {
		t.Errorf("got %q", err)
	}

	if calibrator.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", calibrator.Alpaca.ErrorMessage)
	}
}

func TestNewCalibratorCoverSetCalibratorOff(t *testing.T) {
	calibrator.SetConnected(true)

	var err = calibrator.SetCalibratorOff()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if calibrator.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", calibrator.Alpaca.ErrorMessage)
	}
}

func TestNewCalibratorCoverCloseCover(t *testing.T) {
	calibrator.SetConnected(true)

	var err = calibrator.CloseCover()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if calibrator.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", calibrator.Alpaca.ErrorMessage)
	}
}

func TestNewCalibratorCoverHaltCover(t *testing.T) {
	calibrator.SetConnected(true)

	var err = calibrator.HaltCover()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if calibrator.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", calibrator.Alpaca.ErrorMessage)
	}
}

func TestNewCalibratorCoverOpenCover(t *testing.T) {
	calibrator.SetConnected(true)

	var err = calibrator.OpenCover()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if calibrator.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", calibrator.Alpaca.ErrorMessage)
	}
}
