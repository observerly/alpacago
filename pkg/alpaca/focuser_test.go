package alpaca

import (
	"testing"
)

var focuser = NewFocuser(65535, true, "alpaca.observerly.com", "", -1, 0)

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
	var want string = "https://alpaca.observerly.com"

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

func TestNewFocuserIsAbsolute(t *testing.T) {
	var got, err = focuser.IsAbsolute()

	var want = true || false

	if err != nil {
		t.Errorf("got %q", err)
	}

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}

	if telescope.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", telescope.Alpaca.ErrorMessage)
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

	if telescope.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", telescope.Alpaca.ErrorMessage)
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

	if telescope.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", telescope.Alpaca.ErrorMessage)
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

	if telescope.Alpaca.ErrorNumber != 0 {
		t.Errorf("got %q", telescope.Alpaca.ErrorMessage)
	}
}
