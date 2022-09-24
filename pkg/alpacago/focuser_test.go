package alpacago

import (
	"testing"
)

var focuser = NewFocuser(65535, false, "100.80.84.116", "", -1, 0)

func TestNewFocuserBaseURL(t *testing.T) {
	focuser := NewFocuser(65535, false, "", "0.0.0.0", 8000, 0)

	var got string = focuser.Alpaca.UrlBase
	var want string = "http://0.0.0.0:8000"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewFocuserBaseURLForHost(t *testing.T) {
	var got string = focuser.Alpaca.UrlBase
	var want string = "http://100.80.84.116"

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewFocuserClientID(t *testing.T) {
	focuser := NewFocuser(65535, false, "", "0.0.0.0", 8000, 0)

	var got uint32 = focuser.Alpaca.ClientId
	var want uint32 = 65535

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewFocuserTransactionID(t *testing.T) {
	focuser := NewFocuser(65535, false, "", "0.0.0.0", 8000, 0)

	var got uint32 = focuser.Alpaca.TransactionId
	var want uint32 = 0

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewFocuserDeviceNumber(t *testing.T) {
	focuser := NewFocuser(65535, false, "", "0.0.0.0", 8000, 0)

	var got uint = focuser.DeviceNumber
	var want uint = 0

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestNewFocuserSetConnected(t *testing.T) {
	var err = focuser.SetConnected(true)

	if err != nil {
		t.Errorf("got %q", err)
	}

	if focuser.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", focuser.Alpaca.ErrorMessage)
	}
}

func TestNewFocuserIsAbsolute(t *testing.T) {
	var got, err = focuser.IsAbsolute()

	var want = true || false

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	if focuser.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", focuser.Alpaca.ErrorMessage)
	}
}

func TestNewFocuserIsMoving(t *testing.T) {
	var got, err = focuser.IsMoving()

	var want = false

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	if focuser.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", focuser.Alpaca.ErrorMessage)
	}
}

func TestNewFocuserGetMaxIncrement(t *testing.T) {
	var got, err = focuser.GetMaxIncrement()

	var want int32 = 50000

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %b, wanted %b", got, want)
	}

	if focuser.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", focuser.Alpaca.ErrorMessage)
	}
}

func TestNewFocuserGetMaxStep(t *testing.T) {
	var got, err = focuser.GetMaxStep()

	var want int32 = 50000

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

	if focuser.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", focuser.Alpaca.ErrorMessage)
	}
}

func TestNewFocuserGetPosition(t *testing.T) {
	var got, err = focuser.GetPosition()

	var want int32 = 25000

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

	if focuser.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", focuser.Alpaca.ErrorMessage)
	}
}

func TestNewFocuserGetStepSize(t *testing.T) {
	var got, err = focuser.GetStepSize()

	var want float64 = 20.0

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %f, wanted %f", got, want)
	}

	if focuser.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", focuser.Alpaca.ErrorMessage)
	}
}

func TestNewFocuserGetTemperatureCompensation(t *testing.T) {
	var got, err = focuser.GetTemperatureCompensation()

	var want = false

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	if focuser.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", focuser.Alpaca.ErrorMessage)
	}
}

func TestNewFocuserSetTemperatureCompensation(t *testing.T) {
	var err = focuser.SetTemperatureCompensation(false)

	if err != nil {
		t.Errorf("got %q", err)
	}

	if focuser.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", focuser.Alpaca.ErrorMessage)
	}
}

func TestNewFocuserIsTemperatureCompensationAvailable(t *testing.T) {
	var got, err = focuser.IsTemperatureCompensationAvailable()

	var want = true

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	if focuser.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", focuser.Alpaca.ErrorMessage)
	}
}

func TestNewFocuserGetTemperature(t *testing.T) {
	var got, err = focuser.GetTemperature()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got < -273.15 || got > 120 {
		t.Errorf("got %f, but expected the temperature of the focuser to be a realistic physical value", got)
	}

	if focuser.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", focuser.Alpaca.ErrorMessage)
	}
}

func TestNewFocuserSetHalt(t *testing.T) {
	// To halt, we need to get the focuser moving first!
	var got, err = focuser.IsMoving()

	if err != nil {
		t.Errorf("got %q", err)
	}

	if focuser.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", focuser.Alpaca.ErrorMessage)
	}

	if got {
		var err = focuser.SetHalt()

		if err != nil {
			t.Errorf("got %q", err)
		}

		if focuser.Alpaca.ErrorNumber != 0 {
			t.Errorf("got %q", focuser.Alpaca.ErrorMessage)
		}
	}
}

func TestNewFocuserSetMove(t *testing.T) {
	var _ = focuser.SetConnected(true)

	var err = focuser.SetMove(25000)

	if err != nil {
		t.Errorf("got %q", err)
	}

	if focuser.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", focuser.Alpaca.ErrorMessage)
	}
}
