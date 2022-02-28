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
