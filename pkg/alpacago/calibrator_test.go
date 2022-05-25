package alpacago

import "testing"

var calibrator = NewCoverCalibrator(65535, true, "alpaca.observerly.com", "", -1, 0)

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
	var want string = "https://alpaca.observerly.com"

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
